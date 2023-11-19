package main

import (
	"context"
	"todo-grpc/todo"
	"todo-grpc/todo/models"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func listItems(ctx context.Context) (models.TodoItems, error) {
	tx, err := todo.DB.Begin()

	if err != nil {
		return models.TodoItems{}, err
	}

	list, err := findAllItems(ctx, tx)

	if err != nil {
		return models.TodoItems{}, err
	}

	tx.Commit()

	return list, nil
}

func createItem(ctx context.Context, title string) (models.TodoItem, error) {
	tx, err := todo.DB.Begin()
	if err != nil {
		return models.TodoItem{}, err
	}
	
	todoItem, err := todo.MakeNewItem(title)

	if err != nil {
		return models.TodoItem{}, err
	}

	err = todo.SaveItem(ctx, tx, todoItem)

	if err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	err = tx.Commit()

	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	isDone := todoItem.IsDone()
	todo := models.TodoItem{
		Id:        todoItem.Id.String(),
		Title:     todoItem.Title,
		CreatedAt: timestamppb.New(todoItem.CreatedAt),
		DoneAt:    timestamppb.New(todoItem.DoneAt.Time),
		IsDone:    &isDone,
	}

	return todo, nil
}

func findItem(ctx context.Context, id ulid.ULID) (models.TodoItem, error) {
	tx, err := todo.DB.Begin()

	if err != nil {
		return models.TodoItem{}, err
	}

	item, err := todo.FindItemById(ctx, tx, id)

	if err != nil {
		return models.TodoItem{}, err
	}

	err = tx.Commit()

	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	isDone := item.IsDone()
	todo := models.TodoItem{
		Id:        item.Id.String(),
		Title:     item.Title,
		CreatedAt: timestamppb.New(item.CreatedAt),
		DoneAt:    timestamppb.New(item.DoneAt.Time),
		IsDone:    &isDone,
	}

	return todo, nil
}

func makeItemDone(ctx context.Context, id ulid.ULID) error {
	tx, err := todo.DB.Begin()

	if err != nil {
		return err
	}

	item, err := todo.FindItemById(ctx, tx, id)

	if err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return err
	}

	if err = item.MakeDone(); err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return err
	}

	if err = todo.SaveItem(ctx, tx, item); err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return err
	}

	err = tx.Commit()

	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
		return err
	}

	return nil

}

func updateItem(ctx context.Context, id ulid.ULID, title string) (models.TodoItem, error) {
	tx, err := todo.DB.Begin()

	if err != nil {
		return models.TodoItem{}, err
	}

	item, err := todo.FindItemById(ctx, tx, id)

	if err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	item.Title = title

	if err = todo.SaveItem(ctx, tx, item); err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	err = tx.Commit()

	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
		return models.TodoItem{}, err
	}

	isDone := item.IsDone()
	todo := models.TodoItem{
		Id:        item.Id.String(),
		Title:     item.Title,
		CreatedAt: timestamppb.New(item.CreatedAt),
		DoneAt:    timestamppb.New(item.DoneAt.Time),
		IsDone:    &isDone,
	}

	return todo, nil
}

func deleteItem(ctx context.Context, id ulid.ULID) error {
	tx, err := todo.DB.Begin()

	if err != nil {
		return err
	}

	_, err = todo.DeleteItemById(ctx, tx, id)

	if err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return err
	}
	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
