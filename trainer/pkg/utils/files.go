package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	models "github.com/ivar-mahhonin/food-delivery-classifier/trainer/pkg/models"
	"github.com/navossoc/bayesian"
)

func ReadTrainingData(testDataDir string, stopWordsDir string) (map[string][]string, map[string]struct{}, error) {
	cases, errReadingTestData := readTestData(testDataDir)
	if errReadingTestData != nil {
		log.Fatal("ReadTrainingData: can not read test data: ", errReadingTestData)
		return nil, nil, errReadingTestData
	}

	stopWords, errReadingStopWords := readStopWords(stopWordsDir)
	if errReadingStopWords != nil {
		log.Fatal("ReadTrainingData: can not read stop words")
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

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	classifier.WriteToFile(filePath)
	return err
}

func readStopWords(stopWordsDir string) (map[string]struct{}, error) {
	stopWords, err := readFile[string](stopWordsDir)
	stopWordsMap := make(map[string]struct{})

	if err != nil {
		return nil, err
	}

	for _, s := range stopWords {
		stopWordsMap[s] = struct{}{}
	}

	return stopWordsMap, nil
}

func readFile[T any](fileName string) ([]T, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var data []T

	if !json.Valid([]byte(bytes)) {
		return nil, errors.New("not a valid json")
	}

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
