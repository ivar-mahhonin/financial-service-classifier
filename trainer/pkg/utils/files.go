package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	models "github.com/ivar-mahhonin/food-delivery-classifier/pkg/models"
	"github.com/navossoc/bayesian"
)

func ReadTrainingData(testDataDir string, stopWordsDir string) (map[string][]string, []string, error) {
	cases, errReadingTestData := readTestData(testDataDir)
	if errReadingTestData != nil {
		log.Fatal("Can not read test data: ", errReadingTestData)
		return nil, nil, errReadingTestData
	}

	stopWords, errReadingStopWords := readStopWords(stopWordsDir)
	if errReadingStopWords != nil {
		log.Fatal("Can not read stop words")
		return nil, nil, errReadingStopWords
	}

	return cases, stopWords, nil
}

func ReadModelFromFile(modelFileDir string) (*bayesian.Classifier, error) {
	classifier, err := bayesian.NewClassifierFromFile(modelFileDir)
	if err != nil {
		return nil, err
	}
	return classifier, err
}

func WriteModelToFile(modelFileDir string, classifier *bayesian.Classifier) error {
	filePath := modelFileDir
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	classifier.WriteToFile(filePath)
	return err
}

func readStopWords(stopWordsDir string) ([]string, error) {
	return readFile[string](stopWordsDir)
}

func readFile[T any](fileName string) ([]T, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var data []T

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func readTestData(trainDataDir string) (map[string][]string, error) {
	sources, errReadingTestData := readFile[models.FileTestDataSource](trainDataDir)

	if errReadingTestData != nil {
		return nil, errReadingTestData
	}

	cases := make([]models.FileTestData, len(sources))

	for _, source := range sources {
		cases = append(cases, source.Source)
	}

	data := make(map[string][]string)

	for _, c := range cases {
		if (len(c.Title) > 0 || len(c.Description) > 0) && len(c.Class) > 0 {
			if data[c.Class] == nil {
				data[c.Class] = make([]string, 0)
			}
			data[c.Class] = append(data[c.Class], fmt.Sprintf("%s %s", c.Title, c.Description))
		}
	}
	return data, errReadingTestData
}