package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	models "github.com/ivar-mahhonin/food-delivery-classifier/pkg/models"
	util "github.com/ivar-mahhonin/food-delivery-classifier/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/navossoc/bayesian"
)

const (
	MAX_GO_ROUTINES = 10
)

var classifier *bayesian.Classifier
var wg sync.WaitGroup

func main() {
	stopWordsDir := envVariable("STOP_WORDS_DIR")
	trainDataDir := envVariable("TRAIN_DATA_DIR")
	modelFileDir := envVariable("MODEL_FILE_DIR")

	if stopWordsDir == "" {
		log.Print("STOP_WORDS_DIR is empty")
	}

	if trainDataDir == "" {
		log.Print("TRAIN_DATA_DIR is empty")
	}

	if modelFileDir == "" {
		log.Print("MODEL_FILE_DIR is empty")
	}
	classifier = getBaseModel(modelFileDir, trainDataDir, stopWordsDir)
}

//Reads support cases data, reads the model from the file, and generates a new model if necessary.
//It then waits for all the support cases to be trained before exiting.
func getBaseModel(modelFileDir string, trainDataDir string, stopWordsDir string) *bayesian.Classifier {
	cases, stopWords, errorReadData := util.ReadTrainingData(trainDataDir, stopWordsDir)

	if errorReadData != nil {
		log.Panic(errorReadData)
		os.Exit(1)
	}

	var classes []bayesian.Class

	for k := range cases {
		class := bayesian.Class(k)
		classes = append(classes, class)
	}

	classifier, errorReadingModel := util.ReadModelFromFile(modelFileDir)

	if errorReadingModel != nil {
		log.Print("Can not read model from file")
	}

	if classifier == nil || classifier.Learned() < len(classes) {
		if classifier != nil {
			log.Print("There are more new classes in test data")
		}

		log.Print("Generating new model")
		classifier = createClassifierFromTestData(classes, cases, stopWords)
		modelWriteErr := util.WriteModelToFile(modelFileDir, classifier)
		if modelWriteErr != nil {
			log.Panic(modelWriteErr)
			os.Exit(1)
		}
	} else {
		log.Printf("Found existing model with [%d classes] learned and [%d words] learned for every class", classifier.Learned(), classifier.WordCount())
	}
	wg.Wait()

	return classifier
}

func classify(text string, classes []models.FileTestData, stopWords []string, classifier *bayesian.Classifier) string {
	testTexts := make([]string, 0)
	testTexts = append(testTexts, text)

	tokenized := util.Tokenize(testTexts, stopWords)

	score, likely, strict := classifier.LogScores(
		tokenized,
	)

	println(likely)
	fmt.Printf("%v \n", score)
	fmt.Println("strict", strict)

	return classes[likely].Class
}

// Creates a classifier from support cases, given the classes and stop words.
func createClassifierFromTestData(classes []bayesian.Class, cases map[string][]string, stopWords []string) *bayesian.Classifier {
	classifier := parallelClassifierTraining(cases, classes, stopWords)
	classifier.ConvertTermsFreqToTfIdf()
	return classifier
}

// Trains a classifier in parallel using multiple goroutines.
func parallelClassifierTraining(cases map[string][]string, classes []bayesian.Class, stopWords []string) *bayesian.Classifier {
	classifier := bayesian.NewClassifier(classes...)
	log.Printf("Found %d classes", len(classes))

	tasksChannel := make(chan models.Pair[string, []string])

	for i := 0; i < MAX_GO_ROUTINES; i++ {
		wg.Add(1)
		go func() {
			for sc := range tasksChannel {
				tokens := util.Tokenize(sc.Second, stopWords)
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

func envVariable(key string) string {
	dir, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/../../.env", dir))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
