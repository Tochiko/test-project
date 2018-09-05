package internal

type EntityFirst struct {
	IdOne   int `gorm:"primary_key"`
	IdTwo   int `gorm:"primary_key"`
	Value   string
	Seconds []*EntitySecond `gorm:"-"`
}

func (p *EntityFirst) TableName() string {
	return "entity_first"
}
