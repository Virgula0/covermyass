//go:build !windows

package check

type databaseServerCheck struct{}

func NewDatabaseServerCheck() Check {
	return &databaseServerCheck{}
}

func (s *databaseServerCheck) Name() string {
	return "database-server"
}

func (s *databaseServerCheck) Paths() []string {
	return []string{
		"/var/log/mysqld.log",
		"/var/log/mysql.log",
		"/var/log/mysql/error.log",
		"/var/log/mysql/mysql.log",
		"/var/log/mysql/mysql-slow.log",
	}
}

func init() {
	AddCheck(NewDatabaseServerCheck())
}
