/*
Package cache provides a Hord database driver for a variety of caching strategies. To use this driver, import it as follows:

	import (
	    "github.com/tarmac-project/hord"
	    "github.com/tarmac-project/hord/cache"
	)

# Connecting to the Database

Use the Dial() function to create a new client for interacting with the cache.

	// Handle database connection
	var database hord.Database
	...

	// Handle cache connection
	var cache hord.Database
	...

	var db hord.Database
	db, err := cache.Dial(cache.Config{
		Database: database,
		Cache: 	  cache,
		Type:	 cache.Lookaside,
	})
	if err != nil {
	    // Handle connection error
	}

# Initialize database

Hord provides a Setup() function for preparing a database. This function is safe to execute after every Dial().

	err := db.Setup()
	if err != nil {
	    // Handle setup error
	}

# Database Operations

Hord provides a simple abstraction for working with the cache, with easy-to-use methods such as Get() and Set() to read and write values.

	// Handle database connection
	var database hord.Database
	database, err := cassandra.Dial(cassandra.Config{})
	if err != nil {
		// Handle connection error
	}

	// Handle cache connection
	var cache hord.Database
	cache, err := redis.Dial(redis.Config{})
	if err != nil {
		// Handle connection error
	}

	// Connect to the Cache database
	db, err := cache.Dial(cache.Config{
		Database: database,
		Cache: 	  cache,
		Type:	 cache.Lookaside,
	})
	if err != nil {
	    // Handle connection error
	}

	err := db.Setup()
	if err != nil {
	    // Handle setup error
	}

	// Set a value
	err = db.Set("key", []byte("value"))
	if err != nil {
	    // Handle error
	}

	// Retrieve a value
	value, err := db.Get("key")
	if err != nil {
	    // Handle error
	}
*/
package cache

import (
	"errors"

	"github.com/tarmac-project/hord"
	"github.com/tarmac-project/hord/cache/lookaside"
)

// CacheType is the type of cache to use.
type Type string

const (
	Lookaside Type = "lookaside"
	None      Type = "none"
)

// Config provides the configuration options for the Cache driver.
type Config struct {
	Type     Type
	Database hord.Database
	Cache    hord.Database
}

// NilCache is a nil cache driver that returns dial errors. It fixes the issue when the Dial function returns a nil hord.Database this prevents nil pointer errors.
type NilCache struct{}

var (
	// ErrNoType is returned when the CacheType is invalid.
	ErrNoType = errors.New("invalid CacheType")
)

// Dial will create a new Cache driver using the provided Config. It will return an error if either the Database or Cache values in Config are nil or if a CacheType is not specified.
func Dial(cfg Config) (hord.Database, error) {
	if (cfg.Database == nil) || (cfg.Cache == nil) {
		return &NilCache{}, hord.ErrInvalidDatabase
	}

	switch cfg.Type {
	case Lookaside:
		return lookaside.Dial(lookaside.Config{
			Database: cfg.Database,
			Cache:    cfg.Cache,
		})
	case None:
		return cfg.Database, nil
	default:
		return &NilCache{}, ErrNoType
	}
}

func (nc *NilCache) Setup() error {
	return hord.ErrNoDial
}

func (nc *NilCache) HealthCheck() error {
	return hord.ErrNoDial
}

func (nc *NilCache) Get(_ string) ([]byte, error) {
	return nil, hord.ErrNoDial
}

func (nc *NilCache) Set(_ string, _ []byte) error {
	return hord.ErrNoDial
}

func (nc *NilCache) Delete(_ string) error {
	return hord.ErrNoDial
}

func (nc *NilCache) Keys() ([]string, error) {
	return nil, hord.ErrNoDial
}

func (nc *NilCache) Close() {

}
