package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
)

func main() {
	// database/sql provides the common database API.
	// Real database support comes from drivers.
	db := sql.OpenDB(demoConnector{})
	defer db.Close()

	ctx := context.Background()

	var name string
	if err := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", 1).Scan(&name); err != nil {
		fmt.Println("QueryRowContext error:", err)
		return
	}

	fmt.Println("query result:", name)
}

type demoConnector struct{}

func (demoConnector) Connect(context.Context) (driver.Conn, error) {
	return demoConn{}, nil
}

func (demoConnector) Driver() driver.Driver {
	return demoDriver{}
}

type demoDriver struct{}

func (demoDriver) Open(string) (driver.Conn, error) {
	return demoConn{}, nil
}

type demoConn struct{}

func (demoConn) Prepare(string) (driver.Stmt, error) {
	return nil, driver.ErrSkip
}

func (demoConn) Close() error {
	return nil
}

func (demoConn) Begin() (driver.Tx, error) {
	return nil, driver.ErrSkip
}

func (demoConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return &demoRows{
		columns: []string{"name"},
		values: [][]driver.Value{
			{"Gopher"},
		},
	}, nil
}

type demoRows struct {
	columns []string
	values  [][]driver.Value
	index   int
}

func (r demoRows) Columns() []string {
	return r.columns
}

func (r demoRows) Close() error {
	return nil
}

func (r *demoRows) Next(dest []driver.Value) error {
	if r.index >= len(r.values) {
		return io.EOF
	}

	row := r.values[r.index]
	r.index++
	copy(dest, row)
	return nil
}
