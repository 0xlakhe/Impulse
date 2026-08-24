package conversation

import (
	"net/http"

	"github.com/0xlakhe/Impluse/internal/auth"
	"github.com/0xlakhe/Impluse/internal/httpx"
)

type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) CreateOrGet(w http.ResponseWriter, r *http.Request){
	sellerID:=r.PathValue("sellerId")

	if sellerID==""{
		httpx.Error(w,http.StatusBadRequest,"seller id is required")
		return
	}
	userID,ok:=auth.UserIDFromContext(r.Context())

	if !ok{
		httpx.Error(w,http.StatusUnauthorized,ErrUnauthorized.Error())
		return
	}
	
	// var req CreateConversationRequest

	// if err:=json.NewDecoder(r.Body).Decode(&req); err!=nil{
	// 	httpx.Error(w,http.StatusBadRequest,ErrBadRequest.Error())
	// 	return
	// }
	conversaton,err:=h.service.CreateOrGet(r.Context(),userID,sellerID)
	if err!=nil{
		httpx.Error(w,http.StatusInternalServerError,err.Error())
		return
	}
	httpx.JSON(w,http.StatusOK,CreateConversationResponse{
		ConversationID: conversaton.ID,
	})
}