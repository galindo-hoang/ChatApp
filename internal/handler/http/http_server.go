package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ChatService/internal/configs"
	"github.com/ChatService/internal/logic"
	"io"
	"net/http"
	"time"
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

func enableCors(funcHandler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	fmt.Println("set up")
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		r.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		r.Header.Set("Content-Type", "application/json")
		funcHandler(w, r)
	}
}

func (h *httpServer) Start(ctx context.Context) {

	http.HandleFunc("/v1/sessions", enableCors(h.createSession))
	http.HandleFunc("/v1/sessions/verify", enableCors(h.verifySession))
	http.HandleFunc("/v1/accounts", enableCors(h.createAccount))

	fmt.Printf("listenning address %v\n", h.configs.Http.Address)
	if err := http.ListenAndServe(h.configs.Http.Address, nil); err != nil {
		fmt.Println(err)
	}
}

func jsonResHttp[T any](w http.ResponseWriter, statusCode int, res Response[T]) {
	if statusCode >= 400 {
		fmt.Println(res.Message)
	}
	parsedRes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	w.WriteHeader(statusCode)
	io.WriteString(w, string(parsedRes))
}

func jsonReqHttp[T any](r *http.Request, data *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (h *httpServer) createAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResHttp(w, http.StatusMethodNotAllowed, Response[any]{
			Data:    nil,
			Success: false,
			Message: "method not allowed",
		})
		return
	}

	var parsedBody logic.CreateAccountParams
	if err := jsonReqHttp(r, &parsedBody); err != nil {
		jsonResHttp(w, http.StatusNotAcceptable, Response[any]{
			Data:    nil,
			Success: false,
			Message: "body format error",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := h.accountLogic.CreateAccount(ctx, parsedBody)
	if err != nil {
		jsonResHttp(w, http.StatusBadRequest, Response[any]{
			Data:    err,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	jsonResHttp(w, http.StatusAccepted, Response[logic.CreateAccountResponse]{
		Data:    res,
		Success: true,
		Message: "",
	})
}

func (h *httpServer) createSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResHttp(w, http.StatusMethodNotAllowed, Response[any]{
			Data:    nil,
			Success: false,
			Message: "method not allowed",
		})
		return
	}

	var parsedBody logic.CreateSessionParams
	if err := jsonReqHttp(r, &parsedBody); err != nil {
		jsonResHttp(w, http.StatusNotAcceptable, Response[any]{
			Data:    nil,
			Success: false,
			Message: "body format error",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := h.accountLogic.CreateSession(ctx, parsedBody)
	if err != nil {
		jsonResHttp(w, http.StatusBadRequest, Response[any]{
			Data:    err,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	jsonResHttp(w, http.StatusAccepted, Response[logic.CreateSessionResponse]{
		Data:    res,
		Success: true,
		Message: "",
	})
}

func (h *httpServer) verifySession(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	w.Header().Set("Content-Type", "application/json")

	//(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if r.Method != "GET" {
		jsonResHttp(w, http.StatusMethodNotAllowed, Response[any]{
			Data:    nil,
			Success: false,
			Message: "method not allowed",
		})
		return
	}

	var session = r.Header.Get("session")
	if session == "" {
		jsonResHttp(w, http.StatusUnauthorized, Response[any]{
			Data:    nil,
			Success: true,
			Message: "",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := h.accountLogic.ValidateSession(ctx, session); err != nil {
		jsonResHttp(w, http.StatusUnauthorized, Response[any]{
			Data:    nil,
			Success: true,
			Message: err.Error(),
		})
	} else {
		jsonResHttp(w, http.StatusAccepted, Response[any]{
			Data:    nil,
			Success: true,
			Message: "",
		})
	}
}
