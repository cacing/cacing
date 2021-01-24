package socket

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/hadihammurabi/cacing/storages"
	"github.com/hadihammurabi/cacing/storages/mapstruct"
)

const (
	connHost = "localhost"
	connType = "tcp"
)

var store storages.Storage = mapstruct.NewMapStruct(map[string]mapstruct.Data{})

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
			replySignal := fmt.Sprintf("error=>%s\n", err.Error())
			conn.Write([]byte(replySignal))
		} else {
			fmt.Println("New connection")
			replySignal := fmt.Sprintf("success=>connected\n")
			conn.Write([]byte(replySignal))
		}
	} else if clientMessage[0] == "exec" {
		user := strings.Split(clientMessage[1], " ")
		err := authenticateClient(user[0], user[1])
		if err != nil {
			replySignal := fmt.Sprintf("error=>%s\n", err.Error())
			conn.Write([]byte(replySignal))
		} else {
			command := strings.Split(clientMessage[2], " ")
			if command[0] == "SET" {
				log.Printf("SET %s %s\n", command[1], command[2])
				store.Set(command[1], command[2], 0)
			} else if command[0] == "GET" {
				val, err := store.Get(command[1])
				log.Printf("GET %s, GOT %v\n", command[1], val)
				if err != nil {
					replySignal := fmt.Sprintf("success=>%v\n", val)
					conn.Write([]byte(replySignal))
				}
			}
		}
	}

	handleConnection(conn)
}

func authenticateClient(username string, password string) error {
	return nil
}
