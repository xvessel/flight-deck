package product

import (
	"github.com/jinzhu/gorm"
	"gorm"
)

type ProductInst struct {
	InstId          string //namespace:productName+KubeClusterName+Name
	Name            string
	KubeClusterName string
	ProductName     string
	DesignRevision  string
}

type ProductInstModel struct {
	gorm.Model

	InstId          string //namespace:productName+KubeClusterName+Name
	Name            string
	KubeClusterName string
	ProductName     string
	DesignRevision  string

	Status string
}

type ComponentObj struct {
	ProductInstId   string
	ComponentName   string
	KubeClusterName string
	Role            string            //kubernetes name
	Input           map[string]string //map[string]string
	Output          map[string]string //map[string]string
}

type ComponentObjModel struct {
	gorm.Model

	ProductInstId   string
	ComponentName   string
	KubeClusterName string
	Role            string //kubernetes name
	InputJson       string //map[string]string
	OutputJson      string //map[string]string
}

func ComponentObj2Model(c ComponentObj) ComponentObjModel {
	inputByte, _ := json.Marshal(c.Input)
	outputByte, _ := json.Marshal(c.Output)
	return ComponentObjModel{
		ProductInstId:   model.ProductInstId,
		ComponentName:   model.ComponentName,
		KubeClusterName: model.KubeClusterName,
		Role:            model.Role,
		Input:           string(inputByte),
		Output:          string(outputByte),
	}
}

type ProductModel struct {
	gorm.Model

	Name string
}

func ProductInst2Model(p *ProductInst) ProductInstModel {
	return ProductInstModel{
		InstId:          p.ProductName + "-" + p.KubeClusterName + "-" + p.Name,
		Name:            p.Name,
		KubeClusterName: p.KubeClusterName,
		ProductName:     p.ProductName,
		DesignRevision:  p.DesignRevision,
		Status:          "init",
	}
}
