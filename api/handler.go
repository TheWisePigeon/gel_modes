package api

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "encoding/json"
)

type JsonResponse struct {
    Message string
    HttpCode int
    Failed bool
}


func Handler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "<h1>Hi from Go :)</h1>")
	// u, err := json.Marshal(JsonResponse{Message: "Hello gopher :)", HttpCode: 200, Failed: false})
    //     if err != nil {
    //         panic(err)
    //     }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(u) 
}