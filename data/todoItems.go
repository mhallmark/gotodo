package data

import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/prologic/bitcask"
)

type TodoItem struct {
	Id uuid.UUID
	Message string
	Done bool
	Created time.Time
}

var db *bitcask.Bitcask

func Open() {
	db, _ = bitcask.Open("tmp/db")
}

func Close() {
	db.Close()
}

func Add(todo TodoItem) error {
	var key []byte
	todo.Id.UnmarshalText(key)
	ser, err := json.Marshal(todo)

	if (err != nil) {
		return err
	}

	return db.Put(key, []byte(ser))
}

func List() (<-chan TodoItem, error) {
	items := make(chan TodoItem)
	go func() {
		for k := range db.Keys() {
			ser, err := db.Get(k)

			if (err != nil) {
				fmt.Println(err)
				continue
			}

			var tdi TodoItem
			jErr := json.Unmarshal(ser, tdi)

			if (jErr != nil) {
				fmt.Println(jErr)
				fmt.Println(ser)
				fmt.Println(k)
				continue
			}

			items<- tdi
		}

		close(items)
	}()

	return items, nil
}

func Remove(key string) error {
	return db.Delete([]byte(key))
}