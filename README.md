# pgqry

A simple library for querying postgresql

## API

```go
import email "github.com/richardanaya/pgqry"
import "database/sql"

func CreateConnectionString(host string, port int, user, password, dbname string) string
func Execute(connectionString, query string) (sql.Result, error)
func Query(connectionString, query string) (*sql.Rows, error)
```