package repository

import (
	"github.com/jmoiron/sqlx"
	"selfipvm-api/models"
	"selfipvm-api/common"
	"fmt"
)

//ActivityRepository :repository for activity table
type ActivityRepository interface {
	InsertActivity(activity *models.Activity) error
}

type activityRepository struct {
	Stmts      map[activityKey]*sqlx.Stmt
	NamedStmts map[activityKey]*sqlx.NamedStmt
}

type activityKey int

const (
	insertActivityKey activityKey = iota
)

const (
	insertActivityQuery = `
		INSERT INTO
			activity(content,minutes,date,type,created_at,updated_at)
		VALUES(
			:content,:minutes,:date,:type,current_timestamp,current_timestamp
		)
	`
)

//NewActivityRepository :new ActivityRepository
func NewActivityRepository(postgresqlx *sqlx.DB) ActivityRepository {

	var repo *activityRepository
	var err error
	repo = &activityRepository{
		Stmts:      make(map[activityKey]*sqlx.Stmt),
		NamedStmts: make(map[activityKey]*sqlx.NamedStmt),
	}
	repo.NamedStmts[insertActivityKey],err = common.PrepareNamed(postgresqlx, insertActivityQuery)
	if err != nil{
		fmt.Println(err)
	}
	return repo
}

func (r *activityRepository) InsertActivity(activity *models.Activity) (err error) {

	_, err = r.NamedStmts[insertActivityKey].Exec(activity)
	if err != nil {
		return err
	}
	return err
}
