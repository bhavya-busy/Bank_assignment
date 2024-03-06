package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	db "github.com/my/repo/DB"
	routes "github.com/my/repo/Routes"
)

var dbRef *pg.DB

func main() {
	fmt.Println("Bank systemmm")
	dbRef = db.Connect()

	router := gin.Default()

	routes.BankRoutes(router)
	routes.BranchRoutes(router)
	routes.AccountRoutes(router)
	routes.CustomerRoutes(router)
	routes.MapRoutes(router)
	routes.TransactionRoutes(router)

	router.Run(":3000")

}
