/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:55
 * Filename      : manager.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"cluster"
	"component"
)

type Manager struct {
	db           *gorm.DB
	componentMgr component.Manager
	clusterMgr   cluster.Manager
}

func NewManager(componentMgr component.Manager, clusterMgr cluster.Manager, dbName string) *Manager {
	db, err := gorm.Open("sqlite3", dbName+".db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ProductModel{})
	db.AutoMigrate(&ProductInstModel{})
	db.AutoMigrate(&ComponentObjModel{})
	db.AutoMigrate(&DesignModel{})

	m := &Manager{db: db,
		componentMgr: componentMgr,
		clusterMgr:   clusterMgr}

	m.coordinate()
	return m
}

func (m *Manager) coordinate() {
	go func() {
		for {
			prodInstModels := make([]ProductInstModel, 0)
			m.db.Where("status='init' or status='update'").Find(&prodInstModels)
			for _, model := range prodInstModels {
				design, _ := m.GetDesign(model.ProductName, model.DesignRevision)
				for _, r := range design.ComponentRefers {
					fmt.Println("coordinate ", r)
					err1 := m.CreateComponentObj(model, r)
					err2 := m.UpdateComponentObj(model, r, design.Revision)
					if err1 != nil || err2 != nil {
						fmt.Println("coordinate err", err1, err2)
						//TODO
					}
					for {
						if err := m.ComponentObjReady(model, r.Role, r.ComponentName); err != nil {
							fmt.Println("not ready", r.ComponentName, r.Role)
							time.Sleep(time.Second)
						} else {
							break
						}
					}
				}
				m.db.Model(&ProductInstModel{}).Where("inst_id=?", model.InstId).Update("status", StatusReady)
			}
			m.db.Where("status ='delete'").Find(&prodInstModels)
			//TODO delete
			time.Sleep(time.Second)
		}
	}()
}

func (m *Manager) NewProductInst(prodInst *ProductInst) error {
	//记录数据
	instModel := ProductInst2Model(prodInst)
	err := m.db.Create(&instModel).Error
	if err != nil {
		return err
	}
	//创建产品对应的组件
	design, err := m.GetDesign(prodInst.ProductName, prodInst.DesignRevision)
	if err != nil {
		return err
	}
	for _, refer := range design.ComponentRefers {
		if err := m.CreateComponentObj(instModel, refer); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) GetProductInst(prodInstId string) (model ProductInstModel, err error) {
	err = m.db.Where("inst_id=?", prodInstId).First(&model).Error
	return
}
func (m *Manager) GetProductInst1(prodName, instName string) (model ProductInstModel, err error) {
	err = m.db.Where("product_name=? and name=?", prodName, instName).First(&model).Error
	return
}

func (m *Manager) GetProductInsts(prodName string) (models []ProductInstModel, err error) {
	err = m.db.Where("product_name=?", prodName).Find(&models).Error
	return
}

func (m *Manager) UpdateProductInst(prodInstId string, revision string) error {
	prodInst, err := m.GetProductInst(prodInstId)
	if err != nil {
		return err
	}
	d, err := m.GetDesign(prodInst.ProductName, revision)
	if err != nil {
		return err
	}
	if err = m.CanUpdateDesign(prodInst, d); err != nil {
		return err
	}
	return m.db.Model(&ProductInstModel{}).Where("inst_id = ?", prodInst.InstId).Updates(map[string]interface{}{"status": StatusUpdate, "design_revision": revision}).Error
}

func (m *Manager) CanUpdateDesign(p ProductInstModel, d Design) error {
	d1, err := m.GetDesign(p.ProductName, p.DesignRevision)
	if err != nil {
		return err
	}

	for _, j := range d1.ComponentRefers {
		for _, k := range d.ComponentRefers {
			if j.ComponentName == k.ComponentName {
				err, _ := m.componentMgr.Excute(j.ComponentName, "UPDATE_CHECK", p.KubeClusterName, k.Input, p.InstId, j.Role)
				if err != nil {
					return fmt.Errorf("component %s cannot update", j.ComponentName)
				}
			}
		}
	}
	return nil
}
