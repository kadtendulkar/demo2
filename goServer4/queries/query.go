package queries

import (
	SignUpResolver "server/resolvers"
	schema "server/schemas"

	"github.com/graphql-go/graphql"
)

var rootQueryyy = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "rootQueryyy",
		Fields: graphql.Fields{

			"SignUpMemberList": &graphql.Field{
				Type:        graphql.NewList(schema.SignUpOutput),
				Description: "List of members",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return SignUpResolver.SignUpMemberList, nil
				},
			},
		},
	})
