package internal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"test-project/internal/models"
)

var db *gorm.DB

func InitDatabase(dataBase *gorm.DB) {
	db = dataBase
	db.AutoMigrate(&internal.EntityFirst{})
	db.AutoMigrate(&internal.EntitySecond{})
	db.AutoMigrate(&internal.NestedEntitySecond{})
}

func FindEntitiesSecond() (measurements []*internal.EntitySecond) {
	if err := db.Find(&measurements).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindEntitiesFirst() (products []*internal.EntityFirst) {
	if err := db.Find(&products).Error; err != nil {
		fmt.Println(err)
	}
	return
}
