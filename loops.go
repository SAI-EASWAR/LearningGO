package main

import (
	"fmt"
	"strings"
)

func Loops() {

	//for loops

	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	// using while Loops

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//sum of number from 1 -> 100

	var sum = 0

	for i := 0; i < 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum of numbers from 1 -> 100", sum)

	//print all the even numbers from 1 to 50

	var even []int
	var odd []int
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}

	}
	//Fibonacci Sequence
	//F(n)=F(n−1)+F(n−2)
	a, b := 0, 1
	var fib []int
	for i := 0; i < 10; i++ {
		fib = append(fib, a)

		a, b = b, a+b
	}
	fmt.Println("fibonacci series ", fib)

	palindromeStr := "A man a plan a canal Panama"
	if isPalindrome(palindromeStr) {
		fmt.Println("\"", palindromeStr, "\" is a palindrome string.")
	} else {
		fmt.Println("\"", palindromeStr, "\" is not a palindrome string.")
	}

}

func isPalindrome(str string) bool {
	// Convert the string to lowercase and remove spaces
	sanitizedStr := strings.ReplaceAll(strings.ToLower(str), " ", "")

	return PalindromeHelper(sanitizedStr)
}
func PalindromeHelper(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
