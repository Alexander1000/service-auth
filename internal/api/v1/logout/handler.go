package logout

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	jsonResponse "github.com/Alexander1000/service-auth/internal/response/json"
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

	err := h.storage.Logout(req.Context(), reqData.Token)
	if err != nil {
		jsonResponse.Reply(resp, jsonResponse.ErrorInternalServerError, http.StatusInternalServerError)
	}

	jsonResponse.Reply(resp, response{Result: resultSuccess{Success: true}}, http.StatusOK)
}
