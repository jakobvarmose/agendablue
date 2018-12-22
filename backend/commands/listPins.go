package commands

import (
	"crypto/subtle"
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

func listPins(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	username, _ := obj["username"].(string)

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

	var pins []user.Pin
	err = s.DB.Where("user_id=?", u.ID).Find(&pins).Error
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	pins2 := make([]map[string]interface{}, len(pins))
	for i, pin := range pins {
		pins2[i] = map[string]interface{}{
			"ipnsKey": pin.IPNSKey,
		}
	}

	return http.StatusOK, map[string]interface{}{
		"pins": pins2,
	}
}
