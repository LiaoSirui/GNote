package schema

import (
	"github.com/LiaoSirui/GNote/internal/app/gnote/model"
	"github.com/graphql-go/graphql"
)

var docType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Doc",
		Description: "Doc Model",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"fullPath": &graphql.Field{
				Type: graphql.String,
			},
			"contents": &graphql.Field{
				Type: graphql.String,
			},
			"contentsSplit": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
		},
	},
)

var docQuery = graphql.Field{
	Name:        "QueryPath",
	Description: "Query Path",
	Type:        docType,
	Args: graphql.FieldConfigArgument{
		"fullPath": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		fullPath, _ := p.Args["fullPath"].(string)
		return (&model.Doc{}).Query(fullPath)
	},
}
