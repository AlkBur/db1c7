package test

import (
	"github.com/AlkBur/db1c7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB1(t *testing.T)  {
	db, err := db1c7.Open("files/db1")
	assert.NoError(t, err, "Open DB")

	t.Log(db)
}
