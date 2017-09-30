package main

import (
	"golang.org/x/oauth2"
	"os"
	"fmt"
	"context"
	"github.com/shurcooL/githubql"
	"encoding/json"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubql.NewClient(httpClient)

	type Language struct {
		Name  githubql.String   `json:"name"`
		Color githubql.String   `json:"color"`
	}

	type Repository struct {
		NameWithOwner   githubql.String `json:"owner"`
		Name    githubql.String `json:"name"`
		Url     githubql.String `json:"url"`
		Languages struct {
			Nodes []struct {
				Language `graphql:"... on Language"`
			}
		} `graphql:"languages(first: 5)"`
	}

	var query struct {
		Search struct {
			Nodes []struct {
				Repository `graphql:"... on Repository"`
			}
		} `graphql:"search(first: 5, query: $q, type: $searchType)"`
	}

	variables := map[string]interface{}{
		"q": githubql.String("GraphQL"),
		"searchType":  githubql.SearchTypeRepository,
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(query.Search.Nodes)
	fmt.Printf("%s", bytes)
}

