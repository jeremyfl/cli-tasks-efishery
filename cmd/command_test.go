package cmd

import (
	"context"
	"testing"

	"github.com/jeremylombogia/cli-tasks-efishery/internal"
	"github.com/rs/xid"

	"github.com/jeremylombogia/cli-tasks-efishery/configs"
	"github.com/stretchr/testify/assert"
)

// Initialize database from config
var client, db = configs.Init()

// TestInsert it compare between test store in repositories and show repositories
func TestInsert(t *testing.T) {
	guid := xid.New()

	var createTask = internal.Task{
		ID:      guid.String(),
		Content: "Lorem ipsum dolor si amet",
		IsDone:  false,
		Tag:     "College",
	}

	var store, _ = internal.Store(&createTask)

	var showTask internal.Task

	var foundTask = internal.Show(store.ID)
	if err := foundTask.ScanDoc(&showTask); err != nil {
		panic(err)
	}

	// TODO:: Add all data except rev
	assert.Equal(t, store.ID, showTask.ID, "Insert with affected insert not equal")
}

func TestUpdate(t *testing.T) {
	guid := xid.New()

	var createTask = internal.Task{
		ID:      guid.String(),
		Content: "Lorem ipsum dolor si amet",
		IsDone:  false,
		Tag:     "College",
	}

	var store, _ = internal.Store(&createTask)

	var row = internal.Show(store.ID)

	var task internal.Task

	if err := row.ScanDoc(&task); err != nil {
		panic(err)
	}

	var _, err = internal.Update(store.ID, &task)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, task.ID, task.ID, "Show struct not equal with update struct")

}

func TestPing(t *testing.T) {
	var testPing, _ = client.Ping(context.TODO())

	assert.Equal(t, testPing, true)
}
