package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Counter struct {
	noOfLines  int
	noOfPunc   int
	noOfWord   int
	noOfVowels int
}

func main() {

	startTime := time.Now()


	if len(os.Args) < 2 {
		fmt.Println("Usage: go run task2.go arg1 arg2 ...")
		return
	}


	fileContent, err := ReadFile("D:/golang tutorials/file1.txt")

	if err != nil {
		log.Fatal(err)
	}

	chunks, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}
	
	channel := make(chan Counter)


	
	Chunklenght := len(fileContent) / chunks

	fmt.Println("No of Chunks", os.Args[2])
	fmt.Printf("\n\n")



	for loop := 0; loop < chunks; loop++ {

		fIndex := loop * Chunklenght
		sIndex := (loop + 1) * Chunklenght
		// fmt.Println(fIndex , sIndex)
		go stringIterator(fileContent[fIndex:sIndex], channel)
		// fIndex = sIndex
		// sIndex += Chunklenght
		Counts := <-channel
		fmt.Printf("No of Words of Chunk %d: %d \n", loop+1, Counts.noOfWord)
		fmt.Printf("No of Lines of Chunk %d: %d \n", loop+1, Counts.noOfLines)
		fmt.Printf("No of Vowels of Chunk %d: %d \n", loop+1, Counts.noOfVowels)
		fmt.Printf("No of Punctuation of Chunk %d: %d \n", loop+1, Counts.noOfPunc)
		fmt.Printf("\n\n")
	}

	fmt.Printf("Execution time: %v\n", time.Since(startTime))

}

func ReadFile(filePath string) (string, error) {
	// Read the file's content into a byte slice
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fileContent := string(content)
	return fileContent, nil
}

func stringIterator(content string, channel chan Counter) {
	// fmt.Println("\n \n \n", content)

	counter := Counter{}

	for _, char := range content {
		switch {
		case char == ' ':
			// noOfWord++
			counter.noOfWord++
		case char == '\n':
			// noOfLines++
			counter.noOfLines++
		case (char < 48 && char > 32) || (char < 65 && char > 57) || (char < 97 && char > 90) || (char < 127 && char > 122):
			// noOfPunc++
			counter.noOfPunc++
		case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
			counter.noOfVowels++

		}
	}

	channel <- counter

}
