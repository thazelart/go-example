package hello_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/go-example/internal/hello"
	"testing"
)

// HelloWorld is a public function, we test it in a <package_name>_test.go file
func TestHelloWorld(t *testing.T) {
	expectedResult := "Hello world !"

	testResult := hello.HelloWorld()

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("HelloWorld() mismatch (-want +got):\n%s", diff)
	}
}

// HelloWorld is a public function, we test it in a <package_name>_test.go file
func TestNewPerson(t *testing.T) {
	expectedResult := hello.Person{Name: "test"}

	testResult := hello.NewPerson("test")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("HelloWorld() mismatch (-want +got):\n%s", diff)
	}
}

func TestSayHello(t *testing.T) {
	person := hello.Person{Name: "test"}
	expectedResult := "Hello test !"

	testResult := person.SayHello()

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("HelloWorld() mismatch (-want +got):\n%s", diff)
	}
}
