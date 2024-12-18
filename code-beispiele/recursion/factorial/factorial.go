package factorial

// Factorial berechnet die Fakultät von n.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// FactorialIter berechnet die Fakultät von n iterativ.
func FactorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
