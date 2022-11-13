package action

import (
	"fmt"
	"net/http"
	"text/template"
)

type xrtcClientPushAction struct{}

func NewXrtcClientPushAction() *xrtcClientPushAction {
	return &xrtcClientPushAction{}
}

func writeHtmlErrorResponse(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(err))
}


func (*xrtcClientPushAction) Execute(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hello xrtcclient push action")
	t, err := template.ParseFiles("./static/template/push.tpl")
	if err != nil {
		fmt.Println(err)
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}

	request := make(map[string]string)

	for k, v := range r.Form {
		request[k] = v[0]
	}

	t.Execute(w, request)
	if err != nil {
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}
}

