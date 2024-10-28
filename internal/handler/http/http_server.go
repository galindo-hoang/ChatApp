package http

import (
	"context"
	"fmt"
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/logic"
	"io"
	"net/http"
)

type HttpServer interface {
	Start(ctx context.Context)
}

type httpServer struct {
	accountLogic logic.Account
	configs      configs.Config
}

func NewHttpServer(accountLogic logic.Account, configs configs.Config) HttpServer {
	return &httpServer{
		configs:      configs,
		accountLogic: accountLogic,
	}
}

func (h *httpServer) Start(ctx context.Context) {
	http.HandleFunc("/v1/sessions", createSession)
	http.HandleFunc("/v1/accounts", registerAccount)
	//fmt.Printf("hello world %v\n", h.configs.Http.Port)
	err := http.ListenAndServe(h.configs.Http.Address, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func registerAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start to register account")

	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, "register account success")
}

func createSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start to create session")
	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, "create session success")
}
