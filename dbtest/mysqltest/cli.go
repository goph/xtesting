package mysqltest

import "github.com/goph/env"

// PrepareConfig registers configuration in the EnvVarSet.
func PrepareConfig(e *env.EnvVarSet, config *Config) {
	e.StringVar(&config.Host, "db-host", "", "Database host")
	e.IntVar(&config.Port, "db-port", 3306, "Database port")
	e.StringVar(&config.User, "db-user", "root", "Database user")
	e.StringVar(&config.Pass, "db-pass", "", "Database password")
	e.StringVar(&config.Name, "db-name", "", "Database name")
	e.Var(&config.Params, "db-params", "Database params")
}
