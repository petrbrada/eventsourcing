package core

import (
	"context"
	"errors"
)

// ErrSnapshotNotFound returned when no snapshot is found in the snapshot store
var ErrSnapshotNotFound = errors.New("snapshot not found")

// Snapshot holds current state of an aggregate
type Snapshot struct {
	ID            string
	Type          string
	Version       Version
	GlobalVersion Version
	State         []byte
}

// SnapshotStore interface expose the methods an snapshot store must uphold
type SnapshotStore interface {
	Save(id, aggregateType string, snapshot Snapshot) error
	Get(ctx context.Context, id, aggregateType string) (Snapshot, error)
}
