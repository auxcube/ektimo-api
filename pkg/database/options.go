package database

import "time"

type Option func(*Client)

func MaxIdleConns(n int) Option {
	return func(c *Client) {
		c.maxIdleConns = n
	}
}

func MaxOpenConns(n int) Option {
	return func(c *Client) {
		c.maxOpenConns = n
	}
}

func ConnMaxLifetime(t time.Duration) Option {
	return func(c *Client) {
		c.connMaxLifetime = t
	}
}
