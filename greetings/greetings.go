package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error message.
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a value that embeds the name
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
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

func randomFormat() string {
    formats := []string{
         "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
        "How do you do, %v!",
        "Yo, %v! What's up?",
    }

    fmt.Println(rand.Intn(len(formats)))

    return formats[rand.Intn(len(formats))]
}