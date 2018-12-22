package commands

import (
	"crypto/subtle"
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

func readUserContent(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	username, _ := obj["username"].(string)

	var u user.User
	err := s.DB.Where("username=?", username).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusBadRequest, "user doesn't exist"
		}
		return http.StatusInternalServerError, err.Error()
	}

	if subtle.ConstantTimeCompare(publicKey, u.ContentKey) != 1 {
		return http.StatusBadRequest, "wrong publicKey"
	}

	return http.StatusOK, map[string]interface{}{
		"content": u.Content,
	}
}
