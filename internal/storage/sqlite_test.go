package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHistory(t *testing.T) {
	//Using in-memory SQLite
	store, err := NewSQLiteStorage(":memory:")
	if err != nil {
		t.Errorf("recieved error %v, when trying to open database", err)
	}

	//When 0 entries
	history, err := store.GetHistory()
	if err != nil {
		t.Errorf("recieved error %v, when trying to get history", err)
	}
	assert.Equal(t, len(history), 0)

	//Fake calc result
	testCalcRes := CalculationResult{
		Operation: "multiply",
		A:         4,
		B:         10,
		Result:    14,
		Timestamp: time.Now(),
	}

	//testCalcRes is passed by value, so it cannot be modified in SaveCalculation
	err = store.SaveCalculation(testCalcRes)
	if err != nil {
		t.Errorf("recieved error %v, when trying to save to database", err)
	}

	history, err = store.GetHistory()
	if err != nil {
		t.Errorf("recieved error %v, when trying to get history", err)
	}
	//Want to ensure that only 1 entry is retrieved since only 1 entry is in the database
	assert.Equal(t, len(history), 1)
	assert.Equal(t, testCalcRes.Operation, history[0].Operation)
	assert.Equal(t, testCalcRes.A, history[0].A)
	assert.Equal(t, testCalcRes.B, history[0].B)
	assert.Equal(t, testCalcRes.Result, history[0].Result)
}
