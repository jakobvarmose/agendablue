package commands

import (
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/item"
	"github.com/jakobvarmose/agendablue/backend/login"
	"github.com/jinzhu/gorm"
)

func UpdateItem(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.FormValue("token")
		id := req.FormValue("id")
		data := req.FormValue("data")

		if token == "" {
			sendErrorJSON(w, http.StatusBadRequest, "token required")
			return
		}

		if id == "" {
			sendErrorJSON(w, http.StatusBadRequest, "id required")
			return
		}

		if data == "" {
			sendErrorJSON(w, http.StatusBadRequest, "data required")
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

		item.Data = data

		err = s.DB.Save(&item).Error
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		sendJSON(w, http.StatusOK, map[string]interface{}{})
	}
}
