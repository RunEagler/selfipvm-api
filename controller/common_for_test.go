package controller

import (
	"github.com/goadesign/goa"
	"testing"
	"github.com/jmoiron/sqlx"
	"selfipvm-api/common"
	"fmt"
	"os"
)

var (
	service = goa.New("")
	postgresqlx *sqlx.DB
)

func TestMain(m *testing.M){

	var err error
	postgresqlx,err = common.ConnectPostgres("taileagler","selfipvm")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	result := m.Run()
	os.Exit(result)
}