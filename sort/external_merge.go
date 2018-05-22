package sort

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"io"
	"os"
	"time"
	"io/ioutil"

	"github.com/niftynei/algos/timing"
)

func generateBigFile(filename string, length int) {
	file, err := os.Create(filename)

	file, err = os.OpenFile(
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

func chunkFile(filename string, chunkSize int) string {
	// Given a filename, open the file and chunk it into
	// a set of files that at most chunkSize.
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Make a temporary directory.
	directoryName, err := ioutil.TempDir("", "chunks")
	if err != nil {
		panic(err)
	}
	
	// While there's still stuff to read, output chunk bytes	
	buffer := make([]byte, chunkSize)
	for {
		bytesread, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break;
			}	
			panic(err)
		}
		chunk := sortChunk(buffer[:bytesread])
		writeChunk(chunk, directoryName)
	}		
	
	return directoryName
}

func sortChunk(chunk []byte) []byte {
			
	return sortedChunk
}

func chunkToIntArray(chunk []byte) []int {
	ints := make([]int)	
	newline := []byte("\n")
	for bite, index := range chunk {
		bites = make([]byte)
		if bite == newline {
				
		}	
	}
	index := bytes.Index(chunk, newline)
	if index == -1 {
		// todo:
		return nil
	}
	int(chunk[index+	
}

func writeChunk(chunk []byte, directory string) {
	tmpFile, err := ioutil.TempFile(directory, "chunk_")
	if err != nil {
		panic(err)
	}
	if _, err := tmpFile.Write(chunk); err != nil {
		panic(err)
	}
	if err := tmpFile.Close(); err != nil {
		panic(err)
	}
}

func printFileNames(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ExternalMerge(inputFilename string, numToSort int) (outputFilename string) {
	defer timing.Timer(time.Now(), "External Merge")

	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		generateBigFile(inputFilename, numToSort)
	}
	
	chunkDir := chunkFile(inputFilename, 1024)
	printFileNames(chunkDir)	
	fmt.Printf("directory is %s\n", chunkDir)
	return chunkDir
}
