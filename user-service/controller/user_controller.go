package controller

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"user-service/kafka_handler"
	"user-service/model"
	"user-service/service"

	"github.com/segmentio/kafka-go"
)

// UserController : represent the user's controller contract
type UserController interface {
	AddUser([]byte)
	GetAllUser()
}

type userController struct {
	userService service.UserService
	kafkaWriter kafka.Writer
}

//NewUserController -> returns new user controller
func NewUserController(s service.UserService, kafkaWriter kafka.Writer) UserController {
	return userController{
		userService: s,
		kafkaWriter: kafkaWriter,
	}
}

func (u userController) GetAllUser() {

	users, err := u.userService.GetAll()
	if err != nil {
		kafka_handler.SendToKafka(u.kafkaWriter, []byte("get-all-user-error"), []byte("Failed to get all user"))
		return
	}
	var userData bytes.Buffer
	enc := gob.NewEncoder(&userData)
	err = enc.Encode(users)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	log.Print("[UserController]...Sending response")
	fmt.Println(userData.Bytes())
	kafka_handler.SendToKafka(u.kafkaWriter, []byte("get-all-user-response"), userData.Bytes())
}

func (u userController) AddUser(userData []byte) {
	log.Print("[UserController]...add User")
	var user model.User
	data := bytes.NewBuffer(userData)
	dec := gob.NewDecoder(data)
	err := dec.Decode(&user)
	user, err = u.userService.Save(user)
	if err != nil {
		kafka_handler.SendToKafka(u.kafkaWriter, []byte("create-user-error"), []byte("Failed to create user"))
		return
	}
	kafka_handler.SendToKafka(u.kafkaWriter, []byte("create-user-response"), []byte("Create user success"))
}
