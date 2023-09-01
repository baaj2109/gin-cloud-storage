package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ConnectTest(t *testing.T) {
	_, err := NewDrive()
	assert.Equal(t, err, nil)
}
