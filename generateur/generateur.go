package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Toutes les règles (nombre de chiffres, intervalle des nombres, ...) sont purement arbitraires et peuvent être changées

func main() {
	// Vérifie si un argument est passé
	if len(os.Args) != 2 {
		fmt.Println("Usage: generateur <nombre_d_entrees>")
		return
	}

	// Convertit l'argument en entier
	var count int
	_, err := fmt.Sscanf(os.Args[1], "%d", &count)
	if err != nil {
		fmt.Println("L'argument doit être un entier.")
		return
	}

	// Initialise le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Ouvre le fichier pour écriture
	file, err := os.Create("input")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	defer file.Close()

	// Génère les équations et les écrit dans le fichier
	for i := 0; i < count; i++ {
		// Génère une équation valide
		equation := generateValidEquation()
		// Écrit l'équation dans le fichier
		_, err := fmt.Fprintln(file, equation)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
			return
		}
	}

	fmt.Printf("%d équations générées et sauvegardées\n", count)
}

// Génère une équation valide avec des opérations aléatoires
func generateValidEquation() string {
	for {
		// Génère entre 2 et 5 nombres aléatoires (pour garder les équations simples)
		numCount := rand.Intn(4) + 12 // Entre 2 et 5 nombres
		numbers := make([]int, numCount)
		for i := range numbers {
			numbers[i] = rand.Intn(1000) + 1 // Nombres entre 1 et 1000
		}

		// Génère des opérations aléatoires
		operations := make([]string, numCount-1)
		for i := range operations {
			operations[i] = randomOperation()
		}

		// Calcule le résultat en évaluant les opérations de gauche à droite
		result := numbers[0]
		for i := 0; i < len(operations); i++ {
			nextNum := numbers[i+1]
			switch operations[i] {
			case "+":
				result += nextNum
			case "*":
				result *= nextNum
			case "||":
				// Concaténation : convertit les nombres en chaînes, les combine, puis reconvertit en entier
				resultStr := strconv.Itoa(result) + strconv.Itoa(nextNum)
				result, _ = strconv.Atoi(resultStr)
			}
		}

		// Vérifie si le résultat a entre 4 et 15 chiffres
		if result >= 1000 && result < 1e15 {
			// Formate l'équation
			equation := fmt.Sprintf("%d:", result)
			for _, num := range numbers {
				equation += fmt.Sprintf(" %d", num)
			}
			return equation
		}
	}
}

// Retourne une opération aléatoire (+, *, ||)
func randomOperation() string {
	operations := []string{"+", "*", "||"}
	return operations[rand.Intn(len(operations))]
}
