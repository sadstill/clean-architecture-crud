package converter

import (
	"rest-api-crud/internal/model"
	"rest-api-crud/internal/storage"
)

func ToStorageAuthor(author model.Author) storage.Author {
	return storage.Author{
		ID:   author.ID,
		Name: author.Name,
	}
}

func ToModelAuthor(storageAuthor storage.Author) model.Author {
	return model.Author{
		ID:   storageAuthor.ID,
		Name: storageAuthor.Name,
	}
}
