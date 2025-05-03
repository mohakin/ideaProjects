package models

import (
	"time"

	"example.com/REST_API/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`

	// Something to look out for in the future -
	// - (?, ?, ?, ?, ?) this is how it should be. When first writting this code, I had an extra "," at the end that did not get caught by -
	// gemini, the editor or me. I need to be vary careful about this in the future as it lead to me thorwing a 500 response error -
	// - it took me going back through the lessons one by one and seeing if there was any typos. I was also missing 2 different -
	// - retruns in this code but nothing that would have caused the error as all the did was not kill the process after the error -
	// - occured. This took me literally hours of looking over the problem mutliple times to find it. Don't let this happen again.

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (event Event) Update() error {
	query := `
UPDATE events
SET name = ?, description = ?, location = ?, dateTime = ?
WHERE id = ?
`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registration(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

// If you have a query that changes stuff it's ".Exec", if you have a query that fetches something it's a ".Query".
// "Prepare()" prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times -
// - (potentially with different data for it's placeholders).
// This is only true if the prepared statement is not closed(stmt.Close()) in between those executions. In that case there wouldn't be any advantage.
//
