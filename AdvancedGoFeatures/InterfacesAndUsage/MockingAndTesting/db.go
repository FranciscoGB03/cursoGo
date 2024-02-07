package db

type Database interface {
	Query(string) ([]string, error)
}

func GetDataFromDatabase(db Database) ([]string, error) {
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return results, nil
}
