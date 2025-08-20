package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	if tcp, ok := conn.(*net.TCPConn); ok {
		_ = tcp.SetNoDelay(true)
	}

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Minute))
		reader := bufio.NewReader(conn)
		cmd := command{reader}
		value, err := cmd.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("invalid request: %s\n", err.Error())
			}
			continue
		}
		if value.Type != ARRAY {
			fmt.Println("invalid request, expected array")
			continue
		}
		if len(value.Array) <= 0 {
			fmt.Println("invalid request, expected array lenght > 0")
			continue
		}
		command := string(value.Array[0].Bulk)
		args := value.Array[1:]

		if command == "AUTH" {
			username := string(args[0].Bulk)
			password := string(args[1].Bulk)
			err := authenticateClient(config, username, password)
			if err != nil {
				conn.Write(fmt.Appendf([]byte{}, "-WRONGPASS %s\r\n", err.Error()))
				conn.Close()
				return
			}
			conn.Write(fmt.Appendf([]byte{}, "+OK\r\n"))
			continue
		}

		handler, ok := Handler(config, command, args)
		if !ok {
			fmt.Printf("invalid request, unexpected command: %s\n", command)
			conn.Write(nil)
		} else {
			conn.Write(fmt.Appendf([]byte{}, "%s\r\n", handler()))
		}

	}

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
