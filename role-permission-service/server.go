package main

import (
	"context"
	"log"
	"role-permission-service/controller"
	"role-permission-service/kafka_handler"
	"role-permission-service/model"
	"role-permission-service/repository"
	"role-permission-service/service"
)

func main() {

	db, _ := model.DBConnection()
	permissionRepository := repository.NewPermissionRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	rolePermissionRepository := repository.NewRoleHasPermissionRepository(db)
	modelPermissionRepository := repository.NewModelHasPermissionRepository(db)
	modelRoleRepository := repository.NewModelHasRoleRepository(db)

	if err := permissionRepository.Migrate(); err != nil {
		log.Fatal("Permission migrate err", err)
	}
	if err := roleRepository.Migrate(); err != nil {
		log.Fatal("Role migrate err", err)
	}
	if err := rolePermissionRepository.Migrate(); err != nil {
		log.Fatal("Role migrate err", err)
	}
	if err := modelPermissionRepository.Migrate(); err != nil {
		log.Fatal("Role migrate err", err)
	}
	if err := modelRoleRepository.Migrate(); err != nil {
		log.Fatal("Role migrate err", err)
	}

	permissionService := service.NewPermissionService(permissionRepository)
	kafkaWriter := kafka_handler.NewKafkaWriter("role-permission-topic-response")
	permissionController := controller.NewPermissionController(permissionService, *kafkaWriter)
	kafkaReader := kafka_handler.NewKafkaReader("role-permission-topic")
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error:", err)
			break
		}

		switch key := m.Key; string(key) {
		case "get-all-permission":
			permissionController.GetAllPermission()
			// default:
			// 	fmt.Println("Not found!")
		}

	}

	if err := kafkaReader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
