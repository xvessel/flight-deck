/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:55
 * Filename      : instance.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

/**       |-----------------------v
init--->ready---->update------>delete----->deleted
          ^---------|
**/

var StatusInit string = "init"
var StatusReady string = "ready"
var StatusUpdate string = "update"
var StatusDelete string = "delete"
var StatusDeleted string = "deleted"

type ProductInst struct {
	Name            string
	ProductName     string
	DesignRevision  string
	KubeClusterName string
	//	InstId          string //namespace:productName+KubeClusterName+Name
}

type ProductInstModel struct {
	gorm.Model

	InstId          string `gorm:"UNIQUE_INDEX:inst_id"` //namespace:productName+KubeClusterName+Name
	ProductName     string `gorm:"UNIQUE_INDEX:inst"`
	Name            string `gorm:"UNIQUE_INDEX:inst"`
	KubeClusterName string
	DesignRevision  string

	Status string
}

type ComponentObj struct {
	ProductInstId   string
	Role            string //kubernetes name
	ComponentName   string
	KubeClusterName string
	DesignRevision  string
	Input           map[string]string //map[string]string
	Output          map[string]string //map[string]string
}

type ComponentObjModel struct {
	gorm.Model

	ProductInstId   string `gorm:"UNIQUE_INDEX:obj_id"`
	Role            string `gorm:"UNIQUE_INDEX:obj_id"`
	ComponentName   string
	KubeClusterName string
	DesignRevision  string
	InputJson       string //map[string]string
	OutputJson      string //map[string]string
	Status          string //kubernetes name
}

func ComponentObj2Model(c ComponentObj) ComponentObjModel {
	inputByte, _ := json.Marshal(c.Input)
	outputByte, _ := json.Marshal(c.Output)
	return ComponentObjModel{
		ProductInstId:   c.ProductInstId,
		ComponentName:   c.ComponentName,
		KubeClusterName: c.KubeClusterName,
		Role:            c.Role,
		InputJson:       string(inputByte),
		OutputJson:      string(outputByte),
	}
}

func ProductInst2Model(p *ProductInst) ProductInstModel {
	return ProductInstModel{
		InstId:          p.ProductName + "-" + p.KubeClusterName + "-" + p.Name,
		Name:            p.Name,
		KubeClusterName: p.KubeClusterName,
		ProductName:     p.ProductName,
		DesignRevision:  p.DesignRevision,
		Status:          StatusInit,
	}
}
func ProductInstModel2Inst(p ProductInstModel) ProductInst {
	return ProductInst{
		Name:            p.Name,
		KubeClusterName: p.KubeClusterName,
		ProductName:     p.ProductName,
		DesignRevision:  p.DesignRevision,
	}
}
