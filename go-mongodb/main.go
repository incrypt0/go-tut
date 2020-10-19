package main

import (
	"net/http"

	"github.com/incrypt0/go-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/", index)

	//added route + parameter
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	_ = http.ListenAndServe("localhost:8080", r)

}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Mongo</title>
</head>
<body>
    <a href="/user/32789432">Goto blah</a>
</body>
</html>
	`
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(s))
}
