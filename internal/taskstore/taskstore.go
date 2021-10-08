package Task

import "time"

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

// No enums in go 😢
type Tag int

const (
	Home Tag = iota + 1
	Work
	School
)

// Task methods; API
func New() *Task

// creates a new task store
func (ts *Task) CreateTask(test string, tags []string, due time.Time) int

// Get the task from the datastore with the given id and return a reference/address of the task T to be derefenced by the caller
func (ts *Task) GetTask(id int) (Task, error)

func (ts *Task) GetTasks(id []int) ([]Task, error)

func (ts *Task) GetAllTasks() ([]Task, error)

// Delete the task from the datastore with the given id; else return err
func (ts *Task) DeleteTask(id int) error

// Delete set of tasks from the datastore from a g given list of id's; else return err
func (ts *Task) DeleteTasks(ids []int) error

// Delete all task from the data store
func (ts *Task) DeleteAllTasks() error

func (ts *Task) getTaskByTag(tag Tag) []Task

// Returns all the tasks that have the given due date, in arbitary order
func (ts *Task) getTaskByByDueDate(year int, month time.Month, data int) []Task
