package commands

import (
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/item"
	"github.com/jakobvarmose/agendablue/backend/login"
	"github.com/jinzhu/gorm"
)

func ReadItem(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.FormValue("token")
		id := req.FormValue("id")

		if token == "" {
			sendErrorJSON(w, http.StatusBadRequest, "token required")
			return
		}

		if id == "" {
			sendErrorJSON(w, http.StatusBadRequest, "id required")
			return
		}

		var l login.Login
		err := s.DB.Where("id=?", token).First(&l).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				sendJSON(w, http.StatusBadRequest, map[string]interface{}{
					"error": "login doesn't exist",
				})
				return
			}
			sendInternalServerError(w, err)
			return
		}

		var item item.Item
		err = s.DB.Where("id=? and user_id=?", id, l.UserID).First(&item).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				sendJSON(w, http.StatusBadRequest, map[string]interface{}{
					"error": "item doesn't exist",
				})
				return
			}
			sendInternalServerError(w, err)
			return
		}

		sendJSON(w, http.StatusOK, map[string]interface{}{
			"data": item.Data,
		})
	}
}
