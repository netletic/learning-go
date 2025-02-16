package main

import (
	"fmt"
	"unicode/utf8"
)

func rangeRune(word string) {
	fmt.Printf("range over string '%s' and print char:\n", word)
	fmt.Printf("length of word is %d \n", len(word))
	fmt.Printf("rune count in string is %d\n", utf8.RuneCountInString(word))
	for _, chr := range word {
		fmt.Printf("%c\n", chr)
		fmt.Printf("actual value %v\n", chr)
	}
}

func rangeByte(word string) {
	fmt.Printf("range over string '%s' and print byte:\n", word)
	fmt.Printf("length of word is %d \n", len(word))
	fmt.Printf("rune count in string is %d\n", utf8.RuneCountInString(word))
	for i := range word {
		fmt.Printf("char %c\n", word[i])
		fmt.Printf("actual value %v\n", word[i])
	}
}

func runeArray(word string) {
	fmt.Printf("range over string '%s' and print byte:\n", word)
	fmt.Printf("length of word is %d \n", len(word))
	fmt.Printf("rune count in string is %d\n", utf8.RuneCountInString(word))
	runeArray := []rune(word)
	for i := range runeArray {
		fmt.Printf("%c\n", runeArray[i])
		fmt.Printf("actual value %v\n", runeArray[i])
	}
}

func main() {
	// in Golang, a string is a slice of bytes
	// where each element in the string represents a decimal value
	// which corresponds to a hexadecimal value
	// calling len() on a string will give you the number of bytes in the string
	// accessing an element of a string by index will return a byte

	// ASCII only supported 7 bits, which was only enough space to represent
	// english/latin characters and punctuation

	// instead, unicode maps each character to a codepoint
	// a codepoint is a hexadecimal number that represents a character
	// e.g. U+0041 -> A

	// UTF-8 -- UTF-8 stands for Unicode Transformation Format
	// algorithmic mapping from every codepoint to a unique byte sequence
	// UTF-8 has variable length encoding
	// can represent characters
	// with small values (like A) with one bytes
	// or with large values, with upto four bytes
	// it's also backwards compatible with ASCII
	// UTF-8 is not the dominant encoding – represents 98% of the web

	// a rune in Golang is the same as a codepoint
	// a rune is a decimal value, which refers to a codepoint, and that codepoint
	// has a hexadecimal representation – it's the codepoint hexadecimal representation
	// that determines how many bytes are needed to represent the character
	rangeRune("a£c")
	rangeByte("a£c")
	runeArray("a£c")
}
