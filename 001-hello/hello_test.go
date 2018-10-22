package main

import "testing"

/*
Notes:
- t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
*/

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tester", "")
		want := "Hello, Tester"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jacques", "French")
		want := "Bonjour, Jacques"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloTableDriver(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	var tests = []struct {
		name string
		lang string
		want string
		desc string
	}{
		{"Tester", "", "Hello, Tester", "saying hello to people"},
		{"", "", "Hello, World", "say 'Hello, World' when an empty string is supplied"},
		{"Elodie", "Spanish", "Hola, Elodie", "in Spanish"},
		{"Jacques", "French", "Bonjour, Jacques", "in Spanish"},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := Hello(test.name, test.lang)
			assertCorrectMessage(t, got, test.want)
		})
	}
}
