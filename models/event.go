package models

import (
	"database/sql"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Title       string    `bind:"required"`
	Description string    `bind:"required"`
	DateTime    time.Time `bind:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// later: implement logic to save the event to a database
	query := "INSERT INTO events (title, description, datetime, user_id) VALUES (?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(e.Title, e.Description, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	// later: implement logic to retrieve all events from a database
	query := "SELECT id, title, description, datetime, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.DateTime, &e.UserID); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT id, title, description, datetime, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	if err := row.Scan(&e.ID, &e.Title, &e.Description, &e.DateTime, &e.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No event found
		}
		return nil, err // Other error
	}
	return &e, nil
}

func (e Event) UpdateEvent() error {
	query := "UPDATE events SET title = ?, description = ?, datetime = ?, user_id = ? WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Description, e.DateTime, e.UserID, e.ID)
	return err
}

func (e Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
func (e Event) Unregister(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
