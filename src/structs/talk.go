package structs

import (
	"time"

	"github.com/google/uuid"
)

// Talk defines a talk
type Talk struct {
	// Meta
	UUID uuid.UUID `pg:"type:uuid"`
	ID   int       `pg:",pk"`

	// Actual Talk information
	Title       string
	Description string
	Slides      string
	Nickname    string

	Date          time.Time // the actual date entered
	FormattedDate string    // the date formatted in a way that can be displayed nicely

	// Organization
	Upcoming bool `pg:",use_zero"`
}
