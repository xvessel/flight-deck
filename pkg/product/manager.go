package product

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"cluster"
	"component"
)

type Manager struct {
	db           *gorm.DB
	componentMgr component.Manager
	clusterMgr   cluster.ClusterMgr
}

func NewManager() *Manager {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ProductModel{})
	db.AutoMigrate(&ProductInstModel{})
	db.AutoMigrate(&ComponentObjModel{})

	return &Manager{db: db}
}

func (m *Manager) NewProduct(prodName string) {
	m.db.Create(&ProductModel{Name: prodName})
}

func (m *Manager) NewDesign(m *Design) {
	m.db.Create(Design2Model(m))
}

func (m *Manager) GetDesign(productName string, revision string) Design {
}

func (m *Manager) GetCompnentObj(prodId string, role string) ComponentObj {
	var model ComponentObjModel
	m.db.Where("product_id = ? and role = ?", prodId, role).First(&model)
	obj := ComponentObj{
		ProductInstId:   model.ProductInstId,
		ComponentName:   model.ComponentName,
		KubeClusterName: model.KubeClusterName,
		Role:            model.Role,
		Input:           make(map[string]string),
		Output:          make(map[string]string),
	}
	json.Unmarshal([]byte(model.InputJson), &obj.Input)
	json.Unmarshal([]byte(model.OutputJson), &obj.Output)
	return obj
}

func (m *Manger) CreateComponent(refer ComponentRefer) error {
	args := make(map[string]interface{})
	//渲染参数
	for _, r := range refer.PreRole {
		component := m.GetCompnentObj(prodInst.InstId, r)
		args[r] = component.Output
	}
	for inK, inV := range refer.Input {
		buf := new(bytes.Buffer)
		tpl, _ := template.New("letter").Parse(inV)
		tpl.Excute(buf, args)
		refer.Input[inK] = buf.String()
	}
	//创建组件
	conf, ok := m.ClusterMgr.GetKubeConfig(prodInst.KubeClusterName)
	if !ok {
		return fmt.Errorf("kubeclustername %s not exist", prodInst.KubeClusterName)
	}
	err, out := m.componentMgr.Excute(component.Name, "CREATE", conf, refer.Input, prodInst.InstId, refer.Role)
	if err == nil {
		return err
	}
	//写库
	obj := ComponentObj{
		ProductInstId:   prodInst.InstId,
		ComponentName:   refer.ComponentName,
		KubeClusterName: prodInst.KubeClusterName,
		Role:            refer.Role,
		Input:           refer.Input,
		Output:          out,
	}
	return m.db.Create(ComponentObj2Model(obj)).Error
}

func (m *Manager) NewProductInst(prodInst *ProductInst) error {
	//记录数据
	m.db.Create(ProductInst2Model(ProductInst))
	//创建产品对应的组件
	design := m.GetDesign(prodInst.ProductName, prodInst.DesignRevision)
	for _, refer := range design.ComponentRefers {
		if err := m.CreateComponent(refer); err != nil {
			return err
		}
	}
	//更新状态
	go func() {

	}()
	return true
}

func (m *Manager) UpdateProductInst(prodInst *ProductInst, d Design) {

}

//TODO
func (m *Manager) CanUpdateDesign(p *ProductInst, d Design) (bool, string) {
	d1 := m.GetDesign(p.ProductName, p.DesignRevision)

	for _, j := range d1.ComponentRefers {
		for _, k := range d.ComponentRefers {
			if j.ComponentName == k.ComponentName {
				err, _ := m.componentMgr.Excute(j.ComponentName, "UPDATE_CHECK", p.KubeClusterName, k.Input, p.InstId, j.Role)
				if err != nil {
					return false, fmt.Sprintf("component %s cannot update", j.ComponentName)
				}
			}
		}
	}
	return true, ""
}
