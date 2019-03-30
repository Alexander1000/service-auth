package authorize

import (
	"encoding/json"
	jsonResponse "github.com/Alexander1000/service-auth/internal/response/json"
	"github.com/Alexander1000/service-auth/internal/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	storage storageRepository
}

func New(storage storageRepository) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != "POST" {
		jsonResponse.Reply(resp, jsonResponse.ErrorNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	reqData := request{}
	if rawReqData, err := ioutil.ReadAll(req.Body); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	} else if err = json.Unmarshal(rawReqData, &reqData); err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	res, err := h.storage.Authorize(req.Context(), reqData.Token)
	if err != nil {
		log.Printf("storage err: %v", err)
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	switch res {
	case storage.AuthRefreshed:
		fallthrough
	case storage.AuthDisabled:
		fallthrough
	case storage.AuthNotFound:
		jsonResponse.Reply(resp, jsonResponse.ErrorNotFound, http.StatusOK)
		return
	case storage.AuthExpired:
		respErr := jsonResponse.ErrorBadRequest
		respErr.Error.Message = "Token expired"
		jsonResponse.Reply(resp, respErr, http.StatusOK)
		return
	}

	jsonResponse.Reply(resp, response{Result: resultSuccess{Success: true}}, http.StatusOK)
}
