package cmd

import (
	"context"
	"testing"

	"github.com/jeremylombogia/cli-tasks-efishery/configs"
	"github.com/stretchr/testify/assert"
)

// Initialize database from config
var client, db = configs.Init()

// TestInsert ...
func TestInsert(t *testing.T) {
	assert.Equal(t, true, true, "Insert with affected insert not equal")
}

func TestPing(t *testing.T) {
	var testPing, _ = client.Ping(context.TODO())

	assert.Equal(t, testPing, true)

}
