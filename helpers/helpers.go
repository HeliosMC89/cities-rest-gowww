package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	JSONResponse struct {
		Error bool        `json:"error,omitempty"`
		Code  int         `json:"code,omitempty"`
		Msg   string      `json:"msg,omitempty"`
		Data  interface{} `json:"data,omitempty"`
	}

	DBUpdate struct {
		ID       int64 `json:"id"`
		Affected int64 `json:"affected"`
	}
)

func WriteJSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	_, err = w.Write(resp)
	return err
}

func ReadJSON(r *http.Request, payload interface{}) error {
	buf, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, payload)
}
