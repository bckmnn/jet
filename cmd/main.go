package main

import (
	"github.com/CloudyKit/jet/v6"
	"os"
	"reflect"
)

type User struct {
	Name string
}

var (
	variables = map[string]reflect.Value{
		"user": reflect.ValueOf(User{
			Name: "vlad",
		}),
		"m": reflect.ValueOf(map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
		}),
	}
)

func main() {
	set := jet.NewSet(
		jet.NewOSFileSystemLoader("./cmd/testData"),
	)

	template, err := set.GetTemplate("map.jet")
	if err != nil {
		panic(err)
	}

	if err = template.Execute(os.Stdout, variables, map[string]interface{}{
		"Name": "vlad",
		"Name2": map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
		},
	}); err != nil {
		panic(err)
	}
}
