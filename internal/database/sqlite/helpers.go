package sqlite3

import (
	"fmt"
	"time"
	"todo/internal/log"
	"todo/internal/todo"
	"todo/internal/utils/must"
	"todo/pkg/config"
)

func (c *Connection) InsertTaks(tasks []config.Task) error {
	now := time.Now().UnixNano()
	var start, end int64
	for i, task := range tasks {
		log.Debug.Println("inserting task", task)
		title := task.Title
		duration := must.Eval(time.ParseDuration(task.Duration)).Nanoseconds()

		if i == 0 {
			start = now
		}
		end = start + duration
		if err := c.insertTask(title, task.Desc, start, end); err != nil {
			return fmt.Errorf("'insertTask' '%+v': %v", task, err)
		}
		start = end
	}
	return nil
}

func (c *Connection) insertTask(title, desc string, startTime, endTime int64) error {
	_, err := c.Exec(SQL_INSERT_TASK, title, desc, startTime, endTime)
	return err
}

func (c *Connection) CurrentTask() (task todo.Task, err error) {
	now := time.Now().UnixNano()
	err = c.QueryRow(SQL_SELECT_TASK_BY_TIME, now, now).Scan(&task.Title, &task.Desc, &task.Time.Start, &task.Time.End)
	return
}

func (c *Connection) ExecStmts(stmts []string) error {
	for _, stmt := range stmts {
		if _, err := c.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}
