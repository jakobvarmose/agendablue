package commands

import (
	"fmt"
	"net/http"

	"github.com/jakobvarmose/agendablue/backend/user"
)

func createUser(s *State, publicKey []byte, obj map[string]interface{}) (int, interface{}) {
	fmt.Println(obj)
	username, _ := obj["username"].(string)
	accessKey, _ := obj["accessKey"].([]byte)
	info, _ := obj["info"].([]byte)
	bootstrap, _ := obj["bootstrap"].([]byte)
	contentKey, _ := obj["contentKey"].([]byte)
	content, _ := obj["content"].([]byte)
	loginKey, _ := obj["loginKey"].([]byte)

	u := user.User{
		Username: username,

		AccessKey:  accessKey,
		ContentKey: contentKey,

		Info:      info,
		Bootstrap: bootstrap,
		Content:   content,
	}

	err := s.DB.Create(&u).Error
	if err != nil {
		err = s.DB.Where("username=?", username).First(&user.User{}).Error
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusBadRequest, "user already exists"
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
		"loginId":  login.ID,
		"loginKey": login.LoginKey,
	}
}
