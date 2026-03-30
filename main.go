package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b float64
	var op string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple CLI Calculator")
	fmt.Println("Enter calculation (e.g. 5 + 3):")

	// Read input
	fmt.Fscan(reader, &a, &op, &b)

	// Perform calculation
	switch op {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Invalid operator. Use +, -, *, /")
	}
}

///PALINDROMW

package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindrom("no lemon no melon"))
}

func palindrom(word string) string {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}


ORRRRRRRRRRR TO MAKE IT PRINT TRUE OR FALSE

// palindrome

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}


///METER TO KM AND BACK

package main

import "fmt"

// km to meters
func kmToM(km float64) float64 {
	return km * 1000
}

// meters to km
func mToKm(m float64) float64 {
	return m / 1000
}

func main() {
	fmt.Println("5 km =", kmToM(5), "meters")
	fmt.Println("2000 m =", mToKm(2000), "km")
}



