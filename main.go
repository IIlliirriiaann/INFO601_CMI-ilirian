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
		fmt.Println("Usage: main <fichier d'entrees>")
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
	for testVal, lNb := range entrees {
		// testVal : 44908357, lNb : [21 8 206 35 7]
		l := len(lNb)
		iMax := int(math.Pow(3, float64(l-1)))

		// Pour chaque nombre jusqu'à iMax, on applique le masque en base 3 unique correspondant
		for i := range iMax {
			mask := toBase3(i, l) // mask : 0121
			if testVal == applyMask(lNb, mask) {
				res[testVal] = true
				break
			}
		}
	}

	// Lecture des résultats
	s := 0
	for testVal, valide := range res {
		if valide {
			s += testVal
		}
	}
	println(s)
}

func StrListToIntList(l []string) []int {
	res := make([]int, len(l))
	for i, s := range l {
		res[i], _ = strconv.Atoi(s)
	}
	return res
}

func applyMask(lNb []int, mask string) int {
	l := len(lNb)
	res := lNb[0]

	// Appliquer chaque nombre du masque
	for i := 1; i < l; i++ {
		n := lNb[i]
		// 0 -> + | 1 -> * | 2 -> ||
		switch mask[i] {
		case '0':
			res += n
		case '1':
			res *= n
		case '2':
			res = concat(res, n)
		}

	}
	return res
}

func concat(a, b int) int {
	pow := 10
	tmp := b
	for tmp >= 10 {
		pow *= 10
		tmp /= 10
	}
	return a*pow + b
}

func toBase3(n, l int) string {
	mask := make([]byte, l)
	for i := l - 1; i >= 0; i-- {
		mask[i] = byte('0' + (n % 3))
		n /= 3
	}
	return string(mask)
}
