package commands

import (
	"crypto/subtle"
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

func readUserBootstrap(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	username, _ := obj["username"].(string)
	loginKey, _ := obj["loginKey"].([]byte)

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

	login := user.Login{
		LoginKey: loginKey,
		UserID:   u.ID,
	}

	err = s.DB.Create(&login).Error
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, map[string]interface{}{
		"bootstrap": u.Bootstrap,
		"content":   u.Content,
		"loginId":   login.ID,
		"loginKey":  login.LoginKey,
	}
}
