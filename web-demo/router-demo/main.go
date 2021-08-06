package main

import (
	"net/http"
	"router-demo/controller"
)

func main()  {
	controller.RegisterRouters()
	http.ListenAndServe(":8080",nil)
}
