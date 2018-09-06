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

func FindEntitiesFirst() (products []*internal.EntityFirst) {
	if err := db.Find(&products).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindEntitiesSecond() (measurements []*internal.EntitySecond) {
	if err := db.Find(&measurements).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func GetEntitySecond(idOne int, idTwo int) (entity internal.EntitySecond) {
	if err := db.First(&entity, idOne, idTwo).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindNestedEntitiesSecond(idOneSecond int, idTwoSecond int) (nestedEntities []*internal.NestedEntitySecond) {
	if err := db.Where("id_one_second = ? AND id_two_second = ?", idOneSecond, idTwoSecond).Find(&nestedEntities).Error; err != nil {
		fmt.Println(err)
	}
	return
}
