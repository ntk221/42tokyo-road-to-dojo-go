package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"

	gamev1 "42tokyo-road-to-dojo-go/gen/game/v1"
	"42tokyo-road-to-dojo-go/gen/game/v1/v1connect"
)

type GameServer struct{}

func (s *GameServer) Create(
	ctx context.Context,
	req *connect.Request[gamev1.CreateRequest],
) (*connect.Response[gamev1.CreateResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&gamev1.CreateResponse{
		Status: "OK!",
		Token:  fmt.Sprintf("Hello, %s", req.Msg.UserName),
	})
	res.Header().Set("Game-Version", "v1")
	return res, nil
}

func main() {
	game := &GameServer{}
	mux := http.NewServeMux()
	path, handler := v1connect.NewUserServiceHandler(game)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
