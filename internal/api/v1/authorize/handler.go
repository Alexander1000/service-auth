package authorize

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	jsonResponse "github.com/Alexander1000/service-auth/internal/response/json"
)

type Handler struct {

}

func New() *Handler {
	return &Handler{}
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

	jsonResponse.Reply(resp, response{Result: resultSuccess{Success: true}}, http.StatusOK)
}
