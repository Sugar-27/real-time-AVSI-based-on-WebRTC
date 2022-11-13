package framework

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", entry)
}

// 定义通用的Action接口
type ActionInterface interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

// 定义全局的路由表
var GActionRouter map[string]ActionInterface = make(map[string]ActionInterface)

func responseError(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%d - %s", status, err)))
}

func entry(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
		return
	}
	fmt.Println("=============", r.URL.Path)

	if action, ok := GActionRouter[r.URL.Path]; ok {
		r.ParseForm()
		if action != nil {
			action.Execute(w, r)
		} else {
			responseError(w, r, http.StatusInternalServerError, "Internal server error")
		}
	} else {
		responseError(w, r, http.StatusNotFound, "Not found")
	}
}

func StartHttp() error {
	fmt.Println("start http")
	return http.ListenAndServe(":8080", nil)
}
