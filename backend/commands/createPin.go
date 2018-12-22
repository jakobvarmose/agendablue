package commands

import (
	"net/http"
)

func createPin(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	/*username, _ := obj["username"].(string)
	ipnsKey, _ := obj["ipnsKey"].(string)

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

	p, err := user.NewPin(u.ID, ipnsKey)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	err = s.DB.Create(p).Error
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, map[string]interface{}{}*/
	return http.StatusNotImplemented, nil
}
