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

	if err = template.Execute(os.Stdout, variables, nil); err != nil {
		panic(err)
	}
}
