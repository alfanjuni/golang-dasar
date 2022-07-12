package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("alfan")
	if result != "Hello alfan" {
		t.Error("harusnya Hi alfan")
	}

	fmt.Println("Test hello world done")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("alfan")
	if result != "Hello2 alfan" {
		t.Fatal("harusnya 2 Hi alfan")
	}
	fmt.Println("Test hello world 2 done")
}
