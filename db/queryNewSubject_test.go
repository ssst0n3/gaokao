package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryNewSubject(t *testing.T) {
	assert.NoError(t, QueryNewSubject())
}
