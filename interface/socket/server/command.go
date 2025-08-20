package server

import (
	"bufio"
	"strconv"
)

type VType rune

const (
	ARRAY VType = 42
	BULK  VType = 36
)

type value struct {
	Type  VType
	Array []value
	Bulk  []byte
}

type command struct {
	reader *bufio.Reader
}

func (c command) readLine() ([]byte, int, error) {
	n := 0
	line := make([]byte, 0)
	for {
		b, err := c.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n++
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (c command) Read() (value, error) {
	vType, err := c.reader.ReadByte()
	if err != nil {
		return value{}, err
	}

	switch vType {
	case byte(ARRAY):
		return c.readArray()
	case byte(BULK):
		return c.readBulk()
	}

	return value{}, nil
}

func (c command) readInteger() (int, int, error) {
	line, n, err := c.readLine()
	if err != nil {
		return 0, 0, err
	}

	i, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(i), n, nil
}

func (c command) readArray() (value, error) {
	v := value{}
	v.Type = ARRAY
	vLen, _, err := c.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]value, 0)
	for range vLen {
		vv, err := c.Read()
		if err != nil {
			return v, err
		}
		v.Array = append(v.Array, vv)
	}
	return v, nil
}

func (c command) readBulk() (value, error) {
	v := value{}
	v.Type = BULK

	vLen, _, err := c.readInteger()
	if err != nil {
		return v, err
	}
	v.Bulk = make([]byte, vLen)
	c.reader.Read(v.Bulk)
	c.readLine()
	return v, nil
}
