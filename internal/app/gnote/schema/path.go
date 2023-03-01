package schema

import (
	"github.com/LiaoSirui/GNote/internal/app/gnote/model"
	"github.com/graphql-go/graphql"
)

var pathType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Path",
		Description: "Path Model",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"fullPath": &graphql.Field{
				Type: graphql.String,
			},
			"subPaths": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"subDocs": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
		},
	},
)

var pathQuery = graphql.Field{
	Name:        "QueryPath",
	Description: "Query Path",
	Type:        pathType,
	Args: graphql.FieldConfigArgument{
		"fullPath": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		fullPath, _ := p.Args["fullPath"].(string)
		return (&model.Path{}).Query(fullPath)
	},
}
