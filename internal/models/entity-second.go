package internal

type EntitySecond struct {
	IdOne          int `gorm:"primary_key"`
	IdTwo          int `gorm:"primary_key"`
	Value          string
	NestedEntities []*NestedEntitySecond `gorm:"-"`
}

func (m *EntitySecond) TableName() string {
	return "entity_second"
}

type NestedEntitySecond struct {
	IdOneSecond int `gorm:"primary_key"`
	IdTwoSecond int `gorm:"primary_key"`
	IdThree     int `gorm:"primary_key"`
	ValueOne    float64
	ValueTwo    float64
}

func (mp *NestedEntitySecond) TableName() string {
	return "nested_entity_second"
}
