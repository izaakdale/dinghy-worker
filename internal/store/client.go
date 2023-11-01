package store

import (
	"github.com/dgraph-io/badger/v2"
)

type Client struct {
	*badger.DB
}

func (c *Client) Insert(k, v []byte) error {
	txn := c.NewTransaction(true)
	err := txn.Set(k, v)
	if err != nil {
		txn.Discard()
		return err
	}
	return txn.Commit()
}

func (c *Client) Fetch(k []byte) ([]byte, error) {
	txn := c.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(k)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	// Copy the value to a new byte slice
	value, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *Client) Delete(k []byte) error {
	txn := c.NewTransaction(true)
	err := txn.Delete(k)
	if err != nil {
		txn.Discard()
		return err
	}
	return txn.Commit()
}
