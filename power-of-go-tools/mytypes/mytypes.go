package mytypes

import "strings"

type MyInt int

// Twice multiplies its receiver by 2 and returns the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// MyString is a custom version of the `string` type.
type MyString string

// Len returns the length of the string.
func (s MyString) Len() int {
	return len(s)
}

type MyBuilder struct {
	Contents strings.Builder
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (i *MyInt) Double() {
	*i *= 2
}
