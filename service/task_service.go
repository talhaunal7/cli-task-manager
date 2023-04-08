package service

type TaskService interface {
	Add([]string) error
	List() error
	Do(string) error
}
