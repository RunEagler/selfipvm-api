package common

import (
	"github.com/jmoiron/sqlx"
	"runtime"
	"fmt"
	"database/sql"
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

const TimeFormat = "2006/01/02"


// Preparex creates a prepared statement for queries or executions.
func Preparex(db *sqlx.DB, query string) (*sqlx.Stmt, error) {
	if pc, _, _, ok := runtime.Caller(1); ok {
		fn := runtime.FuncForPC(pc)
		fileName, _ := runtime.FuncForPC(pc).FileLine(pc)
		return db.Preparex(fmt.Sprintf("-- function=%s filename=%s\n%s", fn.Name(), fileName, query))
	}

	return db.Preparex(query)
}

// PrepareNamed creates a prepared statement for queries or executions.
func PrepareNamed(db *sqlx.DB, query string) (*sqlx.NamedStmt, error) {
	if pc, _, _, ok := runtime.Caller(1); ok {
		fn := runtime.FuncForPC(pc)
		fileName, _ := runtime.FuncForPC(pc).FileLine(pc)
		return db.PrepareNamed(fmt.Sprintf("-- function=%s filename=%s\n%s", fn.Name(), fileName, query))
	}

	return db.PrepareNamed(query)
}

//ToPtr :convert string to *string
func ToPtr(s string) *string {
	return &s
}

func ToNoPtr(s *string) string {
	if (s == nil) {
		return ""
	} else {
		return *s
	}
}

func ToNullString(s *string) sql.NullString {
	var nullString sql.NullString
	if s == nil {
		nullString.String = ""
		nullString.Valid = false
	} else {
		nullString.String = *s
		nullString.Valid = true
	}
	return nullString
}

// LogInfo logs the message with the route.
func LogInfo(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogInfo(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}

// LogError logs the error with the route.
func LogError(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogError(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}


//ConvertStringToDate :convert string to date
func ConvertStringToDate(date string) (*time.Time,error){

	var dateTime time.Time
	var err error

	dateTime,err = time.Parse(TimeFormat,date)
	if err != nil{
		return nil,err
	}
	return &dateTime,err
}
