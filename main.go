package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Função que gera a sequência de Fibonacci
func fibonacci(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			seq[i] = 0
		} else if i == 1 {
			seq[i] = 1
		} else {
			seq[i] = seq[i-1] + seq[i-2]
		}
	}
	return seq
}

// Handler HTTP
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, "Missing 'number' parameter", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		http.Error(w, "'number' must be a positive integer", http.StatusBadRequest)
		return
	}

	result := fibonacci(number)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"sequence": result,
	})
}

func main() {
	http.HandleFunc("/fibonacci", fibonacciHandler)

	fmt.Println("🚀 Servidor a correr na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
