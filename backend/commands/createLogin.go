package commands

import (
	"net/http"
)

func createLogin(s *State, publicKey []byte, obj map[string]string) (int, interface{}) {
	/*
		username := obj["username"]

		var u user.User
		err := s.DB.Where("username=?", username).First(&u).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return http.StatusBadRequest, "user doesn't exist"
			}
			return http.StatusInternalServerError, err.Error()
		}

		if subtle.ConstantTimeCompare(publicKey, u.AccessKey) != 1 {
			return http.StatusBadRequest, "wrong publicKey"
		}

		l, err := login.New(u.ID)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		err = s.DB.Create(l).Error
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		return http.StatusOK, map[string]interface{}{
			"token": l.ID,
		}*/
	return http.StatusNotImplemented, nil
}

/*
func CreateLogin(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		username := req.FormValue("username")
		accesskey := req.FormValue("accesskey")

		if username == "" {
			sendErrorJSON(w, http.StatusBadRequest, "username required")
			return
		}

		if accesskey == "" {
			sendErrorJSON(w, http.StatusBadRequest, "accesskey required")
			return
		}

		var u user.User
		err := s.DB.Where("username=?", username).First(&u).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				sendErrorJSON(w, http.StatusBadRequest, "user doesn't exist")
				return
			}
			sendInternalServerError(w, err)
			return
		}

		if !u.CheckAccessKey(accesskey) {
			sendErrorJSON(w, http.StatusForbidden, "invalid accesskey")
			return
		}

		l, err := login.New(u.ID)
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		err = s.DB.Create(l).Error
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		sendJSON(w, http.StatusOK, map[string]interface{}{
			"token": l.ID,
		})
	}
}
*/
