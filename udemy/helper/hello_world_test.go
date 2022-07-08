package helper

import "testing"

func TestHelloWordl(t *testing.T) {
	result := HelloWorld("alfan")
	if result != "Hello alfan" {
		panic("Result is not Hello alfan")
	}
}
