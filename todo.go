package main

import (
	"errors"
	"strconv"
	"time"
)

type Todo struct {
	Title       string     `json:"Title"`
	Completed   bool       `json:"Completed"`
	AddedAt     time.Time  `json:"AddedAt"`
	CompletedAt *time.Time `json:"CompletedAt"`
}

type Todos []Todo

// validate index(index int)
// add(name string)
// delete(index int)
// rename(index int, newName string)
// toggle(index int)

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Invalid index: " + strconv.Itoa(index))
	}
	return nil
}

func (todos *Todos) add(todoName string) error {

	if todoName == "" {
		return errors.New("Invalid name of TODO")
	}

	for _, todo := range *todos {
		if todo.Title == todoName {
			return errors.New("Todo with this name already exist")
		}
	}

	todo := Todo{
		Title:       todoName,
		Completed:   false,
		AddedAt:     time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)

	return nil
}

func (todos *Todos) delete(index int) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[:index-1]...)
	return nil
}

func (todos *Todos) rename(index int, newName string) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = newName

	return nil
}

func (todos *Todos) toggle(index int) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	status := (*todos)[index].Completed
	if !status {
		completedTime := time.Now()
		(*todos)[index].CompletedAt = &completedTime
	} else {
		(*todos)[index].CompletedAt = nil
	}

	(*todos)[index].Completed = !status
	return nil
}
