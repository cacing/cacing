package socket

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	connHost = "localhost"
	connType = "tcp"
)

// Config type
type Config struct {
	Port     string
	Username string
	Password string
}

// NewConfig return new server Config
// or error if fail
func NewConfig(port string, username string, password string) (*Config, error) {
	if port == "" {
		port = "8080"
	}

	if username == "" {
		return nil, errors.New("Username can't be blank")
	}

	if password == "" {
		return nil, errors.New("Password can't be blank")
	}

	config := &Config{
		Port:     port,
		Username: username,
		Password: password,
	}
	return config, nil

}

// RunServer func
func RunServer(config *Config) error {
	fmt.Println("Starting cacing server on " + connHost + ":" + config.Port)
	l, err := net.Listen(connType, connHost+":"+config.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error:", err.Error())
			return nil
		}

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	clientMessage := strings.Split(string(buffer[:len(buffer)-1]), "=>")
	if clientMessage[0] == "connect" {
		user := strings.Split(clientMessage[1], " ")
		err := authenticateClient(user[0], user[1])
		if err != nil {
			replySignal := fmt.Sprintf("error=>%s", err.Error())
			conn.Write([]byte(replySignal))
		} else {
			fmt.Println("New connection")
			replySignal := fmt.Sprintf("success=>connected")
			conn.Write([]byte(replySignal))
		}
	}

	handleConnection(conn)
}

func authenticateClient(username string, password string) error {
	return nil
}
