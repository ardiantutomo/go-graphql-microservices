package main

import (
	"auth-service/controller"
	"auth-service/kafka_handler"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"context"
	"log"
)

func main() {

	db, _ := model.DBConnection()
	userRepository := repository.NewUserRepository(db)

	// if err := userRepository.Migrate(); err != nil {
	// 	log.Fatal("User migrate err", err)
	// }
	authService := service.NewUserService(userRepository)
	kafkaWriter := kafka_handler.NewKafkaWriter("user-topic-response")
	authController := controller.NewUserController(userService, *kafkaWriter)
	kafkaReader := kafka_handler.NewKafkaReader("user-topic")
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error:", err)
			break
		}

		switch key := m.Key; string(key) {
		case "get-all-user":
			userController.GetAllUser()
			// default:
			// 	fmt.Println("Not found!")
		}

	}

	if err := kafkaReader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
