package routes

import (
	"github.com/zhongwq/TestDocker/service"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"path/filepath"
)

type resInfo struct {
	Msg string
}

func NewServer() *negroni.Negroni {
	formatter := render.New();
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)

	n.UseHandler(mx);
	return n;
}


func initRoutes(mx *mux.Router, formatter *render.Render) {
	fmt.Println("Init Routes")
	mx.HandleFunc("/users", UserRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/users/{name:[_a-zA-Z0-9]+}", GetUserByNameHandler(formatter)).Methods("POST")
}

func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.PostForm[`username`][0] == "" {
			formatter.JSON(w,404,"Username can't be null")
		}

		res, msg := service.UserRegister(req.PostForm)
		fmt.Println("msg", msg)
		if res == true {
			formatter.JSON(w,200, resInfo{"Successfully create user!"})
		} else {
			formatter.JSON(w,404,resInfo{msg})
		}
	}
}

func GetUserByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Get by name & passwd")
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		fmt.Println(name)
		reqContent, msg := service.GetUserInfo(req.PostForm)
		fmt.Println(reqContent)
		if reqContent.GetName() != "" {
			r.JSON(w, 200, reqContent)
		} else {
			r.JSON(w,404, resInfo{msg})
		}
	}
}