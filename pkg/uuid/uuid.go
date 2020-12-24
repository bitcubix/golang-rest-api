package uuid

import "github.com/google/uuid"

//Generate returns a UUIDv4 as a string
func Generate() string {
	return uuid.New().String()
}
