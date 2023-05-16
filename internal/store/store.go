package store

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/pkg/errors"
)

var (
	_ Transactioner = (*Client)(nil)

	ErrNotFound = errors.New("no record matches that key") //TODO add return where applicable
)

type Transactioner interface {
	Insert(k, v []byte) error
	Fetch(k []byte) ([]byte, error)
	Delete(k []byte) error
}

func New(db *badger.DB) (*Client, error) {
	return &Client{DB: db}, nil
}
