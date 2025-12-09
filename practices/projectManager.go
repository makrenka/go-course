package practices

import (
	"errors"
	"fmt"
)

type Project struct {
	ID    string
	Name  string
	Tasks []Task
}

type Task struct {
	ID          string
	Title       string
	Description string
	Status      bool
}

func NewProject(id string, name string) (*Project, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}

	return &Project{
		ID:   id,
		Name: name,
	}, nil
}

func NewTask(id string, title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("empty title")
	}
	if description == "" {
		return nil, errors.New("empty description")
	}

	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      true,
	}, nil
}

func (p *Project) AddTask(t Task) error {
	for _, task := range p.Tasks {
		if task.ID == t.ID {
			return errors.New("task with this id already exists")
		}
	}

	p.Tasks = append(p.Tasks, t)

	return nil
}

func (p *Project) UpdateTask(task Task) error {
	var isFound bool
	for _, t := range p.Tasks {
		if t.ID == task.ID {
			t.Title = task.Title
			t.Description = task.Description
			isFound = true
		}
	}

	if !isFound {
		return errors.New("the task is not found")
	}

	return nil
}

func (t *Task) Close(status bool) error {
	if !t.Status {
		return errors.New("the task is not active")
	}

	t.Status = status

	return nil
}

func (t *Task) UpdateDescription(description string) error {
	if !t.Status {
		return errors.New("the task is not active")
	}
	if description == "" {
		return errors.New("empty description")
	}

	t.Description = description

	return nil
}

func (p Project) FilterTasksByStatus(status bool) []Task {
	var res []Task
	for _, task := range p.Tasks {
		if status == task.Status {
			res = append(res, task)
		}
	}
	return res
}

func (p Project) PrintInfo() {
	fmt.Printf("%+v\n", p)
}
