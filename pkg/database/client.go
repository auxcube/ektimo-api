package database

import (
	"context"
	"database/sql"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/auxcube/ektimo-api/ent"
)

const (
	_defaultMaxIdleConns    int           = 10
	_defaultMaxOpenConns    int           = 100
	_defaultConnMaxLifetime time.Duration = time.Hour
)

type Client struct {
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration

	*ent.Client
	db *sql.DB
}

func Open(driverName, dataSourceName string, opts ...Option) (*Client, error) {
	client := &Client{
		maxIdleConns:    _defaultMaxIdleConns,
		maxOpenConns:    _defaultMaxOpenConns,
		connMaxLifetime: _defaultConnMaxLifetime,
	}

	for _, opt := range opts {
		opt(client)
	}

	drv, err := entsql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(client.maxIdleConns)
	db.SetMaxOpenConns(client.maxOpenConns)
	db.SetConnMaxLifetime(client.connMaxLifetime)
	entClient := ent.NewClient(ent.Driver(drv))

	client.db = db
	client.Client = entClient
	return client, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}
