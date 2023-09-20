package main

import (
	gamev1 "42tokyo-road-to-dojo-go/gen/game/v1"
	"42tokyo-road-to-dojo-go/gen/game/v1/v1connect"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGameServer_Create(t *testing.T) {
	game := &GameServer{}
	path, handler := v1connect.NewUserServiceHandler(game)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	server := httptest.NewServer(mux)
	defer server.Close()

	client := v1connect.NewUserServiceClient(&http.Client{}, server.URL)

	req := &connect.Request[gamev1.CreateRequest]{
		Msg: &gamev1.CreateRequest{
			UserName: "test_user",
		},
	}

	res, err := client.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to create: %v", err)
	}

	fmt.Printf("Status: %s\n", res.Msg.Status)
	fmt.Printf("Message: %s\n", res.Msg.Message)
}
