package main

import "fmt"

func simplify(n, d int) (numerator, denominator int) {
	numerator, denominator = n, d
	for i := 2; i <= n; i++ {
		for {
			if numerator == 1 || denominator == 1 {
				break
			}
			if numerator%i == 0 && denominator%i == 0 {
				numerator /= i
				denominator /= i
			} else {
				break
			}

		}
	}
	return numerator, denominator
}

func main() {
	numerators := []float64{}
	denominators := []float64{}

	addNum := func(u, y float64) {
		numerators = append(numerators, u)
		denominators = append(denominators, y)
	}

	for i := 1.0; i < 10.0; i++ {
		for u := 1.0; u < 10.0; u++ {
			for y := 1.0; y < 10.0; y++ {
				if (u*10+i)/(y*10+i) == u/y {
					if u/y < 1.0 {
						addNum(u, y)
					}
				}
				if (i*10+u)/(i*10+y) == u/y {
					if u/y < 1.0 {
						addNum(u, y)
					}
				}
				if (i*10+u)/(y*10+i) == u/y {
					if i/y < 1.0 {
						addNum(u, y)
					}
				}
				if (u*10+i)/(i*10+y) == u/y {
					if u/y < 1.0 {
						addNum(u, y)
					}
				}
			}
		}
	}
	num, den := 1, 1
	for i := range numerators {
		n, d := simplify(int(numerators[i]), int(denominators[i]))
		num *= n
		den *= d
	}
	_, den = simplify(num, den)
	fmt.Println(den)
}
