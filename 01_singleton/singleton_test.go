package singleton_test

import (
	singleton "github.com/elpsyr/go-patterns/01_singleton"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}
