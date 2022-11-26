package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people when a name is supplied", func(t *testing.T) {
		want := "Hello, Charles"
		got := Hello("Charles", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when no string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Fpanish", func(t *testing.T) {
		got := Hello("Charles", "Spanish")
		want := "Hola, Charles"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Charles", "French")
		want := "Bonjour, Charles"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
