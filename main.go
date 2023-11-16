package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL do webhook
	webhookURL := "https://webhook.site/3ae3e0a2-ff88-42e0-a713-b5f2a70b799d"

	// Dados para enviar
	data := map[string]interface{}{
		"nome":   "Jenny",
		"idade":  25,
		"cidade": "São josé dos Campos",
	}

	// Converte os dados para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erro ao converter dados para JSON:", err)
		return
	}

	// Envia a solicitação POST para o webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação:", err)
		return
	}
	defer resp.Body.Close()

	// Exibe a resposta do webhook
	fmt.Println("Resposta do webhook:")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
