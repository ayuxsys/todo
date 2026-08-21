package todo

// todo: add desc
type Task struct {
	Title string
	Desc  string
	Time  Time
}

type Time struct {
	Start int64
	End   int64
}
