/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:55
 * Filename      : design.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"encoding/json"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Product struct {
	Name string
}

type ProductModel struct {
	gorm.Model
	Name string `gorm:"UNIQUE_INDEX"`
}

type Design struct {
	ProductName     string           `json:"ProductName"`
	Revision        string           `json:"Revision"`
	ComponentRefers []ComponentRefer `json:"ComponentRefers"`
}

type ComponentRefer struct {
	Role          string            `json:"Role"`
	ComponentName string            `json:"ComponentName"`
	Input         map[string]string `json:"Input"`
	PreRole       []string          `json:"PreRole"`
}

type DesignModel struct {
	gorm.Model

	ProductName        string `gorm:"UNIQUE_INDEX:design_id"`
	Revision           string `gorm:"UNIQUE_INDEX:design_id"`
	ComponentRefersStr string
}

func EqualComponentRefer(a, b ComponentRefer) bool {
	if a.Role != b.Role || a.ComponentName != b.ComponentName {
		return false
	}
	return reflect.DeepEqual(a.Input, b.Input)
}

func Design2Model(d *Design) DesignModel {
	var ret DesignModel
	ret.ProductName = d.ProductName

	if d.Revision == "" {
		d.Revision = uuid.NewV4().String()
	}
	ret.Revision = d.Revision
	b, _ := json.Marshal(d.ComponentRefers)
	ret.ComponentRefersStr = string(b)
	return ret
}

func DesignModel2Desgin(d DesignModel) Design {
	var ret Design
	ret.ProductName = d.ProductName
	ret.Revision = d.Revision
	ret.ComponentRefers = make([]ComponentRefer, 0)
	json.Unmarshal([]byte(d.ComponentRefersStr), &ret.ComponentRefers)
	return ret
}
