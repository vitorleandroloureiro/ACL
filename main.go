package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
	"github.com/vitorleandroloureiro/ACL/controllers"
	"github.com/vitorleandroloureiro/ACL/models"

)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {

	models.NewDic();

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/user/create/", controllers.CreateHandler)
	log.Fatal(http.ListenAndServe(":9090", router))
}


