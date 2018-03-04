package models

import (
	"database/sql"
	"time"
)

type Activity struct {
	ActivityID int            `db:"activity_id"`
	Minutes    int            `db:"minutes"`
	Content    sql.NullString `db:"content"`
	Date       time.Time      `db:"date"`
	Type       int            `db:"type"`
}
