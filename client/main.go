package main

import (
	"context"
	"fmt"
	"todo-grpc/todo/models"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceTodo() models.TodoClient {
	port := ":7000"

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	return models.NewTodoClient(conn)
}

func main() {
	// todo1 := models.CreateTodoItemRequest{
	// 	Title: "todo 1",
	// }

	todo := serviceTodo()

	fmt.Println("================ test todo")

	// m, err := todo.CreateTodoItem(context.Background(), &todo1)

	// if err != nil {
	// 	log.Debug().Err(err).Msg(err.Error())
	// }
	// m_arr := m.String()
	// log.Info().Msg("tes: " + m_arr)

	m, err := todo.ListTodoItems(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
	}
	m_arr := m.String()
	log.Info().Msg("tes: " + m_arr)

	id := "01HFM4M867K53QFTPMXYYDR35Z"

	req := models.UpdateTodoItemRequest{
		Id:    id,
		Title: "todo 2 edit",
	}

	n, err := todo.UpdateTodoItem(context.Background(), &req)
	if err != nil {
		log.Debug().Err(err).Msg(err.Error())
	}
	n_arr := n.String()
	log.Info().Msg("tes: " + n_arr)
}
