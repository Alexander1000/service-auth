package registration

import (
	"encoding/json"
	"github.com/Alexander1000/service-auth/internal/model"
	jsonResponse "github.com/Alexander1000/service-auth/internal/response/json"
	"io/ioutil"
	"log"
	"net/http"
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

	credentials := make([]model.Credential, 0, len(reqData.Credentials))
	for _, cred := range reqData.Credentials {
		credentials = append(
			credentials,
			model.Credential{Type: cred.Type, ID: cred.ID},
		)
	}

	err := h.storage.Registration(req.Context(), reqData.UserID, reqData.Password, credentials)
	if err != nil {
		log.Printf("storage err: %v", err)
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
		return
	}

	jsonResponse.Reply(resp, response{Result: responseSuccess{Success: true}}, http.StatusOK)
}
