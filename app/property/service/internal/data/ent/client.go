// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/migrate"

	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/community"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/shop"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Community is the client for interacting with the Community builders.
	Community *CommunityClient
	// Shop is the client for interacting with the Shop builders.
	Shop *ShopClient
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
	c.Community = NewCommunityClient(c.config)
	c.Shop = NewShopClient(c.config)
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
		ctx:       ctx,
		config:    cfg,
		Community: NewCommunityClient(cfg),
		Shop:      NewShopClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Community: NewCommunityClient(cfg),
		Shop:      NewShopClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Community.
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
	c.Community.Use(hooks...)
	c.Shop.Use(hooks...)
}

// CommunityClient is a client for the Community schema.
type CommunityClient struct {
	config
}

// NewCommunityClient returns a client for the Community from the given config.
func NewCommunityClient(c config) *CommunityClient {
	return &CommunityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `community.Hooks(f(g(h())))`.
func (c *CommunityClient) Use(hooks ...Hook) {
	c.hooks.Community = append(c.hooks.Community, hooks...)
}

// Create returns a create builder for Community.
func (c *CommunityClient) Create() *CommunityCreate {
	mutation := newCommunityMutation(c.config, OpCreate)
	return &CommunityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Community entities.
func (c *CommunityClient) CreateBulk(builders ...*CommunityCreate) *CommunityCreateBulk {
	return &CommunityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Community.
func (c *CommunityClient) Update() *CommunityUpdate {
	mutation := newCommunityMutation(c.config, OpUpdate)
	return &CommunityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommunityClient) UpdateOne(co *Community) *CommunityUpdateOne {
	mutation := newCommunityMutation(c.config, OpUpdateOne, withCommunity(co))
	return &CommunityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommunityClient) UpdateOneID(id int) *CommunityUpdateOne {
	mutation := newCommunityMutation(c.config, OpUpdateOne, withCommunityID(id))
	return &CommunityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Community.
func (c *CommunityClient) Delete() *CommunityDelete {
	mutation := newCommunityMutation(c.config, OpDelete)
	return &CommunityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommunityClient) DeleteOne(co *Community) *CommunityDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommunityClient) DeleteOneID(id int) *CommunityDeleteOne {
	builder := c.Delete().Where(community.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommunityDeleteOne{builder}
}

// Query returns a query builder for Community.
func (c *CommunityClient) Query() *CommunityQuery {
	return &CommunityQuery{
		config: c.config,
	}
}

// Get returns a Community entity by its id.
func (c *CommunityClient) Get(ctx context.Context, id int) (*Community, error) {
	return c.Query().Where(community.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommunityClient) GetX(ctx context.Context, id int) *Community {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommunityClient) Hooks() []Hook {
	return c.hooks.Community
}

// ShopClient is a client for the Shop schema.
type ShopClient struct {
	config
}

// NewShopClient returns a client for the Shop from the given config.
func NewShopClient(c config) *ShopClient {
	return &ShopClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `shop.Hooks(f(g(h())))`.
func (c *ShopClient) Use(hooks ...Hook) {
	c.hooks.Shop = append(c.hooks.Shop, hooks...)
}

// Create returns a create builder for Shop.
func (c *ShopClient) Create() *ShopCreate {
	mutation := newShopMutation(c.config, OpCreate)
	return &ShopCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Shop entities.
func (c *ShopClient) CreateBulk(builders ...*ShopCreate) *ShopCreateBulk {
	return &ShopCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Shop.
func (c *ShopClient) Update() *ShopUpdate {
	mutation := newShopMutation(c.config, OpUpdate)
	return &ShopUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShopClient) UpdateOne(s *Shop) *ShopUpdateOne {
	mutation := newShopMutation(c.config, OpUpdateOne, withShop(s))
	return &ShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShopClient) UpdateOneID(id int) *ShopUpdateOne {
	mutation := newShopMutation(c.config, OpUpdateOne, withShopID(id))
	return &ShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Shop.
func (c *ShopClient) Delete() *ShopDelete {
	mutation := newShopMutation(c.config, OpDelete)
	return &ShopDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ShopClient) DeleteOne(s *Shop) *ShopDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ShopClient) DeleteOneID(id int) *ShopDeleteOne {
	builder := c.Delete().Where(shop.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShopDeleteOne{builder}
}

// Query returns a query builder for Shop.
func (c *ShopClient) Query() *ShopQuery {
	return &ShopQuery{
		config: c.config,
	}
}

// Get returns a Shop entity by its id.
func (c *ShopClient) Get(ctx context.Context, id int) (*Shop, error) {
	return c.Query().Where(shop.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShopClient) GetX(ctx context.Context, id int) *Shop {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ShopClient) Hooks() []Hook {
	return c.hooks.Shop
}
