package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/TahaRha/optiforge/algorithms" // Replace 'yourmodule' with your module name/path
)

func main() {
	http.HandleFunc("/solve", solveLPHandler)
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// SolveLPHandler is a handler for solving linear programming problems.
func solveLPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var problem algorithms.LPProblem
	if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	solution := algorithms.SolveLPProblem(problem)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(solution); err != nil {
		log.Printf("Could not encode response: %v", err)
	}
}
