package product

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Manager struct {
	db *gorm.DB
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
	m.db.Create(Design2Model(m))
}

func (m *Manager) NewOrUpdateProductInst(prodInst *ProductInst) error {
	var exist ProductInstModel
	m.db.Where("inst_id = ?", prodInst.InstId).First(&exist)
	if exist.InstId == "" {
		m.db.Create(ProductInst2Model(ProductInst))
		return true
	} else {
		var objs []ComponentObjModel
		m.db.Where("product_id = ?", prodInst.InstId).Find(&objs)
		//TODO 1. check update 2.update or not

	}
}
