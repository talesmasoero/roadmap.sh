package model

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	return statusTask[s]
}

var statusTask = map[Status]string{
	StatusTodo:       "todo",
	StatusInProgress: "in-progress",
	StatusDone:       "done",
}
