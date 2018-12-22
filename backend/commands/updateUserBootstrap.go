package commands

import (
	"crypto/subtle"
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/login"
	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

func updateUserBootstrap(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	username, _ := obj["username"].(string)
	info, _ := obj["info"].([]byte)
	bootstrap, _ := obj["bootstrap"].([]byte)
	accessKey, _ := obj["accessKey"].([]byte)

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

	err = s.DB.Where("user_id=?", u.ID).Delete(&login.Login{}).Error
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	u.Info = info
	u.Bootstrap = bootstrap
	u.AccessKey = accessKey

	err = s.DB.Save(&u).Error
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, map[string]interface{}{}
}
