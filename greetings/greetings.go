package greetings

import (
	"fmt"

	"errors"

	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty_name")
	}
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

func randomGreetings() string {
	formats := []string{
		"Hi, %s How Are You",
		"Hey %s, WassUp !",
		"Greeting of the Day %s",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}
