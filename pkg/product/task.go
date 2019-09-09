package product

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type TaskModel struct {
	ProductInstModel

	Status string
}
