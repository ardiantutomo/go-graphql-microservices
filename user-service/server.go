package main

import (
	"context"
	"fmt"
	"log"
	"user-service/controller"
	"user-service/kafka_handler"
	"user-service/model"
	"user-service/repository"
	"user-service/service"
)

func main() {

	db, _ := model.DBConnection()
	userRepository := repository.NewUserRepository(db)

	// if err := userRepository.Migrate(); err != nil {
	// 	log.Fatal("User migrate err", err)
	// }
	userService := service.NewUserService(userRepository)
	kafkaWriter := kafka_handler.NewKafkaWriter("user-topic-response")
	userController := controller.NewUserController(userService, *kafkaWriter)
	kafkaReader := kafka_handler.NewKafkaReader("user-topic")
	fmt.Println("User service started...")
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error:", err)
			break
		}

		switch key := m.Key; string(key) {
		case "get-all-user":
			userController.GetAllUser()
		case "create-user":
			userController.AddUser(m.Value)
			// default:
			// 	fmt.Println("Not found!")
		}

	}

	if err := kafkaReader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
