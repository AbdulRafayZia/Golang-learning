package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"time"
)

type counter struct{
    noOfLines , noOfPunc , noOfWord , noOfVowels int

}

func main() {

    

	startTime := time.Now()
	fileContent, err := ReadFile("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan int)

	fmt.Printf("type of : %T \n", fileContent[0])
	go stringIterator(fileContent[:len(fileContent)/2], c)
	go stringIterator(fileContent[len(fileContent)/2:], c)
	  noOfLines1 := <-c
	noOfPunc1 := <-c
	noOfVowels1 := <-c
	noOfWord1 := <-c



	noOfLines2 := <-c
	noOfPunc2 := <-c
	noOfVowels2 := <-c
	noOfWord2 := <-c
	// noOfLines ,noOfPunc ,noOfVowels ,noOfWord:=stringIterator(fileContent)

	fmt.Printf("no of Words 1st: %d \n", noOfWord1)
	fmt.Printf("no of Lines 1st: %d \n", noOfLines1)
	fmt.Printf("no of Vowels 1st: %d \n", noOfVowels1)
	fmt.Printf("no of Punctuation 1st: %d \n", noOfPunc1)

    fmt.Printf("\n\n")
	fmt.Printf("no of Words 2nd: %d \n", noOfWord2)
	fmt.Printf("no of Lines 2nd: %d \n", noOfLines2)
	fmt.Printf("no of Vowels 2nd: %d \n", noOfVowels2)
	fmt.Printf("no of Punctuation 2nd: %d \n", noOfPunc2)
    fmt.Printf("\n\n")
	fmt.Printf("no of Words: %d \n", noOfWord1+noOfWord2)
	fmt.Printf("no of Lines: %d \n", noOfLines1+noOfLines1)
	fmt.Printf("no of Vowels: %d \n", noOfVowels1+noOfVowels1)
	fmt.Printf("no of Punctuation: %d \n", noOfPunc1+noOfPunc2)

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

func stringIterator(content string, c chan int) {

	noOfLines := 0
	noOfPunc := 0
	noOfWord := 0
	noOfVowels := 0

	for _, char := range content {
		switch {
		case char == ' ':
			noOfWord++
		case char == '\n':
			noOfLines++
		case (char < 48 && char > 32) || (char < 65 && char > 57) || (char < 97 && char > 90) || (char < 127 && char > 122):
			noOfPunc++
		case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
			noOfVowels++

		}
	}

	c <- noOfLines
	c <- noOfPunc
	c <- noOfVowels
	c <- noOfWord

	// close(c)
	// return noOfLines ,noOfPunc ,noOfVowels ,noOfWord

}

// func WordsCount(content string) int {

// 	//    using ascii code

// 	// words:=s
// 	noOfWords := 0
// 	for i := range content {

// 		if content[i] == ' ' {
// 			noOfWords++
// 		}
// 	}

// 	return noOfWords

// }

// func PunctuationCount(content string) int {
// 	// split := strings.Split(s, "")
// 	// words:=s

// 	noOfPunc := 0

// 	for i := range content {
// 		if (content[i] < 48 && content[i] > 32) || (content[i] < 65 && content[i] > 57) || (content[i] < 97 && content[i] > 90) || (content[i] < 127 && content[i] > 122) {
// 			noOfPunc++
// 		}
// 	}

// 	return noOfPunc
// }
// func CountVowels(content string) int {

// 	noOfVowels := 0

// 	for i := range content {

// 		switch content[i]{
// 		case 97:
// 			noOfVowels++
// 		case 101:
// 			noOfVowels++
// 		case 105:
// 			noOfVowels++
// 		case 111:
// 			noOfVowels++
// 		case 117:
// 			noOfVowels++
// 		case 65:
// 			noOfVowels++
// 		case 69:
// 			noOfVowels++
// 		case 73:
// 			noOfVowels++
// 		case 79:
// 			noOfVowels++
// 		case 85:
// 			noOfVowels++
// 		}
// 	}

// 		// for switch execution time avg 647.19
// 		//for  for loop execution time avg 612.60

// 	// for i := range content {
// 	// 	if content[i] == 97 || content[i] == 101 || content[i] == 105 || content[i] == 111 || content[i] == 117 || content[i] == 65 || content[i] == 69 || content[i] == 73 || content[i] == 79 || content[i] == 85 {
// 	// 		noOfVowels++
// 	// 	}
// 	// }

// 	// 	'a' = 97
// 	// 'e' = 101
// 	// 'i' = 105
// 	// 'o' = 111
// 	// 'u' = 117
// 	// The ASCII values for the uppercase vowels are:

// 	// 'A' = 65
// 	// 'E' = 69
// 	// 'I' = 73
// 	// 'O' = 79
// 	// 'U' = 85
// 	return noOfVowels
// }

//  "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// package main
// import (
//     "fmt"
//     "io/ioutil"
//     "log"
//     "time"
// )
// func CountWords(x string, c chan int) {
//     w := 1
//     for _, words := range x {
//         if words == 32 || words == 10 {
//             w++
//         }
//     }
//     c <- w
// }
// func CountVowels(x string) int {
//     vowl := 0
//     for _, words := range x {
//         if words == 65 || words == 69 || words == 73 || words == 79 || words == 85 || words == 97 || words == 101 || words == 105 || words == 111 || words == 117 {
//             vowl++
//         }
//     }
//     return vowl
// }
// func CountLines(x string) int {
//     line := 1
//     for _, words := range x {
//         if words == 13 {
//             line++
//         }
//     }
//     return line
// }
// func CountPunct(x string) int {
//     pun := 0
//     for _, words := range x {
//         if words >= 33 && words <= 47 || words >= 58 && words <= 64 || words >= 91 && words <= 96 || words >= 123 && words <= 126 {
//             pun++
//         }
//     }
//     return pun
// }
// func main() {
//     start := time.Now()
//     c := make(chan int)
//     content, err := ioutil.ReadFile("file1.txt")
//     if err != nil {
//         log.Fatal(err)
//     }
//     data := string(content)
//     //    fmt.Println(data)
//     go CountWords(data, c)
//     wc := <-c
//     fmt.Println("number of words:", wc)
//     fmt.Println("number of puncuations:", CountPunct(data))
//     fmt.Println("number of lines:", CountLines(data))
//     fmt.Println("number of vowels:", CountVowels(data))
//     fmt.Println("Run Time:", time.Since(start))
// }

// package main
// import (
//     "fmt"
//     "time"
//     "io/ioutil"
//     "log"
// )
// func Counts (x string, c chan int)   {
//     lineCount:=1;
//     wordsCount:=1;
//     vowelsCount:=0
//     puncuationsCount:=0
//     for  _,words:=range x{
//         switch{
//         case words == 13 :
//             lineCount++
//         case words == 32 || words == 10 :
//             wordsCount++
//         case words == 65 || words == 69 || words == 73 || words == 79 || words == 85 || words == 97 || words == 101 || words == 105 || words == 111 || words == 117 :
//             vowelsCount++
//         case words >= 33 && words <= 47 || words >= 58 && words <= 64 || words >= 91 && words <= 96 || words >= 123 && words <= 126 :
//             puncuationsCount++
//         }
//     }
//     // return  wordsCount,lineCount,vowelsCount,puncuationsCount
//     c<-lineCount
//     c <-wordsCount
//     c<- vowelsCount
//     c<-puncuationsCount
// }
// func main(){
//     start:=time.Now()
//     c := make(chan int)
//     content, err := ioutil.ReadFile("file1.txt")
//     if err != nil {
//         log.Fatal(err)

//     }

//    fileData := string(content)
//    go Counts(fileData, c)
//    LineCount:=<-c
//    WordsCount:=<-c
//    VowelsCount:=<-c
//    PuncuationsCount:=<-c
//     // wc := <-c
//     fmt.Println("number of lines:", LineCount)
//     fmt.Println("number of words:",WordsCount )
//     fmt.Println("number of vowels:", VowelsCount)
//     fmt.Println("number of puncuations:", PuncuationsCount )
//     fmt.Println("Run Time:", time.Since(start))
// }
