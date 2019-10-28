package main

import (
	"encoding/json"
	"fmt"

	"github.com/tecbot/gorocksdb"
)

type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

// func openDb() {
// 	opts := gorocksdb.NewDefaultOptions()
// 	opts.SetCreateIfMissing(true)
// 	db, err := gorocksdb.OpenDb(opts, "test2.db")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//db.Close()
// }

func main() {

	person := Person{Name: "David", Lastname: "Beckham"}

	//openDb()
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "test2.db")
	if err != nil {
		fmt.Println(err)
	}
	//ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()

	coveculjak, err := json.Marshal(person)

	if err != nil {
		fmt.Println(err)
		return
	}

	//db.Put(person)
	db.Put(wo, []byte("nekiCovek"), []byte(coveculjak))
	//db.Delete(wo, []byte("nekiCovek"))
	//fmt.Println(person)
	db.Close()

}
