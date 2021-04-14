package config

import "fmt"

// GetAddr returns string with server host and port
func (c serverConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// GetDSN returns string with connection string for database
func (c *databaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?multiStatements=true&parseTime=true&loc=UTC&collation=utf8mb4_general_ci",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}
