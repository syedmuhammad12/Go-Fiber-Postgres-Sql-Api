package storage

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type GetVariable struct {
	Variable *string `json:"variable"`
	Value    uint    `json:"value"`
}

type TestVars struct {
	Variable *string `gorm:"primary key;autoIncrement" json:"variable"`
	Value    uint    `json:"value"`
}

func SetVar(db *gorm.DB) error {
	err := db.AutoMigrate(&TestVars{})
	return err
}

// func NewConnection(config *Config) (*gorm.DB, error) {
func NewConnection() (*gorm.DB, error) {
	// dsn := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	// 	config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	// )
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(postgres.Open("postgres://root:secret@host.docker.internal:5432/simple_bank?sslmode=disable"))

	if err != nil {
		log.Fatal("could not migrate db")
		return db, err
	}

	err_a := db.AutoMigrate(&TestVars{})
	if err_a != nil {
		log.Fatal("could not migrate db")
	}
	var lis []GetVariable
	db.Raw(`select * from "test_vars"`).Scan(&lis)

	if len(lis) == 0 {
		db.Exec(`insert into "test_vars" ("variable", "value") values ('a', 0), ('b', 0)`)
	}

	// err_1 := SetVar(db)
	// if err_1 != nil {
	// 	log.Fatal("could not migrate db")
	// }

	// var lis []GetVariable
	// // db.Exec(`CREATE TABLE IF NOT EXISTS "test_vars"("variable", "")`)
	// db.Raw(`select * from "test_vars"`).Scan(&lis)
	// fmt.Println(lis[0].Value + lis[1].Value)

	// if len(lis) == 0 {
	// 	db.Exec(`insert into "test_vars" ("variable", "value") values ('a', 0), ('b', 0)`)
	// }

	return db, nil
}
