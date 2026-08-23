package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/0xlakhe/Impluse/internal/httpx"
	"github.com/0xlakhe/Impluse/internal/user"
)


type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
	var req LoginRequest
	if err:=json.NewDecoder(r.Body).Decode(&req); err!=nil{
		httpx.Error(w,http.StatusBadRequest,err.Error())
		return 
	}

	response,err:=h.service.Login(r.Context(),req)
	if err!=nil{
		switch{
		case errors.Is(err,user.ErrInvalidCredentials):
			httpx.Error(w,http.StatusUnauthorized,err.Error())
		default:
			httpx.Error(w,http.StatusInternalServerError,err.Error())
		}
		return
	}
	httpx.JSON(w,http.StatusOK,response)
}

func(h *Handler) Me(w http.ResponseWriter,r *http.Request,){
	userID,ok:=UserIDFromContext(r.Context())
	if !ok{
		httpx.Error(w,http.StatusUnauthorized,"user not found in context")
		return
	}
	httpx.JSON(w,http.StatusOK,map[string]string{
		"user_id":userID,
	})
}