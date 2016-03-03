package articleStore_test

import (
	"encoding/json"
	"github.com/opinionated/articleStore"
	"os"
	"testing"
)

// TODO: test fail cases (ie folder/ file doesn't exist, bad format etc)

type testStruct struct {
	Name string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func expectFail(err error) {
	if err == nil {
		panic("expected error!")
	}
}

func cleanup(where string) error {
	return os.Remove(where)
}

func TestCreateFolder(t *testing.T) {
	s := articleStore.BuildStore("./tests/", "txt")

	_, err := s.CreateFolder("test1")
	handleError(err)
}

func TestStoreData(t *testing.T) {
	buildJsonWithError := func(name string) []byte {
		resp, err := json.Marshal(testStruct{Name: name})
		handleError(err)
		return resp
	}

	// build the structs
	names := []string{"one", "two"}
	structs := make([][]byte, 2)
	for i, str := range names {
		structs[i] = buildJsonWithError(str)
	}

	store := articleStore.BuildStore("tests", "json")

	// create stores and fill them
	for i, data := range structs {
		_, err := store.CreateFolder(names[i])
		handleError(err)
		err = store.StoreData(data, "base", names[i])
		handleError(err)
	}

	storedFiles, err := store.ListArticles()
	handleError(err)

	for i, name := range storedFiles {
		if name != names[i] {
			t.Errorf("bad file name: %s\n", name)
		}

		data, err := store.GetData("base", names[i])
		handleError(err)

		tmp := testStruct{}
		err = json.Unmarshal(data, &tmp)
		handleError(err)
		if tmp.Name != names[i] {
			t.Errorf("bad file data! name: %s data:%s\n", tmp.Name, data)
		}
	}
}
