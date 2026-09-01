package conversation

import (
	"encoding/json"
	"errors"
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
	
	var req CreateConversationRequest

	if err:=json.NewDecoder(r.Body).Decode(&req); err!=nil{
		httpx.Error(w,http.StatusBadRequest,ErrBadRequest.Error())
		return
	}
	conversaton,err:=h.service.CreateOrGet(r.Context(),userID,sellerID,*req.ProductID)
	if err!=nil{
		httpx.Error(w,http.StatusInternalServerError,err.Error())
		return
	}
	httpx.JSON(w,http.StatusOK,CreateConversationResponse{
		ConversationID: conversaton.ID,
	})
}

func (h *Handler) SendMessage(w http.ResponseWriter,r *http.Request){
	conversationID:=r.PathValue("conversationId")
	userID,ok:=auth.UserIDFromContext(r.Context(),)
	if !ok{
		httpx.Error(w,http.StatusUnauthorized,"authentication required")
		return
	}
	
	var req SendMessageRequest
	err:=json.NewDecoder(r.Body).Decode(&req)

	if err!=nil{
		httpx.Error(w,http.StatusBadRequest,err.Error())
		return
	}
	message,err:=h.service.SendMessage(r.Context(),conversationID,userID,req.Content)
	if err!=nil{
		switch{
		case errors.Is(err,ErrConversationNotFound):
			httpx.Error(w,http.StatusNotFound,err.Error())
		case errors.Is(err,ErrEmptyMessage):
			httpx.Error(w,http.StatusBadRequest,err.Error())
		default:
			httpx.Error(w,http.StatusInternalServerError,err.Error())
		}
		return
	}
	httpx.JSON(w,http.StatusCreated,message)
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request){
	conversationID:=r.PathValue("conversationId")
	if conversationID==""{
		httpx.Error(w,http.StatusBadRequest,"conversation id is required")
		return
	}
	userID,ok:=auth.UserIDFromContext(r.Context())
	if !ok{
		httpx.Error(w,http.StatusUnauthorized,"authentication required")
		return
	}
	messages,err:=h.service.GetMessags(r.Context(),conversationID,userID)
	if err!=nil{
		switch{
		case errors.Is(err,ErrConversationNotFound):
			httpx.Error(w,http.StatusNotFound,err.Error())
		default:
			httpx.Error(w,http.StatusInternalServerError,err.Error())
		}
		return
	}
	response:=make([]MessageResponse,0,len(messages))
	
	for _,message:=range messages{
		response=append(response, toMessageResponse(message))
		httpx.JSON(w,http.StatusOK,response,)
	}
	httpx.JSON(w,http.StatusOK,response)
}