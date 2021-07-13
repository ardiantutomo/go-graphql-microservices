package controller

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"
	"role-permission-service/kafka_handler"
	"role-permission-service/model"
	"role-permission-service/service"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// PermissionController : represent the permission's controller contract
type PermissionController interface {
	AddPermission(*gin.Context)
	GetAllPermission()
}

type permissionController struct {
	permissionService service.PermissionService
	kafkaWriter       kafka.Writer
}

//NewPermissionController -> returns new permission controller
func NewPermissionController(s service.PermissionService, kafkaWriter kafka.Writer) PermissionController {
	return permissionController{
		permissionService: s,
		kafkaWriter:       kafkaWriter,
	}
}

func (u permissionController) GetAllPermission() {

	permissions, err := u.permissionService.GetAll()
	if err != nil {
		kafka_handler.SendToKafka(u.kafkaWriter, []byte("error"), []byte("Failed to get all permission"))
		return
	}
	var permissionData bytes.Buffer
	enc := gob.NewEncoder(&permissionData)
	err = enc.Encode(permissions)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	log.Print("[PermissionController]...Sending response")
	log.Print("total: ", len(permissionData.Bytes()))
	kafka_handler.SendToKafka(u.kafkaWriter, []byte("response"), permissionData.Bytes())
}

func (u permissionController) AddPermission(c *gin.Context) {
	log.Print("[PermissionController]...add Permission")
	var permission model.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission, err := u.permissionService.Save(permission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": permission})
}
