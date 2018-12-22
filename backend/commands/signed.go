package commands

import (
	"encoding/base64"
	"net/http"
	"time"

	cbor "github.com/whyrusleeping/cbor/go"
	"golang.org/x/crypto/nacl/sign"
)

func Signed(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dataString := req.FormValue("data")
		dataBytes, err := base64.RawURLEncoding.DecodeString(dataString)
		if err != nil {
			sendErrorJSON(w, http.StatusBadRequest, "invalid base64 in data")
			return
		}

		var data map[string]interface{}
		err = cbor.Loads(dataBytes, &data)
		if err != nil {
			sendErrorJSON(w, http.StatusBadRequest, "invalid cbor in data")
			return
		}

		publicKey, ok := data["publicKey"].([]byte)
		if !ok {
			sendErrorJSON(w, http.StatusForbidden, "no publicKey")
			return
		}

		if len(publicKey) != 32 {
			sendErrorJSON(w, http.StatusForbidden, "invalid publicKey")
			return
		}

		var publicKeyArray [32]byte
		copy(publicKeyArray[:], publicKey)

		signedMessage, ok := data["signedMessage"].([]byte)
		if !ok {
			sendErrorJSON(w, http.StatusForbidden, "no signedMessage")
			return
		}

		message, ok := sign.Open(nil, signedMessage, &publicKeyArray)
		if !ok {
			sendErrorJSON(w, http.StatusForbidden, "invalid signature")
			return
		}

		var obj map[string]interface{}
		err = cbor.Loads(message, &obj)
		if err != nil {
			sendErrorJSON(w, http.StatusForbidden, "invalid message")
			return
		}

		objDomain, _ := obj["domain"].(string)
		if objDomain != s.Domain {
			sendErrorJSON(w, http.StatusForbidden, "wrong domain")
			return
		}

		objTime, _ := obj["time"].(string)
		t, err := time.Parse(time.RFC3339, objTime)
		if err != nil {
			sendErrorJSON(w, http.StatusForbidden, "invalid time")
			return
		}

		dur, _ := time.ParseDuration("30m")
		if t.Before(time.Now().Add(-dur)) || t.After(time.Now().Add(dur)) {
			sendErrorJSON(w, http.StatusForbidden, "wrong time")
			return
		}

		var status int
		var res interface{}

		switch obj["action"] {
		case "version":
			status, res = version(s, publicKey, obj)

		case "createUser":
			status, res = createUser(s, publicKey, obj)
		case "readUserInfo":
			status, res = readUserInfo(s, publicKey, obj)
		case "readUserBootstrap":
			status, res = readUserBootstrap(s, publicKey, obj)
		case "readUserContent":
			status, res = readUserContent(s, publicKey, obj)
		case "updateUser":
			status, res = updateUser(s, publicKey, obj)
		case "updateUserBootstrap":
			status, res = updateUserBootstrap(s, publicKey, obj)
		case "updateUserContent":
			status, res = updateUserContent(s, publicKey, obj)
		case "deleteUser":
			status, res = deleteUser(s, publicKey, obj)

		case "createPin":
			status, res = createPin(s, publicKey, obj)
		case "deletePin":
			status, res = deletePin(s, publicKey, obj)
		case "listPins":
			status, res = listPins(s, publicKey, obj)

		default:
			status, res = http.StatusBadRequest, "action doesn't exist"
		}

		switch res.(type) {
		case string:
			sendJSON(w, status, map[string]interface{}{
				"error": res,
			})
		default:
			sendJSON(w, status, res)
		}
	}
}
