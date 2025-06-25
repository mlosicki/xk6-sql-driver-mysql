// Package mysql contains MySQL driver registration for xk6-sql.
package mysql

import (
	"github.com/mlosicki/xk6-sql-ext/sql"
	"go.k6.io/k6/js/modules"

	// Blank import required for initialization of driver.
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	id := sql.RegisterDriver("mysql")

	modules.Register("k6/x/sql/driver/mysql", &rootModule{driverID: id})
}
