package untils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, payload any) error {

	if r.Body == nil {
		return fmt.Errorf("Request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

func WriteImage(w http.ResponseWriter, status int, image []byte) {
	w.Header().Add("Content-Type", "image/png")
	w.WriteHeader(status)
	w.Write(image)
}

// func getTime() string {
// 	now := time.Now()
// 	timeZone, _ := now.Zone()
// 	return now.Format("2006-01-02 15:04:05") + " " + timeZone
// }
