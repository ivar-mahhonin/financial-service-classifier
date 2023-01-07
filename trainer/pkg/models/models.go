package util

type FileTestData struct {
	Class       string `json:"product"`
	Title       string `json:"issue"`
	Description string `json:"complaint_what_happened"`
}

type FileTestDataSource struct {
	Source FileTestData `json:"_source"`
}

type StopWords struct {
	Texts []string `json:"texts"`
	Class string   `json:"class"`
}

type Pair[T any, K any] struct {
	First  T
	Second K
}
