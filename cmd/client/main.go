package main

import (
	"context"
	"fmt"
	"github.com/egor-zakharov/goph-keeper/internal/client"
)

func main() {
	c := client.New("localhost:8081")
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	err = c.SignIn("login", "password")
	if err != nil {
		panic(err)
	}
	resp, err := c.UploadFile(context.Background(), "C:\\Users\\edzakharov\\GolandProjects\\1\\goph-keeper\\pkg\\333.exe", "нe")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
