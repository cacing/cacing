package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
)

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
		// message := strings.Split(rawMessage, "=>")
		log.Println(rawMessage)

		// if message[0] == "success" {
		// 	log.Println("> ", message)
		// } else if message[0] == "error" {
		// 	log.Println("Error > ", message)
		// }

		// fmt.Print("Text to send: ")
		// input, _ := reader.ReadString('\n')
		// conn.Write([]byte(input))
	}

}
