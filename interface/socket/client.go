package socket

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
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
	conn, err := net.Dial(connType, connHost+":"+url.Port())
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	password, _ := url.User.Password()
	connectSignal := fmt.Sprintf("connect=>%s %s\n", url.User.Username(), password)
	_, err = conn.Write([]byte(connectSignal))
	if err != nil {
		return err
	}

	// reader := bufio.NewReader(os.Stdin)

	for {
		rawMessage, _ := bufio.NewReader(conn).ReadString('\n')
		message := strings.Split(rawMessage, "=>")

		if message[0] == "success" {
			fmt.Println(message[1])
		} else if message[0] == "error" {
			panic(message[1])
		}

		input := prompt.Input(">>> ", clientCompleter)
		if input == "EXIT" {
			os.Exit(0)
		}

		signal := fmt.Sprintf("exec=>%s %s=>%s\n", url.User.Username(), password, input)
		conn.Write([]byte(signal))

		// fmt.Print("Text to send: ")
		// input, _ := reader.ReadString('\n')
	}

}
