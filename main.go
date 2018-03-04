//go:generate goagen bootstrap -d protas-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"selfipvm-api/controller"
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"selfipvm-api/repository"
	_ "github.com/lib/pq"
	"selfipvm-api/common"
)

var db *sqlx.DB

func main() {
	// Create service
	service := goa.New("protas")

	//Initilize
	db = initializePostgres(service, "taileagler", "selfipvm")

	//NewRepository
	activityRepository := repository.NewActivityRepository(db)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	controller.CommonController(service, activityRepository)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func initializePostgres(service *goa.Service, user string, dbName string) *sqlx.DB {
	db, err := common.ConnectPostgres(user,dbName)
	if err != nil {
		service.LogError("initialize postgresql", "err", err)
		os.Exit(1)
	}
	return db
}
