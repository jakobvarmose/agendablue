package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jakobvarmose/agendablue/backend/commands"
	"github.com/jakobvarmose/agendablue/backend/item"
	"github.com/jakobvarmose/agendablue/backend/user"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Data struct {
	gorm.Model
	Data   []byte
	UserID uint
}

func getenv(name string) string {
	val := os.Getenv(name)
	if val != "" {
		return val
	}
	filename := os.Getenv(name + "_FILE")
	if filename != "" {
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		return string(buf)
	}
	return ""
}

func main() {
	DB_HOST := getenv("DB_HOST")
	DB_USER := getenv("DB_USER")
	DB_PASSWORD := getenv("DB_PASSWORD")
	DB_NAME := getenv("DB_NAME")
	DOMAIN := getenv("DOMAIN")

	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+")/"+DB_NAME+"?charset=utf8mb4&parseTime=true")
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(time.Second)
	}

	err = db.AutoMigrate(&user.User{}, &user.Login{}, &Data{}, &item.Item{}).Error
	if err != nil {
		log.Println(err)
		return
	}

	err = db.Exec(`
			create view if not exists usersview as select
				id,created_at, updated_at, username,
				length(access_key) access_key,
				length(content_key) content_key,
				length(info) info,
				length(bootstrap) bootstrap,
				length(content) content
			from users;
		`).Error
	if err != nil {
		log.Println(err)
		return
	}

	s := &commands.State{
		DB:     db,
		Domain: DOMAIN,
	}

	api := http.NewServeMux()

	api.HandleFunc("/version", commands.Version(s))

	//api.HandleFunc("/createUser", commands.CreateUser(s))
	//api.HandleFunc("/readUserInfo", commands.ReadUserInfo(s))
	//api.HandleFunc("/readUserBootstrap", commands.ReadUserBootstrap(s))
	//api.HandleFunc("/updateUser", commands.UpdateUser(s))
	//api.HandleFunc("/deleteUser", commands.DeleteUser(s))

	//api.HandleFunc("/createLogin", commands.CreateLogin(s))
	api.HandleFunc("/deleteLogin", commands.DeleteLogin(s))
	api.HandleFunc("/deleteOtherLogins", commands.DeleteOtherLogins(s))

	api.HandleFunc("/createItem", commands.CreateItem(s))
	api.HandleFunc("/readItem", commands.ReadItem(s))
	api.HandleFunc("/updateItem", commands.UpdateItem(s))
	api.HandleFunc("/deleteItem", commands.DeleteItem(s))
	api.HandleFunc("/listItems", commands.ListItems(s))

	api.HandleFunc("/signed", commands.Signed(s))

	app := http.NewServeMux()

	app.Handle("/api/v0/", http.StripPrefix("/api/v0", api))

	app.Handle("/", http.FileServer(http.Dir("dist")))

	lst, err := net.Listen("tcp", ":8041")
	if err != nil {
		log.Println(err)
		return
	}
	defer lst.Close()

	err = http.Serve(lst, app)
	if err != nil {
		log.Println(err)
		return
	}
}
