package pqtest_test

import (
	"database/sql/driver"
	"errors"
	"io"
	"time"

	"github.com/cockroachdb/copyist"
)

var _ = driver.ErrBadConn
var _ = io.EOF
var _ = time.Parse
var _ = copyist.Register
var _ = errors.New

func init() {
	copyist.AddRecording(`postgres/github.com/cockroachdb/copyist/pqtest_test.TestQuery`, []copyist.Record{{Typ: copyist.DriverOpen, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`SELECT name FROM customers WHERE id=$1`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{1}}, {Typ: copyist.StmtQuery, Args: copyist.RecordArgs{nil}}, {Typ: copyist.RowsColumns, Args: copyist.RecordArgs{[]string{`name`}}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{`Andy`}, nil}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{}, io.EOF}}})
	copyist.AddRecording(`postgres/github.com/cockroachdb/copyist/pqtest_test.TestInsert`, []copyist.Record{{Typ: copyist.DriverOpen, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`INSERT INTO customers VALUES ($1, $2)`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{2}}, {Typ: copyist.StmtExec, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ResultRowsAffected, Args: copyist.RecordArgs{int64(1), nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`SELECT COUNT(*) FROM customers`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{0}}, {Typ: copyist.StmtQuery, Args: copyist.RecordArgs{nil}}, {Typ: copyist.RowsColumns, Args: copyist.RecordArgs{[]string{`count`}}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{int64(4)}, nil}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{}, io.EOF}}})
	copyist.AddRecording(`postgres/github.com/cockroachdb/copyist/pqtest_test.TestDataTypes`, []copyist.Record{{Typ: copyist.DriverOpen, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`
		CREATE TABLE datatypes
		(i INT, s TEXT, t TIMESTAMP, b BOOL, by BYTES, f FLOAT, d DECIMAL, fa FLOAT[])
	`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{0}}, {Typ: copyist.StmtExec, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`
		INSERT INTO datatypes
		VALUES (1, 'foo', '2000-01-01T10:00:00Z', true, 'ABCD', 1.1, 100.1234, ARRAY(1.1, 2.2))
	`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{0}}, {Typ: copyist.StmtExec, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ResultRowsAffected, Args: copyist.RecordArgs{int64(0), nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`SELECT i, s, t, b, by, f, d, fa FROM datatypes`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{0}}, {Typ: copyist.StmtQuery, Args: copyist.RecordArgs{nil}}, {Typ: copyist.RowsColumns, Args: copyist.RecordArgs{[]string{`i`, `s`, `t`, `b`, `by`, `f`, `d`, `fa`}}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{int64(1), `foo`, copyist.ParseTime(`2000-01-01T10:00:00Z`), true, []uint8{65, 66, 67, 68}, 1.100000000000000089, []uint8{49, 48, 48, 46, 49, 50, 51, 52}, []uint8{123, 49, 46, 49, 44, 50, 46, 50, 125}}, nil}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{}, io.EOF}}})
	copyist.AddRecording(`postgres/github.com/cockroachdb/copyist/pqtest_test.TestTxns`, []copyist.Record{{Typ: copyist.DriverOpen, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnBegin, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`INSERT INTO customers VALUES ($1, $2)`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{2}}, {Typ: copyist.StmtExec, Args: copyist.RecordArgs{nil}}, {Typ: copyist.TxCommit, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnBegin, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`INSERT INTO customers VALUES ($1, $2)`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{2}}, {Typ: copyist.StmtExec, Args: copyist.RecordArgs{nil}}, {Typ: copyist.TxRollback, Args: copyist.RecordArgs{nil}}, {Typ: copyist.ConnPrepare, Args: copyist.RecordArgs{`SELECT COUNT(*) FROM customers`, nil}}, {Typ: copyist.StmtNumInput, Args: copyist.RecordArgs{0}}, {Typ: copyist.StmtQuery, Args: copyist.RecordArgs{nil}}, {Typ: copyist.RowsColumns, Args: copyist.RecordArgs{[]string{`count`}}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{int64(4)}, nil}}, {Typ: copyist.RowsNext, Args: copyist.RecordArgs{[]driver.Value{}, io.EOF}}})
}