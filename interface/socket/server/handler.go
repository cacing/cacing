package server

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func Handler(config *Config, command string, args []value) (func() string, bool) {
	handler, ok := map[string]func() string{
		"COMMAND": func() string {
			return "+OK"
		},
		"PING": func() string {
			if len(args) <= 0 {
				return "+OK"
			}
			res := bytes.Buffer{}
			res.WriteByte('+')
			for _, a := range args {
				res.Write(a.Bulk)
				res.WriteByte(' ')
			}
			return res.String()
		},
		"SET": func() string {
			key := string(args[0].Bulk)
			value := args[1].Bulk
			err := store.Set(key, value, 0)
			if err != nil {
				return "-ERROR " + err.Error()
			}
			return "+OK"
		},
		"GET": func() string {
			key := string(args[0].Bulk)
			value, err := store.Get(key)
			if err != nil {
				return "-ERROR " + err.Error()
			}
			return fmt.Sprintf("+%s", value)
		},
		"DEL": func() string {
			keys := make([]string, len(args))
			for i, a := range args {
				keys[i] = string(a.Bulk)
			}

			_, err := store.Delete(keys...)
			if err != nil {
				return "-ERROR " + err.Error()
			}

			return "+OK"
		},
		"EXPIRE": func() string {
			key := string(args[0].Bulk)
			expire := string(args[1].Bulk)

			expireInt, err := strconv.Atoi(expire)
			if err != nil {
				return "-ERROR " + err.Error()
			}

			_, err = store.SetExpired(key, time.Duration(expireInt)*time.Second)
			if err != nil {
				return "-ERROR " + err.Error()
			}
			return "+OK"
		},
	}[string(command)]
	return handler, ok
}
