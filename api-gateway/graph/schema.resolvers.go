package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-gateway/graph/generated"
	"api-gateway/graph/model"
	"api-gateway/kafka_handler"
	"bytes"
	"context"
	"encoding/gob"
	"log"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	var userData bytes.Buffer
	enc := gob.NewEncoder(&userData)
	err := enc.Encode(input)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	log.Print("[CreateUser]...Sending request")
	kafka_handler.SendToKafka("user-topic", "create-user", userData.String())
	kafkaReader := kafka_handler.NewKafkaReader("user-topic-response", "user-group")
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error:", err)
			break
		}
		if string(m.Key) == "create-user-error" || string(m.Key) == "create-user-response" {
			return string(m.Value), err
		}
	}
	return "-", err
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	log.Print("[GetAllUsers]...Sending request")
	kafka_handler.SendToKafka("user-topic", "get-all-user", "-")

	kafkaReader := kafka_handler.NewKafkaReader("user-topic-response", "user-group")
	var err error
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error:", err)
			break
		}
		if string(m.Key) == "get-all-user-response" {
			data := bytes.NewBuffer(m.Value)
			dec := gob.NewDecoder(data)
			var users []model.User
			// r.users = []*model.User{}
			err = dec.Decode(&users)
			userStruct := make([]*model.User, len(users))

			for i := range userStruct {
				userStruct[i] = &users[i]
			}
			r.users = userStruct
			return r.users, err
		}
	}
	return r.users, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
