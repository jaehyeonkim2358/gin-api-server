// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/jaehyeonkim2358/gin-api-server/ent/board"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Board is the model entity for the Board schema.
type Board struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Board) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case board.FieldID:
			values[i] = new(sql.NullInt64)
		case board.FieldTitle, board.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Board fields.
func (b *Board) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case board.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case board.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case board.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				b.Content = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Board.
// This includes values selected through modifiers, order, etc.
func (b *Board) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Board.
// Note that you need to call Board.Unwrap() before calling this method if this Board
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Board) Update() *BoardUpdateOne {
	return NewBoardClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Board entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Board) Unwrap() *Board {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Board is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Board) String() string {
	var builder strings.Builder
	builder.WriteString("Board(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(b.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Boards is a parsable slice of Board.
type Boards []*Board
