// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ktakenaka/gosample/ent/office"
	"github.com/ktakenaka/gosample/ent/sample"
	"github.com/ktakenaka/gosample/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	officeFields := schema.Office{}.Fields()
	_ = officeFields
	// officeDescName is the schema descriptor for name field.
	officeDescName := officeFields[0].Descriptor()
	// office.NameValidator is a validator for the "name" field. It is called by the builders before save.
	office.NameValidator = func() func(string) error {
		validators := officeDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sampleFields := schema.Sample{}.Fields()
	_ = sampleFields
	// sampleDescCode is the schema descriptor for code field.
	sampleDescCode := sampleFields[2].Descriptor()
	// sample.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	sample.CodeValidator = func() func(string) error {
		validators := sampleDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sampleDescActive is the schema descriptor for active field.
	sampleDescActive := sampleFields[7].Descriptor()
	// sample.DefaultActive holds the default value on creation for the active field.
	sample.DefaultActive = sampleDescActive.Default.(bool)
	// sampleDescCreatedAt is the schema descriptor for created_at field.
	sampleDescCreatedAt := sampleFields[8].Descriptor()
	// sample.DefaultCreatedAt holds the default value on creation for the created_at field.
	sample.DefaultCreatedAt = sampleDescCreatedAt.Default.(func() time.Time)
	// sampleDescUpdatedAt is the schema descriptor for updated_at field.
	sampleDescUpdatedAt := sampleFields[9].Descriptor()
	// sample.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sample.DefaultUpdatedAt = sampleDescUpdatedAt.Default.(func() time.Time)
	// sample.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sample.UpdateDefaultUpdatedAt = sampleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sampleDescID is the schema descriptor for id field.
	sampleDescID := sampleFields[0].Descriptor()
	// sample.DefaultID holds the default value on creation for the id field.
	sample.DefaultID = sampleDescID.Default.(func() string)
}
