package tokenization

import (
	"fmt"
	"os"
	"testing"
)

func Test_WordFromString(t *testing.T) {

	var mInput string
	var mOutput []map[string]uint32
	var mError error

	// Example string...
	mInput = "word1 word2 word3 word4 word5 word6 word7 word99 word1 word2 word3 word4 word5 word6 word100 word7 word1 word2 word3 word4 word5 word6 word7 word1 word2 word3 word4 word5 word6 word7"

	// Testing mapping...
	mOutput, mError = WordFromString(mInput, 10)
	if mError != nil {
		os.Exit(1)
	}

	fmt.Println(mOutput)
}
