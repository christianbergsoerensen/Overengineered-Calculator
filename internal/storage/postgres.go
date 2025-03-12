package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQLStorage struct {
	db *sql.DB
}

func NewPostgreSQLStorage(dbURL string) (*PostgreSQLStorage, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS calculations (
			id SERIAL PRIMARY KEY,
			operation TEXT NOT NULL,
			a DOUBLE PRECISION NOT NULL,
			b DOUBLE PRECISION NOT NULL,
			result DOUBLE PRECISION NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
		)
	`)

	if err != nil {
		return nil, err
	}

	return &PostgreSQLStorage{db: db}, nil
}

func (s *PostgreSQLStorage) SaveCalculation(result CalculationResult) error {
	_, err := s.db.Exec(`
		INSERT INTO calculations (operation, a, b, result, timestamp)
		VALUES ($1, $2, $3, $4, $5)
	`, result.Operation, result.A, result.B, result.Result, result.Timestamp)

	return err
}

func (s *PostgreSQLStorage) GetHistory() ([]CalculationResult, error) {
	rows, err := s.db.Query(`
		SELECT id, operation, a, b, result, timestamp FROM calculations
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
