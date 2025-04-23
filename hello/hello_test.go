package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Amit", "")
		want := "Hello, Amit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello world when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Amit", "Spanish")
		want := "Hola, Amit"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Amit", "French")
		want := "Bonjour, Amit"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Have %q Wanted %q", got, want)
	}
}
