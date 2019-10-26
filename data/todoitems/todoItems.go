package todoitems

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/prologic/bitcask"
)

// TodoItem encapsulates all the information needed to persist a todo entry.
type TodoItem struct {
	ID      uuid.UUID
	Message string
	Done    bool
	Created time.Time
}

var db *bitcask.Bitcask

// Open opens the connection to the local database.
func Open() error {
	bitcask.WithSync(true)
	d, err := bitcask.Open("tmp/db")
	db = d
	return err
}

// Close closes the connection to the local database.
func Close() {
	db.Close()
}

// Add adds a TodoItem to the local database.
func Add(todo TodoItem) error {
	key, err := todo.ID.MarshalText()
	ser, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return db.Put(key, []byte(ser))
}

// List returns a channel that streams TodoItems from the database. It also returns
// an error channel that streams any errors that come about.
func List(getAll bool) (<-chan TodoItem, <-chan error, <-chan byte) {
	items := make(chan TodoItem)
	errs := make(chan error)
	done := make(chan byte)
	keys := db.Keys()

	go func() {
		for {
			key, more := <-keys
			if !more {
				break
			}

			ser, err := db.Get(key)

			if err != nil {
				errs <- err
			}

			var tdi TodoItem
			jErr := json.Unmarshal(ser, &tdi)

			if jErr != nil {
				errs <- err
			}

			if getAll || !tdi.Done {
				items <- tdi
			}
		}

		done <- 0
	}()

	return items, errs, done
}

// Done accepts a slice of string ids that map to a stored todo item, and sets
// the "Done" state to true.
// All changed TodoItems are streamed back through the TodoItem channel, and any
// errors that come about are streamed back through the error channel.
func Done(ids []string) (<-chan TodoItem, <-chan error, <-chan byte) {
	changedItems := make(chan TodoItem)
	errs := make(chan error)
	done := make(chan byte)

	go func() {
		for _, id := range ids {
			err := db.Scan([]byte(id), func(key []byte) error {
				ser, err := db.Get(key)

				if err != nil {
					return err
				}

				var item TodoItem
				json.Unmarshal(ser, &item)
				item.Done = true
				ser, err = json.Marshal(item)

				if err != nil {
					return err
				}

				err = db.Put(key, ser)

				if err != nil {
					return err
				}

				changedItems <- item

				return nil
			})

			if err != nil {
				errs <- err
			}
		}

		done <- 0
	}()

	return changedItems, errs, done
}

// Remove removes a TodoItem from the database.
func Remove(ids []string) (<-chan string, <-chan error, <-chan byte) {
	dKeys := make(chan string)
	errs := make(chan error)
	done := make(chan byte)

	go func() {
		for _, id := range ids {
			err := db.Scan([]byte(id), func(key []byte) error {
				dErr := db.Delete(key)

				if dErr != nil {
					return dErr
				}

				dKeys <- string(key)
				return nil
			})

			if err != nil {
				errs <- err
			}
		}

		done <- 0
	}()

	return dKeys, errs, done
}
