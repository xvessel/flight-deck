package product

import (
	"github.com/jinzhu/gorm"
	"gorm"
)

type ProductInst struct {
	InstId          string
	Name            string
	KubeClusterName string
	ProductName     string
	DesignRevision  string
	Input           map[string]string
}

type ProductInstModel struct {
	gorm.Model

	//productName+KubeClusterName+Name
	//KubeNamespace
	InstId          string
	Name            string
	KubeClusterName string
	ProductName     string
	DesignRevision  string
	Input           string
}

func ProductInst2Model(p *ProductInst) ProductInstModel {
	return ProductInstModel{
		InstId:          p.ProductName + "-" + p.KubeClusterName + "-" + p.Name,
		Name:            p.Name,
		KubeClusterName: p.KubeClusterName,
		ProductName:     p.ProductName,
		DesignRevision:  p.DesignRevision,
		Input:           p.Input,
	}
}

type ComponentObjModel struct {
	gorm.Model

	ProductInstId   string
	ComponentName   string
	KubeClusterName string
	Role            string //kubernetes name
}

type ProductModel struct {
	gorm.Model

	Name string
}
