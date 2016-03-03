package articleStore

import (
	"io/ioutil"
	"os"
)

// Store holds articles and related article data
type Store struct {
	// where the head of the filesystem is
	basePath string
	// the kind of file used to store data
	fileType string
}

// TODO: sanitize the path to make sure we only look at subdirs

// BuildStore creates a new store with provided data.
func BuildStore(basePath, fileType string) Store {
	// TODO: autodetect the path
	return Store{basePath, fileType}
}

// CreateFolder creates a new folder to hold an article and it's data.
// This won't overwrite any existing data.
func (s Store) CreateFolder(folderName string) (string, error) {
	prefix, err := s.getPrefix(folderName)
	if err != nil {
		return folderName, err
	}
	err = os.MkdirAll(prefix, 0777)
	return folderName, err
}

// StoreData stores an article's data in a file specified by fileName in an
// article's folder.
func (s Store) StoreData(data []byte, fileName, articleName string) error {
	// TODO: tighten this up
	prefix, err := s.getPrefix(articleName)
	if err != nil {
		return err
	}

	file, err := os.Create(prefix + fileName + "." + s.fileType)

	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return file.Close()
}

// GetData from the article store.
func (s Store) GetData(typeName, articleName string) ([]byte, error) {
	prefix, err := s.getPrefix(articleName)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(prefix + typeName + "." + s.fileType)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

// ListArticles all the articles in the base directory.
func (s Store) ListArticles() ([]string, error) {
	dir, err := os.Open(s.basePath)
	if err != nil {
		return nil, err
	}

	defer dir.Close()
	return dir.Readdirnames(0)
}

// TODO: sanitation here
func (s Store) getPrefix(articleName string) (string, error) {
	return s.basePath + "/" + articleName + "/", nil
}
