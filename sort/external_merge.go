package sort

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/niftynei/algos/timing"
)

func generateBigFile(filename string, length int) {
	file, err := os.Create(filename)

	file, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < length; i++ {
		_, err := file.WriteString(fmt.Sprintf("%d\n", rand.Intn(math.MaxInt64)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ExternalMerge(inputFilename string, numToSort int) (outputFilename string) {
	defer timing.Timer(time.Now(), "External Merge")

	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		generateBigFile(inputFilename, numToSort)
	}

	outputFilename = inputFilename
	return outputFilename
}
