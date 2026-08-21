package sqlite3

import (
	"os"
	"testing"
	"todo/pkg/config"

	"github.com/stretchr/testify/require"
)

var mockTaks = []config.Task{
	{Title: "example task 1", Desc: "some task 1", Duration: "1h"},
	{Title: "example task 2", Desc: "some task 2", Duration: "2h"},
}

func TestSQLITE(t *testing.T) {
	database := "./test.db"
	os.Remove(database)
	conn := Connect(Dsn(database))
	defer conn.Close()
	require.NoError(t, conn.InitTables())
	require.NoError(t, conn.InsertTaks(mockTaks))
	task, err := conn.CurrentTask()
	require.NoError(t, err)
	t.Log(task)
}
