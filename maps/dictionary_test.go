package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknow")
		if err == nil {
			t.Fatalf("expected to get an error.")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	defination := "this is just a test"

	t.Run("new word", func(t *testing.T) {
		dictionary.Add(word, defination)
		assertDefinition(t, dictionary, word, defination)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary.Add(word, defination)
		err := dictionary.Add(word, defination)
		if err == nil {
			t.Fatalf("expected to get an error.")
		}

		assertError(t, err, ErrAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	defination := "this is just a test"

	t.Run("existing word", func(t *testing.T) {
		dictionary.Add(word, defination)
		err := dictionary.Delete(word)
		assertError(t, err, nil)
	})
	t.Run("unknow word", func(t *testing.T) {
		err := dictionary.Delete(word)
		if err == nil {
			t.Fatalf("expected to get an error.")
		}
		assertError(t, err, ErrWordDoesNotExists)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, defination string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, defination)
}
