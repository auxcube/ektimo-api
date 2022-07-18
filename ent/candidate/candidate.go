// Code generated by entc, DO NOT EDIT.

package candidate

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the candidate type in the database.
	Label = "candidate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the candidate in the database.
	Table = "candidates"
)

// Columns holds all SQL columns for candidate fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusNoQuizAssigned is the default value of the Status enum.
const DefaultStatus = StatusNoQuizAssigned

// Status values.
const (
	StatusNoQuizAssigned Status = "no-quiz-assigned"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusNoQuizAssigned:
		return nil
	default:
		return fmt.Errorf("candidate: invalid enum value for status field: %q", s)
	}
}