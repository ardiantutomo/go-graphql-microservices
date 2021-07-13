package controller

import (
	"auth-service/kafka_handler"

	"github.com/segmentio/kafka-go"
)

// AuthController : represent the auth's controller
type AuthController interface {
	Login()
}

type authController struct {
	kafkaWriter kafka.Writer
}

//NewUserController -> returns new user controller
func NewAuthController(kafkaWriter kafka.Writer) AuthController {
	return authController{
		kafkaWriter: kafkaWriter,
	}
}

func (u authController) Login() {

	kafka_handler.SendToKafka(u.kafkaWriter, []byte("response"), userData.Bytes())
}
