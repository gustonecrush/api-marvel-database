package models

type Hero struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	HeroName string `gorm:"type:varchar(300)" json:"hero_name"`
	Gender string `gorm:"type:varchar(1)" json:"gender"`
}