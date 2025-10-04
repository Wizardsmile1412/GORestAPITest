package models

import "example.com/restAPIProject/db"

type Registration struct {
	ID      int64
	EventID int64
	UserID  int64
}

func GetAllRegistrations() ([]Registration, error) {
	query := `SELECT * FROM registrations`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations []Registration

	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID)

		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}
