package handlers

import (
	"fmt"
	"net/http"
)

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dsd")
}
