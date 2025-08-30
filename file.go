package go_file_manipulation

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CreateFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message + "\n")

	return nil
}

func ReadFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		message += string(line) + "\n"
	}

	var log string = "All line has been read"
	defer func() {
		fmt.Println(log)
	}()

	return message, nil
}

func AddMessage(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message + "\n")

	return nil
}
