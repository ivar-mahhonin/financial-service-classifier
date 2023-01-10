package util

import (
	"io/ioutil"
	"os"
	"testing"

	models "github.com/ivar-mahhonin/food-delivery-classifier/trainer-service/pkg/models"
	"github.com/navossoc/bayesian"
)

func TestGetExistingBaseModel(t *testing.T) {
	cases := map[string][]string{
		"class1": {"This is a test sentence for class 1", "This is another test sentence for class 1"},
		"class2": {"This is a test sentence for class 2", "This is another test sentence for class 2"},
	}
	stopWords := map[string]struct{}{"is": {}}
	classifier := CreateClassifierFromTestData([]bayesian.Class{"class1", "class2"}, cases, stopWords)
	err := WriteModelToFile("test_dir/test_model.gob", classifier)
	if err != nil {
		t.Errorf("Error writing model to file: %v", err)
	}
	defer os.RemoveAll("test_dir")
	baseModel, err := GetBaseModel("test_dir/test_model.gob", "", "")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if baseModel.Learned() != 2 {
		t.Errorf("Incorrect number of classes learned")
	}
}

func TestTrainingNewBaseModel(t *testing.T) {
	err := ioutil.WriteFile("test_data.json", []byte(`[{"_source": {"issue": "title1", "complaint_what_happened": "description1", "product": "class1"}}, {"_source": {"issue": "title2", "complaint_what_happened": "description2", "product": "class1"}}, {"_source": {"issue": "title3", "complaint_what_happened": "description3", "product": "class2"}}, {"_source": {"issue": "title4", "complaint_what_happened": "description4", "product": "class2"}}]`), 0666)
	if err != nil {
		t.Errorf("Error creating test data file: %v", err)
	}
	defer os.Remove("test_data.json")

	err = ioutil.WriteFile("stop_words.json", []byte(`["a", "b", "c"]`), 0666)
	if err != nil {
		t.Errorf("Error creating stop words file: %v", err)
	}
	defer os.Remove("stop_words.json")

	baseModel, err := GetBaseModel("test_dir/test_model.gob", "test_data.json", "stop_words.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if baseModel.Learned() != 2 {
		t.Errorf("Incorrect number of classes learned")
	}
	defer os.RemoveAll("test_dir")
}

func TestCreateClassifierFromTestData(t *testing.T) {
	// Prepare test data
	classes := []bayesian.Class{"class1", "class2"}
	cases := map[string][]string{
		"class1": {"This is a test case for class 1", "This is another test case for class 1"},
		"class2": {"This is a test case for class 2", "This is another test case for class 2"},
	}
	stopWords := map[string]struct{}{"for": {}}

	// Call the function being tested
	classifier := CreateClassifierFromTestData(classes, cases, stopWords)

	// Check if the returned classifier is not nil
	if classifier == nil {
		t.Error("Expected classifier, got nil")
	}

	// Check if the number of classes learned is equal to the expected number of classes
	if classifier.Learned() != len(classes) {
		t.Errorf("Expected %d classes, got %d", len(classes), classifier.Learned())
	}

}

func TestClassify(t *testing.T) {
	classes := []models.FileTestData{
		{Class: "class1"},
		{Class: "class2"},
		{Class: "class3"},
	}
	stopWords := map[string]struct{}{
		"the": {},
		"is":  {},
	}
	classifier := bayesian.NewClassifier(bayesian.Class("class1"), bayesian.Class("class2"), bayesian.Class("class3"))
	classifier.Learn([]string{"this", "is", "a", "text"}, bayesian.Class("class1"))
	classifier.Learn([]string{"this", "is", "another", "text"}, bayesian.Class("class2"))
	classifier.Learn([]string{"yet", "another", "text"}, bayesian.Class("class3"))

	t.Run("Correct classification for class1", func(t *testing.T) {
		text := "this is a text"
		result, score := classify(text, classes, stopWords, classifier)
		expected := "class1"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}

		if score < 90 {
			t.Errorf("Expected %f score, to be more than 90", score)
		}
	})

	t.Run("Correct classification for class2", func(t *testing.T) {
		text := "this is another text"
		result, score := classify(text, classes, stopWords, classifier)
		expected := "class2"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
		if score < 90 {
			t.Errorf("Expected %f score, to be more than 90", score)
		}
	})
}
