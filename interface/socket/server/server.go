package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/hadihammurabi/cacing/interface/socket"
	"github.com/hadihammurabi/cacing/storages"
	"github.com/hadihammurabi/cacing/storages/mapstruct"
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
	log.Println("Starting cacing server on " + socket.ConnHost + ":" + config.Port)
	l, err := net.Listen(socket.ConnType, socket.ConnHost+":"+config.Port)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Error:", err.Error())
			return nil
		}

		go handleConnection(config, c)
	}
}

func handleConnection(config *Config, conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("Client left.")
		conn.Close()
		return
	}

	clientMessage := string(buffer[:len(buffer)-1])
	command := socket.NewCommandFromMessage(clientMessage)

	user := strings.Split(command.User, " ")
	err = authenticateClient(config, user[0], user[1])
	if err != nil {
		replySignal := fmt.Sprintf("error=>%s\n", err.Error())
		conn.Write([]byte(replySignal))
	}

	if command.Type == socket.SignalConnect {
		log.Println("New client connected.")
		replySignal := fmt.Sprintf("success=>connected\n")
		conn.Write([]byte(replySignal))
	} else if command.Type == socket.SignalExec {
		exec := socket.NewExecFromCommandPayload(command.Payload)
		switch exec.Type {
		case socket.ExecSet:
			log.Printf("SET %s %s\n", exec.Args[0], exec.Args[1])
			store.Set(exec.Args[0], exec.Args[1], 0)
			conn.Write([]byte("\n"))
		case socket.ExecGet:
			val, err := store.Get(exec.Args[0])
			if err != nil {
				log.Println(err)
			} else {
				replySignal := fmt.Sprintf("success=>%v\n", val)
				conn.Write([]byte(replySignal))
			}
		}
	}

	handleConnection(config, conn)
}

func authenticateClient(config *Config, username string, password string) error {
	fmt.Println(config)
	if username == config.Username && password == config.Password {
		return nil
	}

	return errors.New("invalid username or password")
}
