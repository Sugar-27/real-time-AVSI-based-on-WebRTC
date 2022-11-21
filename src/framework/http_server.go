package framework

import (
	"fmt"
	"net/http"
	"signaling/src/glog"
	"strconv"
)

func init() {
	http.HandleFunc("/", entry)
}

// 定义通用的Action接口
type ActionInterface interface {
	Execute(w http.ResponseWriter, cr *ComRequest)
}

// 定义全局的路由表
var GActionRouter map[string]ActionInterface = make(map[string]ActionInterface)

type ComRequest struct {
	R      *http.Request
	Logger *ComLog
	LogId  uint32
}

func responseError(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%d - %s", status, err)))
}

func getRealClientIP(r *http.Request) string {
	ip := r.RemoteAddr

	if rip := r.Header.Get("X-Real-IP"); rip != "" {
		ip = rip
	} else if rip = r.Header.Get("X-Forwarded-IP"); rip != "" {
		ip = rip
	}

	return ip
}

func entry(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
		return
	}
	fmt.Println("=============", r.URL.Path)

	if action, ok := GActionRouter[r.URL.Path]; ok {
		if action != nil {
			cr := &ComRequest{
				R:      r,
				Logger: &ComLog{},
				LogId:  GetLogId32(),
			}
			cr.Logger.AddNotice("logID", strconv.Itoa(int(cr.LogId)))
			cr.Logger.AddNotice("url", r.URL.Path)
			cr.Logger.AddNotice("referer", r.Header.Get("Referer"))
			cr.Logger.AddNotice("cookie", r.Header.Get("Cookie"))
			cr.Logger.AddNotice("ua", r.Header.Get("User-Agent"))
			cr.Logger.AddNotice("clientIP", r.RemoteAddr)
			cr.Logger.AddNotice("realClientIP", getRealClientIP(r))
			r.ParseForm()

			// 添加请求的参数到log中
			for k, v := range r.Form {
				cr.Logger.AddNotice(k, v[0])
			}
			
			cr.Logger.TimeBegin("totalCost")
			action.Execute(w, cr)
			cr.Logger.TimeEnd("totalCost")
			cr.Logger.Infof("")
		} else {
			responseError(w, r, http.StatusInternalServerError, "Internal server error")
		}
	} else {
		responseError(w, r, http.StatusNotFound, "Not found")
	}
}

func RegisterStaticUrl() {
	fs := http.FileServer(http.Dir(gconf.httpStaticDir))
	http.Handle(gconf.httpStaticPrefix, http.StripPrefix(gconf.httpStaticPrefix, fs))
}

func StartHttp() error {
	// fmt.Println("start http")
	glog.Infof("start http server on port: %d", gconf.httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", gconf.httpPort), nil)
}

func StartHttps() error {
	// fmt.Println("start http")
	glog.Infof("start https server on port: %d", gconf.httpsPort)
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", gconf.httpsPort), gconf.httpsCert, gconf.httpsKey, nil)
}
