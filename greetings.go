package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello. Devuelve un saludo

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New(" Nombre vacío")

	}
	//message := fmt.Sprintf("Hola, %v ¡Bienvenido!", name)
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

// Devuelve uno o varios mensajes aleatorios
func randomFormat() string {
	formats := []string{
		"Hola, %v ¡Bienvenido!",
		"Que bueno verte, %v",
		"Que onda, %v?",
		"Todo bien? %v",
	}
	return formats[rand.Intn(len(formats))]
}
