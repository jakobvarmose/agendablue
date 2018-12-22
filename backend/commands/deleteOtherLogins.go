package commands

import (
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/login"
	"github.com/jinzhu/gorm"
)

func DeleteOtherLogins(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.FormValue("token")

		if token == "" {
			sendErrorJSON(w, http.StatusBadRequest, "token required")
			return
		}

		var l login.Login
		err := s.DB.Where("id=?", token).First(&l).Error
		if err == gorm.ErrRecordNotFound {
			sendJSON(w, http.StatusBadRequest, map[string]interface{}{
				"error": "login doesn't exist",
			})
			return
		}
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		err = s.DB.Where("user_id=? and id<>?", l.UserID, token).Delete(&login.Login{}).Error
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		sendJSON(w, http.StatusOK, map[string]interface{}{})
	}
}
