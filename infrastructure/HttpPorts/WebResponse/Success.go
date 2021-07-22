package webresponse

import (
	"encoding/json"
	"net/http"
)

type ResponseApi struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

func Succes(result interface{}, status int) *ResponseApi {

	return &ResponseApi{
		Success: true,
		Status:  status,
		Result:  result,
	}

}

func (r *ResponseApi) Send(w http.ResponseWriter) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)

}
