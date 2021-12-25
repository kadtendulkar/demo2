package schemas

import (
	"github.com/graphql-go/graphql"
)

var SignUpInput = graphql.FieldConfigArgument{
	"username": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
var SignUpOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SignUpDetails",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
