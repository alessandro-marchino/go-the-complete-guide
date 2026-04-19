package models

import (
	"time"

	"example.com/rest_api/db"
)

type Event struct {
	ID int64 `json:"id"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location string `json:"location" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`
	UserID int `json:"-"`
}

var events = []Event {}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
}

func GetAllEvents() []Event {
	return events
}
