package util

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/navossoc/bayesian"
)

func TestReadFile(t *testing.T) {
	expected := []string{"a", "b", "c"}
	err := ioutil.WriteFile("test.json", []byte(`["a", "b", "c"]`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	result, err := readFile[string]("test.json")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestReadFileInvalidJSON(t *testing.T) {
	err := ioutil.WriteFile("test.json", []byte(`["a", "b", "c"`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	_, err = readFile[string]("test.json")
	if err == nil {
		t.Errorf("Expected error reading invalid JSON data")
	}
}

func TestReadFileNotExist(t *testing.T) {
	_, err := readFile[string]("test.json")
	if err == nil {
		t.Errorf("Expected error reading non-existent file")
	}
}

func TestReadFileEmpty(t *testing.T) {
	err := ioutil.WriteFile("test.json", []byte{}, 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	_, err = readFile[string]("test.json")
	if err == nil {
		t.Errorf("Expected error reading empty file, which is not valid json")
	}
}

func TestReadStopWords(t *testing.T) {
	expected := map[string]struct{}{"a": {}, "b": {}, "c": {}}
	err := ioutil.WriteFile("test.json", []byte(`["a", "b", "c"]`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	result, err := readStopWords("test.json")
	if err != nil {
		t.Errorf("Error reading stop words file: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestReadTestData(t *testing.T) {
	expected := map[string][]string{"class1": {"title1 description1", "title2 description2"}, "class2": {"title3 description3", "title4 description4"}}
	err := ioutil.WriteFile("test.json", []byte(`[{"_source": {"issue": "title1", "complaint_what_happened": "description1", "product": "class1"}}, {"_source": {"issue": "title2", "complaint_what_happened": "description2", "product": "class1"}}, {"_source": {"issue": "title3", "complaint_what_happened": "description3", "product": "class2"}}, {"_source": {"issue": "title4", "complaint_what_happened": "description4", "product": "class2"}}]`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	result, err := readTestData("test.json")
	if err != nil {
		t.Errorf("Error reading test data file: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case failed: got %v, want %v", result, expected)
	}
}

func TestReadTestDataInvalidJSON(t *testing.T) {
	err := ioutil.WriteFile("test.json", []byte(`[{"_source": {"issue": "titl`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	_, err = readTestData("test.json")
	if err == nil {
		t.Errorf("Expected error reading invalid JSON data")
	}
}

func TestReadTestDataNotExist(t *testing.T) {
	_, err := readTestData("test.json")
	if err == nil {
		t.Errorf("Expected error reading non-existent file")
	}
}

func TestReadTestDataNoClass(t *testing.T) {
	err := ioutil.WriteFile("test.json", []byte(`[{"_source": {"issue": "title1", "complaint_what_happened": "description1", "product": "class1"}}, {"_source": {"issue": "title2", "complaint_what_happened": "description2", "product": "class1"}}, {"_source": {"issue": "title3", "complaint_what_happened": "description3"}}, {"_source": {"issue": "title4", "complaint_what_happened": "description4"}}]`), 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test.json")
	result, err := readTestData("test.json")
	if err != nil {
		t.Errorf("Error reading test data file: %v", err)
	}
	if len(result) > 1 {
		t.Errorf("Test case failed: got %v, want only class1 data", result)
	}
	if len(result["class1"]) != 2 {
		t.Errorf("Test case failed: got %v, want 2 items in class1 data", result["class1"])
	}
}

var classes = []bayesian.Class{bayesian.Class("class1"), bayesian.Class("class2")}
var classifier = bayesian.NewClassifier(classes...)

func TestWriteModelToFile(t *testing.T) {
	err := WriteModelToFile("test_dir/test_model.gob", classifier)
	if err != nil {
		t.Errorf("Error writing model to file: %v", err)
	}
	defer os.RemoveAll("test_dir")
	if _, err := os.Stat("test_dir/test_model.gob"); os.IsNotExist(err) {
		t.Errorf("Model file not created")
	}
}

func TestWriteModelToFileExist(t *testing.T) {
	err := ioutil.WriteFile("test_model.gob", []byte{}, 0666)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer os.Remove("test_model.gob")
	err = WriteModelToFile("test_model.gob", classifier)
	if err != nil {
		t.Errorf("Error writing model to file: %v", err)
	}
}
func TestWriteModelToFileInvalidName(t *testing.T) {
	err := WriteModelToFile("/test_model", classifier)
	defer os.RemoveAll("/test_model")
	if err == nil {
		t.Errorf("Expected error writing model to file with invalid name")
	}
}

func TestWriteModelToFileToNonExistingdDirectory(t *testing.T) {
	err := WriteModelToFile("non-existing/path/test_model.gob", classifier)
	defer os.RemoveAll("non-existing")
	if err != nil {
		t.Errorf("Expected to wrtie model file to non existing directory")
	}
}
