package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountData(t *testing.T) {
	a := AccountData{}
	assert.Empty(t, a.Password)
	assert.Empty(t, a.Id)
}
