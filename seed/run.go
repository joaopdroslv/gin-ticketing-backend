package seed

import "database/sql"

func Run(db *sql.DB) error {

	if err := Users(db, 100); err != nil {
		return err
	}

	// Other seeders come here...

	return nil
}
