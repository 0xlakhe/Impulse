package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/0xlakhe/Impluse/internal/httpx"
)

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request){
	var req RegisterRequest
	if err:=json.NewDecoder(r.Body).Decode(&req); err!=nil{
		httpx.Error(w,http.StatusBadRequest,err.Error())
		return
	}
	response,err:=h.service.Register(r.Context(),req)
	
	if err!=nil{
		switch{
		case errors.Is(err, ErrEmailAlreadyExists):
			http.Error(w,err.Error(),http.StatusConflict)
			return
		
		case errors.Is(err,ErrUserNameTaken):
			http.Error(w,err.Error(),http.StatusConflict)
			return
		default:
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}
	}
	httpx.JSON(w,http.StatusCreated,response)
}
