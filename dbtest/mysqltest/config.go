package mysqltest

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Config holds information necessary for connecting to a database.
type Config struct {
	Host string
	Port int
	User string
	Pass string
	Name string

	Params DsnParams
}

// Validate checks that the configuration is valid.
func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("database host is required")
	}

	if c.Port == 0 {
		return errors.New("database port is required")
	}

	if c.User == "" {
		return errors.New("database user is required")
	}

	if c.Name == "" {
		return errors.New("database name is required")
	}

	return nil
}

// DSN returns a MySQL driver compatible data source name.
func (c Config) DSN() string {
	var params string

	if len(c.Params) > 0 {
		params = "?" + c.Params.String()
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
		params,
	)
}

// DsnParams implements the common Value interface implemented by flag/pflag/env packages.
type DsnParams map[string]string

func (p DsnParams) Set(val string) error {
	for _, v := range strings.Split(val, "&") {
		param := strings.SplitN(v, "=", 2)
		if len(param) != 2 {
			continue
		}

		p[param[0]] = param[1]
	}

	return nil
}

func (DsnParams) Type() string {
	return "query string"
}

func (p DsnParams) String() string {
	var query string

	for key, value := range p {
		if query != "" {
			query += "&"
		}

		query += key + "=" + value
	}

	return query
}
