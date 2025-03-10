package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Vérifie l'argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: main_optimise <fichier d'entrees>")
		return
	}

	filename := os.Args[1]
	f, _ := os.Open(filename)
	sc := bufio.NewScanner(f)
	entrees := make(map[int][]int) // entrees : {44908357:[21 8 206 35 7], ... }
	res := make(map[int]bool)      // res : {44908357: true, ...}

	// Lecture des entrées
	for sc.Scan() {
		line := sc.Text()
		p := strings.Split(line, ": ")
		testVal, _ := strconv.Atoi(p[0])
		entrees[testVal] = StrListToIntList(strings.Fields(p[1]))
		res[testVal] = false

	}

	// Boucle principale
	s := 0
	for testVal, lNb := range entrees {
		// testVal : 44908357, lNb : [21 8 206 35 7]
		if valideVal(testVal, lNb) {
			s += testVal
			res[testVal] = true
		}
	}
	print(s)
}

func StrListToIntList(l []string) []int {
	res := make([]int, len(l))
	for i, s := range l {
		res[i], _ = strconv.Atoi(s)
	}
	return res
}

// Vérifie si l'écriture de testVal se termine par nb et renvoie le reste
// endsWith(1234, 34) renvoie (true, 12)
func endsWith(testVal, nb int) (bool, int) {

	// Nombre de chiffres de nb
	nbChiffres := int(math.Log10(float64(nb))) + 1
	puissance := int(math.Pow(10, float64(nbChiffres)))

	// Si testVal se termine par nb
	if testVal%puissance == nb {
		return true, testVal / puissance
	}

	return false, testVal
}

// Utilise une approche descendante pour savoir si une "valeur de test" est correcte
func valideVal(testVal int, lNb []int) bool {

	l := len(lNb)

	// Cas terminaux: Toutes les opérations ont été effectuées ou la valeur 0 a été atteinte
	if testVal == 0 || l == 0 {
		return testVal == 0 && l == 0 // Si la valeur de test est bonne, il devrait rester 0
	}

	// On prend le dernier nombre de la liste
	nb := lNb[l-1]
	reste := lNb[:l-1]

	// On essaie toutes les opérations possibles :
	// 1. Concaténation
	if concat, valSuiv := endsWith(testVal, nb); concat {
		if valideVal(valSuiv, reste) {
			return true
		}
	}

	// 2. Multiplication
	if nb != 0 && testVal%nb == 0 {
		valSuiv := testVal / nb
		if valideVal(valSuiv, reste) {
			return true
		}
	}

	// 3. Addition
	valSuiv := testVal - nb
	if valSuiv >= 0 { // On s'arrête si la valeur est négative puisque le résultat cible est 0
		if valideVal(valSuiv, reste) {
			return true
		}
	}

	// Aucune combinaison trouvée
	return false
}
