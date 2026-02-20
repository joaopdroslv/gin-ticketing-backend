package seed

import (
	"database/sql"
	"go-gin-ticketing-backend/internal/shared/enums"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func Users(db *sql.DB, amount int) error {

	query := `
		INSERT INTO users (
			user_status_id,
			name,
			email,
			birthdate
		) VALUES (?, ?, ?, ?)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < amount; i++ {

		birthdate := gofakeit.DateRange(
			time.Now().AddDate(-80, 0, 0),
			time.Now().AddDate(-18, 0, 0),
		)

		if _, err := stmt.Exec(
			int64(enums.PasswordCreation),
			gofakeit.Name(),
			gofakeit.Email(),
			birthdate.Format("2006-01-02"),
		); err != nil {
			return err
		}
	}

	return nil
}
