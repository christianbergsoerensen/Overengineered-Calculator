package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS calculations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			operation TEXT NOT NULL,
			a REAL NOT NULL,
			b REAL NOT NULL,
			result REAL NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
		)
	`)

	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil

}

func (s *SQLiteStorage) SaveCalculation(result CalculationResult) error {
	_, err := s.db.Exec(`
		INSERT INTO calculations (operation,a,b,result,timestamp)
		VALUES (?,?,?,?,?)`, result.Operation, result.A, result.B, result.Result, result.Timestamp,
	)

	return err
}

func (s *SQLiteStorage) GetHistory() ([]CalculationResult, error) {
	rows, err := s.db.Query(`
		SELECT * FROM calculations
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := []CalculationResult{}
	for rows.Next() {
		res := CalculationResult{}
		err := rows.Scan(&res.ID, &res.Operation, &res.A, &res.B, &res.Result, &res.Timestamp)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil

}
