package main

import (
	"fmt"
	"unicode"
)


func isValidLuhn(cardNumber string) bool {
	var sum int
	var alternate bool

	
	for i := len(cardNumber) - 1; i >= 0; i-- {
		
		if !unicode.IsDigit(rune(cardNumber[i])) {
			continue
		}

		
		n := int(cardNumber[i] - '0')

		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
		alternate = !alternate
	}

	
	return sum%10 == 0
}

func main() {
	// Test cases
	testCases := []struct {
		number string
		valid  bool
	}{
		{"4539578763621486", true}, 
		{"4485732711235813", true}, 
		{"6011514433546201", true}, 
		{"6011514433546202", false}, 
		{"378282246310005", true},  
		{"371449635398431", true},  
		{"30569309025904", true},   
		{"38520000023237", true},    
	}

	for _, tc := range testCases {
		fmt.Printf("Testing card number: %s, valid: %v\n", tc.number, isValidLuhn(tc.number) == tc.valid)
	}
}
