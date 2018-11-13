package controllers

import (
	"fmt"
	"net/http"
	"wonder_go_df/application/services"
)

func Index(w http.ResponseWriter, r *http.Request) {
	services.HoHello("index")
	fmt.Fprint(w, "Wonder Welcome You!")
}
