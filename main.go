package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isThisAValidChar(needle rune) bool {
	wanted := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	low := 0
	high := len(wanted) - 1

	for low <= high {
		median := (low + high) / 2

		if wanted[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(wanted) || wanted[low] != needle {
		return false
	}

	return true
}

func parseToInt(num rune) int {
	wanted := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	low := 0
	high := len(wanted) - 1

	for low <= high {
		median := (low + high) / 2

		if wanted[median] < num {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(wanted) || wanted[low] != num {
		return 0
	}

	return int(wanted[low]) - 48
}

func removeUnwantedCharacters(in string) string {
	out := ""

	for _, char := range in {
		if isThisAValidChar(char) {
			out += string(char)
		}
	}

	return out
}

func isThisADummyCPF(cpf string) bool {
	dummies := []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
		"01234567890",
	}

	for _, d := range dummies {
		if d == cpf {
			return true
		}
	}

	return false
}

func getVerificationDigits(cpf []int) [2]int {
	digits := [2]int{0, 0}
	counts := [2]int{10, 11}
	var cp []int

	cp = append(cp, cpf...)

	if len(cp) < 11 {
		cp = append(cp, 0, 0)
	}

	for i, c := range counts {
		sum := 0
		count := c
		for _, n := range cp {
			sum += n * count
			count--

			if count <= 1 {
				break
			}
		}

		remainer := sum % 11
		if remainer <= 1 {
			digits[i] = 0
		} else {
			digits[i] = 11 - remainer
		}

		cp[9] = digits[0]
	}

	return digits
}

func Verify(cpf string) (bool, error) {
	processed := removeUnwantedCharacters(cpf)

	if len(processed) != 11 {
		return false, fmt.Errorf("invalid length! a valid cpf must be 11 characters long")
	}

	if isThisADummyCPF(processed) {
		return false, fmt.Errorf("invalid cpf! '%s' is not a valid cpf", processed)
	}

	numbers := []int{}
	for _, c := range processed {
		numbers = append(numbers, parseToInt(c))
	}

	digits := getVerificationDigits(numbers)

	if numbers[9] != digits[0] {
		return false, fmt.Errorf("invalid first verifier digit! expected '%d'; got '%d'", digits[0], numbers[9])
	}

	if numbers[10] != digits[1] {
		return false, fmt.Errorf("invalid second verifier digit! expected '%d'; got '%d'", digits[1], numbers[10])
	}

	return true, nil
}

func GenerateCPF() string {
	cpf := ""
	numbers := []int{}

	for i := 0; i < 9; i++ {
		rand.Seed(time.Now().UnixNano())
		numbers = append(numbers, rand.Intn(10))
		cpf += fmt.Sprint(numbers[i])
	}

	digits := getVerificationDigits(numbers)

	for _, d := range digits {
		cpf += fmt.Sprint(d)
	}

	return cpf
}

func main() {
}
