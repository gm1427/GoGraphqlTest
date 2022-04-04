package main

import (
	"context"
	"fmt"
	s "strings"

	"github.com/machinebox/graphql"
)

func main() {
	graphqlClient := graphql.NewClient("https://gitlab.com/api/graphql")
	graphqlRequest := graphql.NewRequest(
		`
		query ($ln: Int=5}){
			projects(last:$ln){
			  nodes{
				name
				description
				forksCount
			  }
			}
		  }
	`)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	process(graphqlResponse)
}
func process(in interface{}) {
	var names string = ""
	var sum int = 0
	for key, element := range in.(map[string]interface{})["projects"].(map[string]interface{})["nodes"].([]interface{}) {
		name := fmt.Sprint(element.(map[string]interface{})["name"])
		if key == 0 {
			names = name
		} else {
			names = s.Join([]string{names, name}, ",")
		}
		j := int(element.(map[string]interface{})["forksCount"].(float64))
		sum = sum + j
	}
	println("names: \n", names)
	println("sum: \n", sum)
}
