package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	models "github.com/my/repo/Models"
	//"github.com/my/repo/models"
)

var DB *pg.DB

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "1234",
		Addr:     "localhost:5432",
		Database: "bank",
	}
	DB = pg.Connect(opts)
	if DB == nil {
		fmt.Println("Failed to connect to db")
		os.Exit(100)
	}
	fmt.Println("Connection to db successful")
	createSchema()

	return DB
}

func createSchema() error {
	models := []interface{}{

		(*models.Bank)(nil),
		(*models.Branch)(nil),
		(*models.Account)(nil),
		(*models.Customer)(nil),
		(*models.Transaction)(nil),
		(*models.Map)(nil),
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			fmt.Printf("table not created %v\n", err)
			return err
		}
		log.Println("Tables succesffuly created ")
	}
	return nil
}
