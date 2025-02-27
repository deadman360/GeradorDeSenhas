package main

import (
	"fmt"
	"math/rand"
)

func main() {
	psswd := generatePassword(800, true, true, true, true)
	fmt.Println(psswd)
}

func generatePassword(tamanho int, lower bool, upper bool, number bool, spec bool) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	var options []string
	if lower {
		options = append(options, lowerCase)
	}
	if upper {
		options = append(options, upperCase)
	}
	if number {
		options = append(options, numbers)
	}
	if spec {
		options = append(options, special)
	}
	if options == nil {
		return "Error"
	}
	var senha string

	for len(senha) < tamanho {
		for _, value := range options {
			if len(senha) == tamanho {
				break
			}
			senha = senha + string(value[rand.Intn((len(value)-1))])
		}
	}

	return senha
}
