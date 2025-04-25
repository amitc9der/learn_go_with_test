package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound          = DictionaryErr("could not find the word you are looking for")
	ErrAlreadyExists     = DictionaryErr("Word and Meaning already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot perfom operation on word because it does not exist")
)

func (d *Dictionary) Search(word string) (string, error) {
	defination, exists := (*d)[word]
	if !exists {
		return "", ErrNotFound
	}

	return defination, nil
}

func (d *Dictionary) Add(word, definition string) error {
	_, err := (*d).Search(word)

	switch err {
	case nil:
		return ErrAlreadyExists
	case ErrNotFound:
		(*d)[word] = definition
	default:
		return err
	}

	return nil
}

func (d *Dictionary) Delete(word string) error {
	_, err := (*d).Search(word)

	switch err {
	case nil:
		delete((*d), word)
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}
	return nil
}

func (d *Dictionary) Update(word, definition string) error {
	_, err := (*d).Search(word)
	switch err {
	case nil:
		(*d)[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}
	return nil
}
