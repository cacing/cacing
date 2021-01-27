package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/hadihammurabi/cacing/interface/socket"
	uuid "github.com/satori/go.uuid"
)

func clientCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "SET", Description: "Set value of a key"},
		{Text: "GET", Description: "Get value from a key"},
		{Text: "DEL", Description: "Delete key and value"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// ConnectTo func
func ConnectTo(url *url.URL) error {
	conn, err := net.Dial(socket.ConnType, socket.ConnHost+":"+url.Port())
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	password, _ := url.User.Password()
	connectCommand := socket.CommandToMessage(&socket.Command{
		Type: socket.SignalConnect,
		User: fmt.Sprintf("%s %s", url.User.Username(), password),
	})
	_, err = conn.Write([]byte(connectCommand))
	if err != nil {
		return err
	}

	// reader := bufio.NewReader(os.Stdin)

	var id uuid.UUID

	for {
		rawMessage, _ := bufio.NewReader(conn).ReadString('\n')
		commandFromServer := socket.NewCommandFromMessage(rawMessage)

		switch commandFromServer.Type {
		case socket.SignalSuccess:
			if commandFromServer.Payload == "login" {
				id = uuid.FromStringOrNil(commandFromServer.User)
				fmt.Println("Connected with id:", id)
			} else if commandFromServer.Payload == string(socket.ExecSet) {
			} else {
				fmt.Println(commandFromServer.Payload)
			}
			fmt.Printf("%s\n\n", commandFromServer.Headers["TIME"])
		case socket.SignalError:
			log.Fatalln(commandFromServer.Payload)
		}

		input := prompt.Input(">>> ", clientCompleter)
		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			os.Exit(0)
		}

		signal := socket.CommandToMessage(&socket.Command{
			Type:    socket.SignalExec,
			User:    id.String(),
			Payload: input,
		})
		conn.Write([]byte(signal))

		// fmt.Print("Text to send: ")
		// input, _ := reader.ReadString('\n')
	}

}
