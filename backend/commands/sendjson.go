package commands

import (
	"encoding/base64"
	"log"
	"net/http"

	cbor "github.com/whyrusleeping/cbor/go"
)

func sendJSON(w http.ResponseWriter, status int, val interface{}) {
	/*buf, err := json.Marshal(val)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/
	src, err := cbor.Dumps(val)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(src))+1)
	base64.RawURLEncoding.Encode(dst, src)
	dst[len(dst)-1] = '\n'

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write(dst)
}

func sendErrorJSON(w http.ResponseWriter, status int, err interface{}) {
	sendJSON(w, status, map[string]interface{}{
		"error": err,
	})
}

func sendInternalServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	sendErrorJSON(w, http.StatusInternalServerError, "internal server error")
}
