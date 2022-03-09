package main

import (
	"fmt"

	"box/internal/resource"
)

func main() {
	// cmd.Execute()

	fs := resource.NewRepositoryFS("./")

	res, err := fs.ListDir("./")
	if err != nil {
		panic(err)
	}

	for _, r := range res {
		fmt.Println(r)
	}

	fs.MoveDir("./hej", "./hej2")
}
