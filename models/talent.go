package models

type Talent struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(100), unique" json:"code"`
	Nama string `gorm:"type:varchar(250)" json:"nama"`
}
