package httpx

import "net/http"

type ErrorResponse struct{
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, status int,message string){
	JSON(w,status,ErrorResponse{
		Error:message,
	})
}