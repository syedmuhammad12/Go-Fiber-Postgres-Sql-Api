package models

import "gorm.io/gorm"

type TestVars struct {
	Variable *string `gorm:"primary key;autoIncrement" json:"variable"`
	Value    uint    `json:"value"`
}

func SetVar(db *gorm.DB) error {
	err := db.AutoMigrate(&TestVars{})
	return err
}
