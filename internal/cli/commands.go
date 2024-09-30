package cli

import (
	"context"
	"fmt"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"os"
)

const incorrectInput = "incorrect input data, see prompt"

func (e *Executor) HandleCommands(raw string) {
	data := e.parseRawData(raw)
	command := data[0]
	args := data[1:]

	switch command {
	case "signUp":
		if len(args) != 2 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.SignUp(
			args[0], args[1],
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("success")
		return

	case "signIn":
		if len(args) != 2 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.SignIn(
			args[0], args[1],
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		go func() {
			changes, err := e.client.SubscribeToChanges(context.Background())
			if err != nil {
				return
			}
			for {
				recv, err := changes.Recv()
				if err != nil {
					fmt.Println("Ошибка при получении сообщения:", err)
					return
				}

				fmt.Printf("Warning: new action \"%s\" on product \"%s\" with id \"%s\"\n",
					recv.Action, recv.Product, recv.Id)
			}
		}()
		fmt.Println("success")
		return

	case "get-cards":
		result, err := e.client.GetCards(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range *result {
			fmt.Printf("%#v\n", val)
		}
		return
	case "create-card":
		if len(args) < 4 {
			fmt.Println(incorrectInput)
			return
		}
		result, err := e.client.CreateCard(context.Background(),
			args[0],
			args[1],
			args[2],
			args[3],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", *result)
		return
	case "update-card":
		if len(args) < 5 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.UpdateCard(context.Background(),
			args[0],
			args[1],
			args[2],
			args[3],
			args[4],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", models.Card{
			ID:             args[0],
			Number:         args[1],
			ExpirationDate: args[2],
			HolderName:     args[3],
			CVV:            args[4],
		})
		return
	case "delete-card":
		if len(args) < 1 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.DeleteCard(context.Background(),
			args[0],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		return

	case "upload-file":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		result, err := e.client.UploadFile(context.Background(), args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", *result)
		return
	case "get-files":
		result, err := e.client.GetFiles(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range *result {
			fmt.Printf("%#v\n", val)
		}
		return
	case "download-file":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.DownloadFile(context.Background(), args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		return
	case "delete-file":
		if len(args) < 1 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.DeleteFile(context.Background(),
			args[0],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		return

	case "get-text":
		result, err := e.client.GetTextData(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range *result {
			fmt.Printf("%#v\n", val)
		}
		return
	case "create-text":
		if len(args) < 2 {
			fmt.Println(incorrectInput)
			return
		}
		result, err := e.client.CreateTextData(context.Background(),
			args[0],
			args[1],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", *result)
		return
	case "update-text":
		if len(args) < 3 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.UpdateTextData(context.Background(),
			args[0],
			args[1],
			args[2],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", models.TextData{
			ID:   args[0],
			Meta: args[1],
			Text: args[2],
		})
		return
	case "delete-text":
		if len(args) < 1 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.DeleteCard(context.Background(),
			args[0],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		return

	case "get-auth":
		result, err := e.client.GetAuthData(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, val := range *result {
			fmt.Printf("%#v\n", val)
		}
		return
	case "create-auth":
		if len(args) < 3 {
			fmt.Println(incorrectInput)
			return
		}
		result, err := e.client.CreateAuthData(context.Background(),
			args[0],
			args[1],
			args[2],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", *result)
		return
	case "update-auth":
		if len(args) < 4 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.UpdateAuthData(context.Background(),
			args[0],
			args[1],
			args[2],
			args[3],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%#v\n", models.AuthData{
			ID:       args[0],
			Meta:     args[1],
			Login:    args[2],
			Password: args[3],
		})
		return
	case "delete-auth":
		if len(args) < 1 {
			fmt.Println(incorrectInput)
			return
		}
		err := e.client.DeleteAuthData(context.Background(),
			args[0],
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		return

	case "exit":
		fmt.Println("Good bye")
		os.Exit(1)
		return
	}
}
