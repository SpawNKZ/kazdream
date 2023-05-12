package entities

const (
	FileName           = "mobydick.txt"
	CountWordsForPrint = 20
)

type Word struct {
	Value []byte
	Count int
}
