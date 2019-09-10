package product

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Design struct {
	ProductName     string
	Revision        string
	ComponentRefers []ComponentRefer
}

type ComponentRefer struct {
	Role          string
	ComponentName string
	Input         map[string]string
	PreRole       []string
}

type DesignModel struct {
	gorm.Model

	ProductName        string
	Revision           string `gorm:"UNIQUE_INDEX"`
	ComponentRefersStr string
}

//TODO
func EqualComponentRefer(a, b ComponentRefer) bool {
}

//TODO
func CanUpdateDesign(a, b Design) bool {
	aComps := make(map[string]ComponentRefer)
	bComps := make(map[string]ComponentRefer)

}

func Design2Model(d *Design) DesignModel {
	var ret DesignModel
	ret.ProductName = d.ProductName

	if d.Revision == "" {
		d.Revision == uuid.Must(uuid.NewV4())
	}
	ret.Revision = d.Revision
	b, _ := json.Marshal(d.ComponentRefers)
	ret.ComponentRefersStr = string(b)
}

func DesignModel2Desgin(d *DesignModel) Design {
	var ret Design
	ret.ProductName = d.ProductName
	ret.Revision = d.Revision
	json.Unmarshal(d.ComponentRefersStr, &ret)
	return ret
}