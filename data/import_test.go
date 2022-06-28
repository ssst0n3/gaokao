package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRead(t *testing.T) {
	err := Read("A安徽2021.xlsx")
	assert.NoError(t, err)
}
