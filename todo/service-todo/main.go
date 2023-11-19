package main

import (
	"context"
	"net"
	"todo-grpc/todo"
	"todo-grpc/todo/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoServer struct {
	models.UnimplementedTodoServer
}

func (TodoServer) GetTodoItem(ctx context.Context, param *models.TodoItemId) (*models.TodoItem, error) {
	itemId, err := ulid.Parse(param.Id)

	if err != nil {
		return &models.TodoItem{}, err
	}

	item, err := findItem(ctx, itemId)

	if err != nil {
		return &models.TodoItem{}, err
	}

	return &item, nil
}

func (TodoServer) ListTodoItems(ctx context.Context, void *emptypb.Empty) (*models.TodoItems, error) {
	todoItems, err := listItems(ctx)

	if err != nil {
		return nil, err
	}

	return &todoItems, nil
}

func (TodoServer) CreateTodoItem(ctx context.Context, param *models.CreateTodoItemRequest) (*models.TodoItem, error) {
	todoItem, err := createItem(ctx, param.Title)

	if err != nil {
		return &models.TodoItem{}, err
	}

	log.Info().Msg("create new item: " + todoItem.Id)

	return &todoItem, nil
}

func (TodoServer) UpdateTodoItem(ctx context.Context, param *models.UpdateTodoItemRequest) (*models.TodoItem, error) {
	itemId, err := ulid.Parse(param.Id)

	if err != nil {
		return &models.TodoItem{}, err
	}

	todoItem, err := updateItem(ctx, itemId, param.Title)

	if err != nil {
		return &models.TodoItem{}, err
	}

	return &todoItem, nil
}

func (TodoServer) DeleteTodoItem(ctx context.Context, param *models.TodoItemId) (*emptypb.Empty, error) {
	itemId, err := ulid.Parse(param.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	err = deleteItem(ctx, itemId)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (TodoServer) MakeTodoItemDone(ctx context.Context, param *models.TodoItemId) (*emptypb.Empty, error) {
	itemId, err := ulid.Parse(param.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	err = makeItemDone(ctx, itemId)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func main() {
	todo.SetupDatabase()

	srv := grpc.NewServer()

	var todoSrv TodoServer
	models.RegisterTodoServer(srv, todoSrv)

	log.Info().Msg("Starting Todo RPC Server at :7000")

	l, err := net.Listen("tcp", ":7000")

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start server at :7000")
	}

	log.Fatal().Err(srv.Serve(l))
}
