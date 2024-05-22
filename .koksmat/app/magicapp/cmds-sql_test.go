package magicapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSQL(t *testing.T) {
	sql := `SELECT 1`
	x, err := ExecuteSQL(sql)
	if err != nil {
		t.Errorf("ExecuteSQL() error = %v", err)
	}

	assert.Equal(t, int64(1), x, "ExecuteSQL() = %v, want %v", x, 1)

}

func TestSelectSQL(t *testing.T) {
	sql := `SELECT 1 AS col1`
	x, err := SelectSQL(sql)
	if err != nil {
		t.Errorf("ExecuteSQL() error = %v", err)
	}

	assert.Equal(t, `[{"col1":1}]`, *x, "%s  ExecuteSQL() = %v, want %v", sql, x, 1)

}
