// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ktakenaka/gosample/ent/migrate"

	"github.com/ktakenaka/gosample/ent/office"
	"github.com/ktakenaka/gosample/ent/sample"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Office is the client for interacting with the Office builders.
	Office *OfficeClient
	// Sample is the client for interacting with the Sample builders.
	Sample *SampleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Office = NewOfficeClient(c.config)
	c.Sample = NewSampleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Office: NewOfficeClient(cfg),
		Sample: NewSampleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Office: NewOfficeClient(cfg),
		Sample: NewSampleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Office.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Office.Use(hooks...)
	c.Sample.Use(hooks...)
}

// OfficeClient is a client for the Office schema.
type OfficeClient struct {
	config
}

// NewOfficeClient returns a client for the Office from the given config.
func NewOfficeClient(c config) *OfficeClient {
	return &OfficeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `office.Hooks(f(g(h())))`.
func (c *OfficeClient) Use(hooks ...Hook) {
	c.hooks.Office = append(c.hooks.Office, hooks...)
}

// Create returns a create builder for Office.
func (c *OfficeClient) Create() *OfficeCreate {
	mutation := newOfficeMutation(c.config, OpCreate)
	return &OfficeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Office entities.
func (c *OfficeClient) CreateBulk(builders ...*OfficeCreate) *OfficeCreateBulk {
	return &OfficeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Office.
func (c *OfficeClient) Update() *OfficeUpdate {
	mutation := newOfficeMutation(c.config, OpUpdate)
	return &OfficeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OfficeClient) UpdateOne(o *Office) *OfficeUpdateOne {
	mutation := newOfficeMutation(c.config, OpUpdateOne, withOffice(o))
	return &OfficeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OfficeClient) UpdateOneID(id int) *OfficeUpdateOne {
	mutation := newOfficeMutation(c.config, OpUpdateOne, withOfficeID(id))
	return &OfficeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Office.
func (c *OfficeClient) Delete() *OfficeDelete {
	mutation := newOfficeMutation(c.config, OpDelete)
	return &OfficeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OfficeClient) DeleteOne(o *Office) *OfficeDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OfficeClient) DeleteOneID(id int) *OfficeDeleteOne {
	builder := c.Delete().Where(office.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OfficeDeleteOne{builder}
}

// Query returns a query builder for Office.
func (c *OfficeClient) Query() *OfficeQuery {
	return &OfficeQuery{
		config: c.config,
	}
}

// Get returns a Office entity by its id.
func (c *OfficeClient) Get(ctx context.Context, id int) (*Office, error) {
	return c.Query().Where(office.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OfficeClient) GetX(ctx context.Context, id int) *Office {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySamples queries the samples edge of a Office.
func (c *OfficeClient) QuerySamples(o *Office) *SampleQuery {
	query := &SampleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(office.Table, office.FieldID, id),
			sqlgraph.To(sample.Table, sample.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, office.SamplesTable, office.SamplesColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OfficeClient) Hooks() []Hook {
	return c.hooks.Office
}

// SampleClient is a client for the Sample schema.
type SampleClient struct {
	config
}

// NewSampleClient returns a client for the Sample from the given config.
func NewSampleClient(c config) *SampleClient {
	return &SampleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sample.Hooks(f(g(h())))`.
func (c *SampleClient) Use(hooks ...Hook) {
	c.hooks.Sample = append(c.hooks.Sample, hooks...)
}

// Create returns a create builder for Sample.
func (c *SampleClient) Create() *SampleCreate {
	mutation := newSampleMutation(c.config, OpCreate)
	return &SampleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Sample entities.
func (c *SampleClient) CreateBulk(builders ...*SampleCreate) *SampleCreateBulk {
	return &SampleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Sample.
func (c *SampleClient) Update() *SampleUpdate {
	mutation := newSampleMutation(c.config, OpUpdate)
	return &SampleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SampleClient) UpdateOne(s *Sample) *SampleUpdateOne {
	mutation := newSampleMutation(c.config, OpUpdateOne, withSample(s))
	return &SampleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SampleClient) UpdateOneID(id string) *SampleUpdateOne {
	mutation := newSampleMutation(c.config, OpUpdateOne, withSampleID(id))
	return &SampleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Sample.
func (c *SampleClient) Delete() *SampleDelete {
	mutation := newSampleMutation(c.config, OpDelete)
	return &SampleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SampleClient) DeleteOne(s *Sample) *SampleDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SampleClient) DeleteOneID(id string) *SampleDeleteOne {
	builder := c.Delete().Where(sample.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SampleDeleteOne{builder}
}

// Query returns a query builder for Sample.
func (c *SampleClient) Query() *SampleQuery {
	return &SampleQuery{
		config: c.config,
	}
}

// Get returns a Sample entity by its id.
func (c *SampleClient) Get(ctx context.Context, id string) (*Sample, error) {
	return c.Query().Where(sample.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SampleClient) GetX(ctx context.Context, id string) *Sample {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SampleClient) Hooks() []Hook {
	return c.hooks.Sample
}
