package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Fibonacci sequence starting at 1,1...
func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}

	seq := []int{1, 1}
	for i := 2; i < n; i++ {
		next := seq[i-1] + seq[i-2]
		seq = append(seq, next)
	}
	return seq
}

// HTTP Handler
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Parâmetro 'n' é obrigatório. Exemplo: /fibonacci?n=7", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		http.Error(w, "Parâmetro 'n' deve ser um número inteiro positivo", http.StatusBadRequest)
		return
	}

	seq := Fibonacci(n)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"sequence": seq,
	})
}

func main() {
	http.HandleFunc("/fibonacci", fibonacciHandler)

	fmt.Println("🚀 Servidor a correr na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
