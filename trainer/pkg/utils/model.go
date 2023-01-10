package util

import (
	"log"
	"sync"

	models "github.com/ivar-mahhonin/food-delivery-classifier/trainer/pkg/models"
	"github.com/navossoc/bayesian"
)

var wg sync.WaitGroup

const (
	MAX_GO_ROUTINES = 10
)

//Reads support cases data, reads the model from the file, and generates a new model if necessary.
//It then waits for all the support cases to be trained before exiting.
func GetBaseModel(modelFileDir string, trainDataDir string, stopWordsDir string) (*bayesian.Classifier, error) {
	classifier, errorReadingModel := ReadModelFromFile(modelFileDir)

	if errorReadingModel != nil {
		log.Print("Can not read model from file: ", modelFileDir)
	}

	if classifier == nil {
		if classifier != nil {
			log.Print("There are more new classes in test data")
		}

		cases, stopWords, errorReadData := ReadTrainingData(trainDataDir, stopWordsDir)

		if errorReadData != nil {
			log.Panic(errorReadData)
			return nil, errorReadData
		}

		var classes []bayesian.Class

		for k := range cases {
			class := bayesian.Class(k)
			classes = append(classes, class)
		}

		log.Print("Generating new model")
		classifier = CreateClassifierFromTestData(classes, cases, stopWords)
		modelWriteErr := WriteModelToFile(modelFileDir, classifier)
		if modelWriteErr != nil {
			log.Panic(modelWriteErr)
			return nil, modelWriteErr
		}
	} else {
		log.Printf("Found existing model with [%d classes] learned and [%d words] learned for every class", classifier.Learned(), classifier.WordCount())
	}
	wg.Wait()

	return classifier, nil
}

// Creates a classifier from support cases, given the classes and stop words.
func CreateClassifierFromTestData(classes []bayesian.Class, cases map[string][]string, stopWords map[string]struct{}) *bayesian.Classifier {
	classifier := ParallelClassifierTraining(cases, classes, stopWords)
	classifier.ConvertTermsFreqToTfIdf()
	return classifier
}

// Trains a classifier in parallel using multiple goroutines.

func ParallelClassifierTraining(cases map[string][]string, classes []bayesian.Class, stopWords map[string]struct{}) *bayesian.Classifier {
	classifier := bayesian.NewClassifier(classes...)
	log.Printf("Found %d classes", len(classes))

	tasksChannel := make(chan models.Pair[string, []string])

	for i := 0; i < MAX_GO_ROUTINES; i++ {
		wg.Add(1)
		go func() {
			for sc := range tasksChannel {
				tokens := Tokenize(sc.Second, stopWords)
				classifier.Learn(tokens, bayesian.Class(sc.First))
				log.Printf("Trained '%s' class with %d tickets", sc.First, len(sc.Second))
			}
			wg.Done()
		}()
	}

	for k, v := range cases {
		tasksChannel <- models.Pair[string, []string]{First: k, Second: v}
	}

	close(tasksChannel)

	wg.Wait()
	return classifier
}

func classify(text string, classes []models.FileTestData, stopWords map[string]struct{}, classifier *bayesian.Classifier) string {
	testTexts := make([]string, 0)
	testTexts = append(testTexts, text)

	tokenized := Tokenize(testTexts, stopWords)

	_, likely, _ := classifier.LogScores(
		tokenized,
	)

	return classes[likely].Class
}
