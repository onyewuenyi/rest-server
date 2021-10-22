package Task
import (
    "database/sql"
	import "time"
)


type task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// No enums in go 😢. This is a workaround to express 
type Tag int

const (
	Home Tag = iota + 1
	Work
	School
)

// Task methods; API
func New() *TaskStore

// creates a new task store
func (t *task) CreateTask(db *sql.DB) error {

}

// // Get the task from the datastore with the given id and return a reference/address of the task T to be derefenced by the caller
// func (t *task) GetTask(id int) (Task, error)

// func (t *task) GetTasks(id []int) ([]Task, error)

// func (t *task) GetAllTasks() ([]Task, error)

// // Delete the task from the datastore with the given id; else return err
// func (t *task) DeleteTask(id int) error

// // Delete set of tasks from the datastore from a g given list of id's; else return err
// func (t *task) DeleteTasks(ids []int) error

// // Delete all task from the data store
// func (t *task) DeleteAllTasks() error

// func (t *task) getTaskByTag(tag Tag) []Task

// // Returns all the tasks that have the given due date, in arbitary order
// func (t *task) getTaskByByDueDate(year int, month time.Month, data int) []Task
