package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Erro eu buscar URL")
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Erro ao gerar UUID")
	}

	fmt.Printf("Hello: %s, %s\n", resp.Status, uuid)
}
