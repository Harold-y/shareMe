package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DB *leveldb.DB // Global variable of DB
)

func StartDB(name string) *leveldb.DB {
	DB, err := leveldb.OpenFile(name+".db", nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return DB
}
