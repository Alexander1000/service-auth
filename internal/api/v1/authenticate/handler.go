package authenticate

import (
	"net/http"
	jsonResponse "github.com/Alexander1000/service-auth/internal/response/json"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Handler struct {
	storage storageRepository
}

func New(storage storageRepository) *Handler {
	return &Handler{
		storage: storage,
	}
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

	token, err := h.storage.Authenticate(req.Context(), reqData.Credential, reqData.Password)
	if err != nil {
		log.Printf("storage err: %v", err)
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	jsonResponse.Reply(resp, response{Result: *token}, http.StatusOK)
}