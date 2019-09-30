/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:55
 * Filename      : manager_design.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"fmt"
)

func (m *Manager) NewDesign(d *Design) error {
	_, err := m.GetProduct(d.ProductName)
	if err != nil {
		return fmt.Errorf("GetProduct %v %v", *d, err)
	}
	for _, j := range d.ComponentRefers {
		spec, err := m.componentMgr.Spec(j.ComponentName)
		if err != nil {
			return fmt.Errorf("Component %v not found ", j.ComponentName)
		}
		for _, in := range spec.Inputs {
			if _, ok := j.Input[in.Name]; !ok {
				return fmt.Errorf("input %v not exist", in.Name)
			}
		}
	}
	//ComponentRefer validate
	sort := make([]ComponentRefer, 0)
	overRole := make(map[string]interface{})
	for len(d.ComponentRefers) != 0 {
		found := false
		for i, j := range d.ComponentRefers {
			if _, ok := overRole[j.Role]; ok {
				return fmt.Errorf("role %s twice", j.Role)
			}
			preok := true
			for _, r := range j.PreRole {
				if _, ok := overRole[r]; !ok {
					preok = false
				}
			}
			if preok {
				sort = append(sort, j)
				overRole[j.Role] = nil
				d.ComponentRefers = append(d.ComponentRefers[:i], d.ComponentRefers[i+1:]...)
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("circle depends")
		}
	}
	d.ComponentRefers = sort
	dm := Design2Model(d)
	return m.db.Create(&dm).Error
}

func (m *Manager) GetDesign(productName string, revision string) (Design, error) {
	var model DesignModel
	err := m.db.Where("product_name = ? and revision = ?", productName, revision).First(&model).Error
	return DesignModel2Desgin(model), err
}

func (m *Manager) GetDesigns(productName string) ([]Design, error) {
	model := make([]DesignModel, 0)
	err := m.db.Where("product_name = ? ", productName).Find(&model).Error

	ret := make([]Design, 0)
	for _, j := range model {
		ret = append(ret, DesignModel2Desgin(j))
	}
	return ret, err
}

func (m *Manager) NewProduct(prodName string) error {
	return m.db.Create(&ProductModel{Name: prodName}).Error
}

func (m *Manager) Products() ([]ProductModel, error) {
	models := make([]ProductModel, 0)
	err := m.db.Find(&models).Error
	return models, err
}

func (m *Manager) GetProduct(prodName string) (ProductModel, error) {
	var mod ProductModel
	err := m.db.Where("name=?", prodName).First(&mod).Error
	return mod, err
}
