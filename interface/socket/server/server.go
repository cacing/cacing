package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cacing/cacing/interface/socket/client"

	"github.com/cacing/cacing/interface/socket"
	"github.com/cacing/cacing/storage"
	"github.com/cacing/cacing/storage/mapstruct"
)

var store storage.Storage = mapstruct.NewMapStruct(map[string]mapstruct.Data{})
var clientPool = client.NewPool()

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

	clientMessage := buffer[:len(buffer)-1]
	command, err := socket.NewCommandFromMessage(clientMessage)
	if err != nil {
		replySignal, _ := socket.CommandToMessage(&socket.Command{
			Type:    socket.SignalError,
			Payload: fmt.Sprintf("%s", err.Error()),
		})
		conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
	}

	resolveCommand(config, conn, command)
	handleConnection(config, conn)
}

func authenticateClient(config *Config, username string, password string) error {
	if username == config.Username && password == config.Password {
		return nil
	}

	return errors.New("invalid username or password")
}

func clientExists(id string) error {
	exists, _ := clientPool.IsExists(id)
	if exists {
		return nil
	}

	return errors.New("invalid client id")
}

func resolveCommand(config *Config, conn net.Conn, command *socket.Command) {
	start := time.Now()
	if command.Type == socket.SignalConnect {
		user := strings.Split(command.User, " ")
		err := authenticateClient(config, user[0], user[1])
		if err != nil {
			finish := time.Since(start)
			replySignal, _ := socket.CommandToMessage(&socket.Command{
				Type:    socket.SignalError,
				Payload: fmt.Sprintf("%s", err.Error()),
				Headers: socket.CommandHeader{
					"TIME": finish.String(),
				},
			})
			conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
		}
		newClientID, _ := clientPool.Add()
		finish := time.Since(start)
		replySignal, _ := socket.CommandToMessage(&socket.Command{
			Type:    socket.SignalSuccess,
			User:    newClientID.String(),
			Payload: "login",
			Headers: socket.CommandHeader{
				"TIME": finish.String(),
			},
		})
		conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
		log.Println("New client connected.")
	} else if command.Type == socket.SignalExec {
		err := clientExists(command.User)
		if err != nil {
			finish := time.Since(start)
			replySignal, _ := socket.CommandToMessage(&socket.Command{
				Type:    socket.SignalError,
				Payload: fmt.Sprintf("%s", err.Error()),
				Headers: socket.CommandHeader{
					"TIME": finish.String(),
				},
			})
			conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
		}
		exec := socket.NewExecFromCommandPayload(command.Payload)
		switch exec.Type {
		case socket.ExecSet:
			log.Printf("SET %s %s\n", exec.Args[0], exec.Args[1])
			store.Set(exec.Args[0], exec.Args[1], 0)
			finish := time.Since(start)
			replySignal, _ := socket.CommandToMessage(&socket.Command{
				Type:    socket.SignalSuccess,
				User:    command.User,
				Payload: string(socket.ExecSet),
				Headers: socket.CommandHeader{
					"TIME": finish.String(),
				},
			})
			conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
		case socket.ExecGet:
			val, err := store.Get(exec.Args[0])
			finish := time.Since(start)
			if err != nil {
				replySignal, _ := socket.CommandToMessage(&socket.Command{
					Type:    socket.SignalError,
					User:    command.User,
					Payload: fmt.Sprintf("%s", err.Error()),
					Headers: socket.CommandHeader{
						"TIME": finish.String(),
					},
				})
				conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
			} else {
				replySignal, _ := socket.CommandToMessage(&socket.Command{
					Type:    socket.SignalSuccess,
					User:    command.User,
					Payload: fmt.Sprintf("%v", val),
					Headers: socket.CommandHeader{
						"TIME": finish.String(),
					},
				})
				conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
			}
		case socket.ExecDel:
			val, err := store.Delete(exec.Args[0])
			finish := time.Since(start)
			if err != nil {
				replySignal, _ := socket.CommandToMessage(&socket.Command{
					Type:    socket.SignalError,
					User:    command.User,
					Payload: fmt.Sprintf("%s", err.Error()),
					Headers: socket.CommandHeader{
						"TIME": finish.String(),
					},
				})
				conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
			} else {
				replySignal, _ := socket.CommandToMessage(&socket.Command{
					Type:    socket.SignalSuccess,
					User:    command.User,
					Payload: fmt.Sprintf("%v", val),
					Headers: socket.CommandHeader{
						"TIME": finish.String(),
					},
				})
				conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
			}
		case socket.ExecExists:
			exists := store.Exists(exec.Args[0])
			result := 0
			if exists {
				result = 1
			}
			finish := time.Since(start)
			replySignal, _ := socket.CommandToMessage(&socket.Command{
				Type:    socket.SignalSuccess,
				User:    command.User,
				Payload: fmt.Sprint(result),
				Headers: socket.CommandHeader{
					"TIME": finish.String(),
				},
			})
			conn.Write([]byte(fmt.Sprintf("%s\n", replySignal)))
		}
	}
}
