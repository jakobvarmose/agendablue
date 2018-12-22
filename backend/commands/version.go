package commands

import (
	"net/http"
)

func Version(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		sendJSON(w, http.StatusOK, map[string]interface{}{
			"version": "1",
		})
	}
}

func version(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"version": "1",
	}
}
