package sqlite3

var tableStmts = []string{
	SQL_CREATE_TASKS_TABLE,
}

var (
	SQL_CREATE_TASKS_TABLE = `
		CREATE TABLE IF NOT EXISTS tasks (
			title VARCHAR(500) NOT NULL,
			desc VARCHAR(700),
			start_time BIGINT,
			end_time BIGINT
		);
	`

	SQL_INSERT_TASK = `INSERT INTO tasks (title, desc, start_time, end_time) VALUES (?, ?, ?, ?);`

	// selects a task using a unix timestamp
	SQL_SELECT_TASK_BY_TIME = `SELECT title, desc, start_time, end_time FROM tasks WHERE start_time <= ? AND end_time >= ? LIMIT 1`
)
