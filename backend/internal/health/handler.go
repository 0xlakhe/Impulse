package health

import (
	"encoding/json"
	"net/http"
)

type Respone struct{
	Status string `json:"status"`
	Database string `json:"database"`
}

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service:service}
}

func (h *Handler) Handle(w http.ResponseWriter,r *http.Request){
	response:=Respone{
		Status: "ok",
		Database: "ok",
	}
	if err:=h.service.CheckDatabase(r.Context());err!=nil{
		response.Database="unavailable"
	}
	w.Header().Set("Content-Type","application/json")
	err:=json.NewEncoder(w).Encode(response)
	if err!=nil{
		http.Error(w,"failed to encode response",http.StatusInternalServerError)
		return
	}
}
