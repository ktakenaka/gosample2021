// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/ktakenaka/gosample2021/ent/office"
	"github.com/ktakenaka/gosample2021/ent/predicate"
	"github.com/ktakenaka/gosample2021/ent/sample"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOffice = "Office"
	TypeSample = "Sample"
)

// OfficeMutation represents an operation that mutates the Office nodes in the graph.
type OfficeMutation struct {
	config
	op             Op
	typ            string
	id             *int
	code           *string
	name           *string
	clearedFields  map[string]struct{}
	samples        map[string]struct{}
	removedsamples map[string]struct{}
	clearedsamples bool
	done           bool
	oldValue       func(context.Context) (*Office, error)
	predicates     []predicate.Office
}

var _ ent.Mutation = (*OfficeMutation)(nil)

// officeOption allows management of the mutation configuration using functional options.
type officeOption func(*OfficeMutation)

// newOfficeMutation creates new mutation for the Office entity.
func newOfficeMutation(c config, op Op, opts ...officeOption) *OfficeMutation {
	m := &OfficeMutation{
		config:        c,
		op:            op,
		typ:           TypeOffice,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOfficeID sets the ID field of the mutation.
func withOfficeID(id int) officeOption {
	return func(m *OfficeMutation) {
		var (
			err   error
			once  sync.Once
			value *Office
		)
		m.oldValue = func(ctx context.Context) (*Office, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Office.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOffice sets the old Office of the mutation.
func withOffice(node *Office) officeOption {
	return func(m *OfficeMutation) {
		m.oldValue = func(context.Context) (*Office, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OfficeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OfficeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OfficeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCode sets the "code" field.
func (m *OfficeMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *OfficeMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Office entity.
// If the Office object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OfficeMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *OfficeMutation) ResetCode() {
	m.code = nil
}

// SetName sets the "name" field.
func (m *OfficeMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *OfficeMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Office entity.
// If the Office object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OfficeMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *OfficeMutation) ResetName() {
	m.name = nil
}

// AddSampleIDs adds the "samples" edge to the Sample entity by ids.
func (m *OfficeMutation) AddSampleIDs(ids ...string) {
	if m.samples == nil {
		m.samples = make(map[string]struct{})
	}
	for i := range ids {
		m.samples[ids[i]] = struct{}{}
	}
}

// ClearSamples clears the "samples" edge to the Sample entity.
func (m *OfficeMutation) ClearSamples() {
	m.clearedsamples = true
}

// SamplesCleared reports if the "samples" edge to the Sample entity was cleared.
func (m *OfficeMutation) SamplesCleared() bool {
	return m.clearedsamples
}

// RemoveSampleIDs removes the "samples" edge to the Sample entity by IDs.
func (m *OfficeMutation) RemoveSampleIDs(ids ...string) {
	if m.removedsamples == nil {
		m.removedsamples = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.samples, ids[i])
		m.removedsamples[ids[i]] = struct{}{}
	}
}

// RemovedSamples returns the removed IDs of the "samples" edge to the Sample entity.
func (m *OfficeMutation) RemovedSamplesIDs() (ids []string) {
	for id := range m.removedsamples {
		ids = append(ids, id)
	}
	return
}

// SamplesIDs returns the "samples" edge IDs in the mutation.
func (m *OfficeMutation) SamplesIDs() (ids []string) {
	for id := range m.samples {
		ids = append(ids, id)
	}
	return
}

// ResetSamples resets all changes to the "samples" edge.
func (m *OfficeMutation) ResetSamples() {
	m.samples = nil
	m.clearedsamples = false
	m.removedsamples = nil
}

// Where appends a list predicates to the OfficeMutation builder.
func (m *OfficeMutation) Where(ps ...predicate.Office) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *OfficeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Office).
func (m *OfficeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OfficeMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.code != nil {
		fields = append(fields, office.FieldCode)
	}
	if m.name != nil {
		fields = append(fields, office.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OfficeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case office.FieldCode:
		return m.Code()
	case office.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OfficeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case office.FieldCode:
		return m.OldCode(ctx)
	case office.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Office field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OfficeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case office.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case office.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Office field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OfficeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OfficeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OfficeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Office numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OfficeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OfficeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OfficeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Office nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OfficeMutation) ResetField(name string) error {
	switch name {
	case office.FieldCode:
		m.ResetCode()
		return nil
	case office.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Office field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OfficeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.samples != nil {
		edges = append(edges, office.EdgeSamples)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OfficeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case office.EdgeSamples:
		ids := make([]ent.Value, 0, len(m.samples))
		for id := range m.samples {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OfficeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedsamples != nil {
		edges = append(edges, office.EdgeSamples)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OfficeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case office.EdgeSamples:
		ids := make([]ent.Value, 0, len(m.removedsamples))
		for id := range m.removedsamples {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OfficeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedsamples {
		edges = append(edges, office.EdgeSamples)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OfficeMutation) EdgeCleared(name string) bool {
	switch name {
	case office.EdgeSamples:
		return m.clearedsamples
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OfficeMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Office unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OfficeMutation) ResetEdge(name string) error {
	switch name {
	case office.EdgeSamples:
		m.ResetSamples()
		return nil
	}
	return fmt.Errorf("unknown Office edge %s", name)
}

// SampleMutation represents an operation that mutates the Sample nodes in the graph.
type SampleMutation struct {
	config
	op            Op
	typ           string
	id            *string
	office_id     *int
	addoffice_id  *int
	code          *string
	size          *sample.Size
	amount        *float64
	addamount     *float64
	memo          *string
	url           **url.URL
	active        *bool
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Sample, error)
	predicates    []predicate.Sample
}

var _ ent.Mutation = (*SampleMutation)(nil)

// sampleOption allows management of the mutation configuration using functional options.
type sampleOption func(*SampleMutation)

// newSampleMutation creates new mutation for the Sample entity.
func newSampleMutation(c config, op Op, opts ...sampleOption) *SampleMutation {
	m := &SampleMutation{
		config:        c,
		op:            op,
		typ:           TypeSample,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSampleID sets the ID field of the mutation.
func withSampleID(id string) sampleOption {
	return func(m *SampleMutation) {
		var (
			err   error
			once  sync.Once
			value *Sample
		)
		m.oldValue = func(ctx context.Context) (*Sample, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Sample.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSample sets the old Sample of the mutation.
func withSample(node *Sample) sampleOption {
	return func(m *SampleMutation) {
		m.oldValue = func(context.Context) (*Sample, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SampleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SampleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Sample entities.
func (m *SampleMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SampleMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetOfficeID sets the "office_id" field.
func (m *SampleMutation) SetOfficeID(i int) {
	m.office_id = &i
	m.addoffice_id = nil
}

// OfficeID returns the value of the "office_id" field in the mutation.
func (m *SampleMutation) OfficeID() (r int, exists bool) {
	v := m.office_id
	if v == nil {
		return
	}
	return *v, true
}

// OldOfficeID returns the old "office_id" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldOfficeID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOfficeID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOfficeID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOfficeID: %w", err)
	}
	return oldValue.OfficeID, nil
}

// AddOfficeID adds i to the "office_id" field.
func (m *SampleMutation) AddOfficeID(i int) {
	if m.addoffice_id != nil {
		*m.addoffice_id += i
	} else {
		m.addoffice_id = &i
	}
}

// AddedOfficeID returns the value that was added to the "office_id" field in this mutation.
func (m *SampleMutation) AddedOfficeID() (r int, exists bool) {
	v := m.addoffice_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetOfficeID resets all changes to the "office_id" field.
func (m *SampleMutation) ResetOfficeID() {
	m.office_id = nil
	m.addoffice_id = nil
}

// SetCode sets the "code" field.
func (m *SampleMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *SampleMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *SampleMutation) ResetCode() {
	m.code = nil
}

// SetSize sets the "size" field.
func (m *SampleMutation) SetSize(s sample.Size) {
	m.size = &s
}

// Size returns the value of the "size" field in the mutation.
func (m *SampleMutation) Size() (r sample.Size, exists bool) {
	v := m.size
	if v == nil {
		return
	}
	return *v, true
}

// OldSize returns the old "size" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldSize(ctx context.Context) (v sample.Size, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSize is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSize requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSize: %w", err)
	}
	return oldValue.Size, nil
}

// ResetSize resets all changes to the "size" field.
func (m *SampleMutation) ResetSize() {
	m.size = nil
}

// SetAmount sets the "amount" field.
func (m *SampleMutation) SetAmount(f float64) {
	m.amount = &f
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *SampleMutation) Amount() (r float64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds f to the "amount" field.
func (m *SampleMutation) AddAmount(f float64) {
	if m.addamount != nil {
		*m.addamount += f
	} else {
		m.addamount = &f
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *SampleMutation) AddedAmount() (r float64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *SampleMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetMemo sets the "memo" field.
func (m *SampleMutation) SetMemo(s string) {
	m.memo = &s
}

// Memo returns the value of the "memo" field in the mutation.
func (m *SampleMutation) Memo() (r string, exists bool) {
	v := m.memo
	if v == nil {
		return
	}
	return *v, true
}

// OldMemo returns the old "memo" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldMemo(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMemo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMemo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMemo: %w", err)
	}
	return oldValue.Memo, nil
}

// ClearMemo clears the value of the "memo" field.
func (m *SampleMutation) ClearMemo() {
	m.memo = nil
	m.clearedFields[sample.FieldMemo] = struct{}{}
}

// MemoCleared returns if the "memo" field was cleared in this mutation.
func (m *SampleMutation) MemoCleared() bool {
	_, ok := m.clearedFields[sample.FieldMemo]
	return ok
}

// ResetMemo resets all changes to the "memo" field.
func (m *SampleMutation) ResetMemo() {
	m.memo = nil
	delete(m.clearedFields, sample.FieldMemo)
}

// SetURL sets the "url" field.
func (m *SampleMutation) SetURL(u *url.URL) {
	m.url = &u
}

// URL returns the value of the "url" field in the mutation.
func (m *SampleMutation) URL() (r *url.URL, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldURL(ctx context.Context) (v *url.URL, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ClearURL clears the value of the "url" field.
func (m *SampleMutation) ClearURL() {
	m.url = nil
	m.clearedFields[sample.FieldURL] = struct{}{}
}

// URLCleared returns if the "url" field was cleared in this mutation.
func (m *SampleMutation) URLCleared() bool {
	_, ok := m.clearedFields[sample.FieldURL]
	return ok
}

// ResetURL resets all changes to the "url" field.
func (m *SampleMutation) ResetURL() {
	m.url = nil
	delete(m.clearedFields, sample.FieldURL)
}

// SetActive sets the "active" field.
func (m *SampleMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the value of the "active" field in the mutation.
func (m *SampleMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// OldActive returns the old "active" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActive: %w", err)
	}
	return oldValue.Active, nil
}

// ResetActive resets all changes to the "active" field.
func (m *SampleMutation) ResetActive() {
	m.active = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *SampleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SampleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SampleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *SampleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *SampleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Sample entity.
// If the Sample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SampleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *SampleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the SampleMutation builder.
func (m *SampleMutation) Where(ps ...predicate.Sample) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SampleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Sample).
func (m *SampleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SampleMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.office_id != nil {
		fields = append(fields, sample.FieldOfficeID)
	}
	if m.code != nil {
		fields = append(fields, sample.FieldCode)
	}
	if m.size != nil {
		fields = append(fields, sample.FieldSize)
	}
	if m.amount != nil {
		fields = append(fields, sample.FieldAmount)
	}
	if m.memo != nil {
		fields = append(fields, sample.FieldMemo)
	}
	if m.url != nil {
		fields = append(fields, sample.FieldURL)
	}
	if m.active != nil {
		fields = append(fields, sample.FieldActive)
	}
	if m.created_at != nil {
		fields = append(fields, sample.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, sample.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SampleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sample.FieldOfficeID:
		return m.OfficeID()
	case sample.FieldCode:
		return m.Code()
	case sample.FieldSize:
		return m.Size()
	case sample.FieldAmount:
		return m.Amount()
	case sample.FieldMemo:
		return m.Memo()
	case sample.FieldURL:
		return m.URL()
	case sample.FieldActive:
		return m.Active()
	case sample.FieldCreatedAt:
		return m.CreatedAt()
	case sample.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SampleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sample.FieldOfficeID:
		return m.OldOfficeID(ctx)
	case sample.FieldCode:
		return m.OldCode(ctx)
	case sample.FieldSize:
		return m.OldSize(ctx)
	case sample.FieldAmount:
		return m.OldAmount(ctx)
	case sample.FieldMemo:
		return m.OldMemo(ctx)
	case sample.FieldURL:
		return m.OldURL(ctx)
	case sample.FieldActive:
		return m.OldActive(ctx)
	case sample.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case sample.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Sample field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SampleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sample.FieldOfficeID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOfficeID(v)
		return nil
	case sample.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case sample.FieldSize:
		v, ok := value.(sample.Size)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSize(v)
		return nil
	case sample.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case sample.FieldMemo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemo(v)
		return nil
	case sample.FieldURL:
		v, ok := value.(*url.URL)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case sample.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case sample.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case sample.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Sample field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SampleMutation) AddedFields() []string {
	var fields []string
	if m.addoffice_id != nil {
		fields = append(fields, sample.FieldOfficeID)
	}
	if m.addamount != nil {
		fields = append(fields, sample.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SampleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case sample.FieldOfficeID:
		return m.AddedOfficeID()
	case sample.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SampleMutation) AddField(name string, value ent.Value) error {
	switch name {
	case sample.FieldOfficeID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOfficeID(v)
		return nil
	case sample.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown Sample numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SampleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(sample.FieldMemo) {
		fields = append(fields, sample.FieldMemo)
	}
	if m.FieldCleared(sample.FieldURL) {
		fields = append(fields, sample.FieldURL)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SampleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SampleMutation) ClearField(name string) error {
	switch name {
	case sample.FieldMemo:
		m.ClearMemo()
		return nil
	case sample.FieldURL:
		m.ClearURL()
		return nil
	}
	return fmt.Errorf("unknown Sample nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SampleMutation) ResetField(name string) error {
	switch name {
	case sample.FieldOfficeID:
		m.ResetOfficeID()
		return nil
	case sample.FieldCode:
		m.ResetCode()
		return nil
	case sample.FieldSize:
		m.ResetSize()
		return nil
	case sample.FieldAmount:
		m.ResetAmount()
		return nil
	case sample.FieldMemo:
		m.ResetMemo()
		return nil
	case sample.FieldURL:
		m.ResetURL()
		return nil
	case sample.FieldActive:
		m.ResetActive()
		return nil
	case sample.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case sample.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Sample field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SampleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SampleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SampleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SampleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SampleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SampleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SampleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Sample unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SampleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Sample edge %s", name)
}
