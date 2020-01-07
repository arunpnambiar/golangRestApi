package models

type user struct {
	uniqueid  uint64 `gorm:"primary_key;autoincrement" json:"userid"`
	firstname string `gorm:"size:50;notnull" json:"firstname"`
	lastname  string `gorm:"size:50;notnull" json:"lastname"`
	username  string `gorm:"size:50;notnull;unique" json:"lastname"`
	password  string `gorm:"size:200;notnull" json:"password"`
}
