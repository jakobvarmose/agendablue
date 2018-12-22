package commands

import (
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

func readUserInfo(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	username, _ := obj["username"].(string)

	var u user.User
	err := s.DB.Where("username=?", username).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusBadRequest, "user doesn't exist"
		}
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, map[string]interface{}{
		"info": u.Info,
	}
}
