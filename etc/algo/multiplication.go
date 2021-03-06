package algo

import "fmt"

// a la russe
func Multiplication() {
	var nA, nB, result int
	fmt.Scan(&nA, &nB)

	for nA >= 1 {
		if nA&1 == 1 {
			result += nB
		}

		nA >>= 1 // nA /= 2
		nB <<= 1 // nB *= 2
	}

	fmt.Println(result)
}
