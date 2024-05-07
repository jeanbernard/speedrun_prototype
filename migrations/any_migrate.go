package main

import (
	database "developer/any/db"
	dbmodels "developer/any/dbmodels/models"
	"fmt"
)

func main() {
	db := database.NewSQLiteDatabase()
	err := anyMigrate(db)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// TODO: find another way to add models
func anyMigrate(db database.Database) error {
	//delete tables every time for now
	if err := db.GetDb().Migrator().DropTable(&dbmodels.Run{}, &dbmodels.Game{}, &dbmodels.Category{}); err != nil {
		panic("failed dropping tables")
	}

	if err := db.GetDb().AutoMigrate(&dbmodels.Run{}, &dbmodels.Game{}, &dbmodels.Category{}); err != nil {
		panic("failed migration")
	}
	return nil
}
