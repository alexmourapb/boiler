// Package entity contains all the entities of the project
//go:generate msgp -tests=false
package entity

import "time"

// User is the entity of the user
type User struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
