package resolvers

import (
	"context"
	"fmt"
	"log"
	mongoConnector "server/data"
	"time"

	"github.com/graphql-go/graphql"
)

var SignUpMemberList []SignUpMember

type SignUpMember struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(params graphql.ResolveParams) (interface{}, error) {
	// marshall and cast the argument value
	username, _ := params.Args["username"].(string)
	//fmt.Println(text)
	// figure out new id
	password, _ := params.Args["password"].(string)

	// perform mutation operation here
	// for e.g. create a Todo and save to DB.
	newSignUpMember := SignUpMember{
		Username: username,
		Password: password,
	}

	SignUpMemberList = append(SignUpMemberList, newSignUpMember)
	fmt.Println(SignUpMemberList)

	client := mongoConnector.MongoConnector()
	collection := client.Database("members").Collection("golangMembers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, newSignUpMember)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(collection)
	// return &model.Member{
	// 	ID:       res.InsertedID.(primitive.ObjectID).Hex(),
	// 	Username: input.Username,
	// 	Password: input.Password,
	// }

	// return the new Todo object that we supposedly save to DB
	// Note here that
	// - we are returning a `Todo` struct instance here
	// - we previously specified the return Type to be `todoType`
	// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
	return newSignUpMember, nil
}
