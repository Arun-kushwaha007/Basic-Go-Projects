package todo  // This line indicates that this file is part of the "todo" package, which is used for managing a simple command-line todo list application.

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
