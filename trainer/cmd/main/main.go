package main

import (
	"log"
	"os"

	util "github.com/ivar-mahhonin/food-delivery-classifier/trainer/pkg/utils"
	"github.com/navossoc/bayesian"
)

var classifier *bayesian.Classifier

func main() {
	stopWordsDir := util.GetEnvVariable("STOP_WORDS_DIR")
	trainDataDir := util.GetEnvVariable("TRAIN_DATA_DIR")
	modelFileDir := util.GetEnvVariable("MODEL_FILE_DIR")

	if stopWordsDir == "" || trainDataDir == "" || modelFileDir == "" {
		if stopWordsDir == "" {
			log.Print("STOP_WORDS_DIR is empty")
		}
		if trainDataDir == "" {
			log.Print("TRAIN_DATA_DIR is empty")
		}
		if modelFileDir == "" {
			log.Print("MODEL_FILE_DIR is empty")
		}
		os.Exit(1)
	}

	classifier = util.GetBaseModel(modelFileDir, trainDataDir, stopWordsDir)
}
