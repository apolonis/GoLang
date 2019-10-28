package main

import (
	// "github.com/DanielMorsing/rocksdb"
	"github.com/tecbot/gorocksdb"
	// _ "github.com/cockroachdb/c-rocksdb"
)

func main() {
	// opts := NewDefaultOptions()
	// opts.SetCreateIfMissing(true)

	// transactionDBOpts := NewDefaultTransactionDBOptions()

	// db, err := OpenTransactionDb(opts, transactionDBOpts, "test-db-1")
	// ensure.Nil(t, err)
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	gorocksdb.OpenDb(opts, "test.db")

}
