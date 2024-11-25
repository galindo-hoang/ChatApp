package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ChatService/internal/handler"
	"github.com/ChatService/internal/logic"
	"io"
	"net/http"
	"strings"
	"time"
)

func Middleware(funcHandler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("running....")
		// Set json
		w.Header().Set("Content-Type", "application/json")

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Handle preflight (OPTIONS) request
		if r.Method == "OPTIONS" {
			return
		}
		funcHandler(w, r)
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

func getLogic(w http.ResponseWriter) (logic.Account, func(), error) {
	accountLogic, clearLogic, err := handler.GetAccountLogic("")

	if err != nil {
		fmt.Printf("get accountLogic fail: %v\n", err.Error())
		jsonResHttp(w, http.StatusBadRequest, Response[any]{
			Data:    err,
			Success: false,
			Message: err.Error(),
		})
		clearLogic()
		return nil, nil, err
	}

	return accountLogic, clearLogic, nil
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
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

	accountLogic, clearLogic, err := getLogic(w)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
		clearLogic()
	}()

	res, err := accountLogic.CreateAccount(ctx, parsedBody)
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

func CreateSession(w http.ResponseWriter, r *http.Request) {
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

	accountLogic, clearLogic, err := getLogic(w)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
		clearLogic()
	}()

	res, err := accountLogic.CreateSession(ctx, parsedBody)
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

func VerifySession(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		jsonResHttp(w, http.StatusMethodNotAllowed, Response[any]{
			Data:    nil,
			Success: false,
			Message: "method not allowed",
		})
		return
	}

	var session = strings.Split(r.Header.Get("Authorization"), " ")[1]
	if session == "" {
		jsonResHttp(w, http.StatusUnauthorized, Response[any]{
			Data:    nil,
			Success: true,
			Message: "",
		})
	}

	accountLogic, clearLogic, err := getLogic(w)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
		clearLogic()
	}()

	if err := accountLogic.ValidateSession(ctx, session); err != nil {
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
