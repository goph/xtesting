package mysqltest

// value is the common interface used by flag/pflag/env packages.
type value interface {
	Set(string) error
	Type() string
	String() string
}

// configurator implements the config methods required by the database config.
type configurator interface {
	StringVar(p *string, name string, value string, usage string)
	IntVar(p *int, name string, value int, usage string)
	Var(value value, name string, usage string)
}

// PrepareConfig registers configuration in the configurator.
func PrepareConfig(c configurator, config *Config) {
	c.StringVar(&config.Host, "db-host", "", "Database host")
	c.IntVar(&config.Port, "db-port", 3306, "Database port")
	c.StringVar(&config.User, "db-user", "root", "Database user")
	c.StringVar(&config.Pass, "db-pass", "", "Database password")
	c.StringVar(&config.Name, "db-name", "", "Database name")
	c.Var(&config.Params, "db-params", "Database params")
}
