// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ktakenaka/gosample2021/ent/sample"
)

// Sample is the model entity for the Sample schema.
type Sample struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// OfficeID holds the value of the "office_id" field.
	OfficeID int `json:"office_id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Size holds the value of the "size" field.
	Size sample.Size `json:"size,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo *string `json:"memo,omitempty"`
	// URL holds the value of the "url" field.
	URL *url.URL `json:"url,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	office_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sample) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sample.FieldURL:
			values[i] = new([]byte)
		case sample.FieldActive:
			values[i] = new(sql.NullBool)
		case sample.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case sample.FieldOfficeID:
			values[i] = new(sql.NullInt64)
		case sample.FieldID, sample.FieldCode, sample.FieldSize, sample.FieldMemo:
			values[i] = new(sql.NullString)
		case sample.FieldCreatedAt, sample.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case sample.ForeignKeys[0]: // office_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sample", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sample fields.
func (s *Sample) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sample.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case sample.FieldOfficeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field office_id", values[i])
			} else if value.Valid {
				s.OfficeID = int(value.Int64)
			}
		case sample.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case sample.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				s.Size = sample.Size(value.String)
			}
		case sample.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				s.Amount = value.Float64
			}
		case sample.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				s.Memo = new(string)
				*s.Memo = value.String
			}
		case sample.FieldURL:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.URL); err != nil {
					return fmt.Errorf("unmarshal field url: %w", err)
				}
			}
		case sample.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				s.Active = value.Bool
			}
		case sample.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sample.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case sample.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field office_id", value)
			} else if value.Valid {
				s.office_id = new(int)
				*s.office_id = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Sample.
// Note that you need to call Sample.Unwrap() before calling this method if this Sample
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sample) Update() *SampleUpdateOne {
	return (&SampleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sample entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sample) Unwrap() *Sample {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sample is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sample) String() string {
	var builder strings.Builder
	builder.WriteString("Sample(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", office_id=")
	builder.WriteString(fmt.Sprintf("%v", s.OfficeID))
	builder.WriteString(", code=")
	builder.WriteString(s.Code)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", s.Size))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", s.Amount))
	if v := s.Memo; v != nil {
		builder.WriteString(", memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", url=")
	builder.WriteString(fmt.Sprintf("%v", s.URL))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", s.Active))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Samples is a parsable slice of Sample.
type Samples []*Sample

func (s Samples) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
