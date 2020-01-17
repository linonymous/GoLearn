package main

import (
	"github.com/linonymous/GoLearn/mockntest/custom_mocking/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockDB struct {
	db.DML
}

func (m *MockDB) ListTable() []string  {
	return []string{"abc", "mno", "pqr"}
}
func TestGetTables(t *testing.T) {
	d := &MockDB{}
	a := GetTables(d)
	assert.Equal(t, a[0], "abc")
	assert.Equal(t, a[1], "mno")
	assert.Equal(t, a[2], "pqr")
}