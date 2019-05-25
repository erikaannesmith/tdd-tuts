package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world if empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eloise", "Spanish")
		want := "Hola, Eloise"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"

		assertCorrectMessage(t, got, want)
	})
}