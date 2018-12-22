package commands

import (
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/item"
	"github.com/jakobvarmose/agendablue/backend/login"
	"github.com/jinzhu/gorm"
)

func CreateItem(s *State) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.FormValue("token")
		data := req.FormValue("data")

		if token == "" {
			sendErrorJSON(w, http.StatusBadRequest, "token required")
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
				sendErrorJSON(w, http.StatusBadRequest, "login doesn't exist")
				return
			}
			sendInternalServerError(w, err)
			return
		}

		item := item.Item{
			UserID: l.UserID,
			Data:   data,
		}

		err = s.DB.Create(&item).Error
		if err != nil {
			sendInternalServerError(w, err)
			return
		}

		sendJSON(w, http.StatusOK, map[string]interface{}{
			"id": item.ID,
		})
	}
}
