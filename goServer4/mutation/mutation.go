package mutation

import (
	"context"
	"log"
	mongoConnector "server/data"
	SignUpResolver "server/resolvers"
	schema "server/schemas"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"SignUp": &graphql.Field{
			Type:        schema.SignUpOutput,
			Description: "SignUp Member Details",
			Args:        schema.SignUpInput,
			Resolve:     SignUpResolver.SignUp,
		},
	},
})

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "rootQuery",
		Fields: graphql.Fields{

			"SignUpMemberList": &graphql.Field{
				Type:        graphql.NewList(schema.SignUpOutput),
				Description: "List of members",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					client := mongoConnector.MongoConnector()
					collection := client.Database("members").Collection("golangMembers")
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					cur, err := collection.Find(ctx, bson.D{})
					if err != nil {
						log.Fatal(err)
					}
					var members []*SignUpResolver.SignUpMember
					for cur.Next(ctx) {
						var member *SignUpResolver.SignUpMember
						err := cur.Decode(&member)
						if err != nil {
							log.Fatal(err)
						}
						members = append(members, member)
					}
					return members, nil
					//return SignUpResolver.SignUpMemberList, nil
				},
			},
		},
	})
