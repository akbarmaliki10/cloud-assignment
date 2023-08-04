package usecase

import "amitshekar-clean-arch/domain"

type TodoUsecaseImpl struct {
	repo domain.TodoRepo
}

// make a function that act like a constructor
func NewTodoUsecase(todoRepo domain.TodoRepo) domain.TodoUsecase {
	return &TodoUsecaseImpl{
		repo: todoRepo,
	}
}

// receiver function or more like classes/struct method in python/java
func (u *TodoUsecaseImpl) CreateTodo(todo *domain.Todo) error {
	err := u.repo.CreateTodo(todo)
	return err
}

func (u *TodoUsecaseImpl) GetTodo(name *string) (*domain.Todo, error) {
	todo, err := u.repo.GetTodo(name)
	return todo, err
}

func (u *TodoUsecaseImpl) GetAll() ([]*domain.Todo, error) {
	todos, err := u.repo.GetAll()
	return todos, err
}

func (u *TodoUsecaseImpl) UpdateTodo(todo *domain.Todo) error {
	err := u.repo.UpdateTodo(todo)
	return err
}

func (u *TodoUsecaseImpl) DeleteTodo(name *string) error {
	err := u.repo.DeleteTodo(name)
	return err
}
