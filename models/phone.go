package models

type Phone struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Nama     string `gorm:"type:varchar(128)" json:"nama"`
	Ram      int64  `gorm:"primaryKey" json:"ram"`
	Internal int64  `gorm:"primaryKey" json:"internal"`
}
