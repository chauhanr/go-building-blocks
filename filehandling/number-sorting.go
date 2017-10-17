package filehandling

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"strconv"
)


const ARRAY_LENGTH=10000000

// function that read form a file and prints output
func ReadFile(path string) error {
	inputFile, inputErr := os.Open(path)

	if inputErr != nil{
		return inputErr
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, _, readErr:= inputReader.ReadLine()
		if readErr == io.EOF {
			return readErr
		}
		fmt.Println(string(inputString))
	}
	return nil
}


// function that reads a file bunch at a time using scanner
func ReadFileToSparseArray(path string) ([ARRAY_LENGTH]int, error){
	var bitArray [ARRAY_LENGTH]int

	inputFile, inputErr := os.Open(path)
	//fmt.Printf("Reading file %s", path)
	if inputErr != nil{
		return bitArray, inputErr
	}
	defer inputFile.Close()
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan(){
    	var line = scanner.Text()
		//fmt.Printf("value %s", line)
    	number, err := strconv.Atoi(line)
    	if err != nil{
    		return bitArray, err
		}
		//fmt.Printf("value %d", number)
		bitArray[number] = 1;
	}
	return bitArray, nil
}

/* test if the file is storted. this checks if every number if larger than the previous number read from the
 the file. */
func IsFileSorted(path string) (bool,error){
     sortedFile, err := os.Open(path)
     if err != nil {
     	return false, err
	 }
	 defer sortedFile.Close()

	 scanner := bufio.NewScanner(sortedFile)
	 prev := -1
     for scanner.Scan(){
     	curr, errConvert := strconv.Atoi(scanner.Text());
     	if errConvert != nil{
     		return false, nil
		}
     	if prev > curr {
     		return false, nil
		}
		prev = curr
	 }
	return true, nil
}

/*
	Sort large file will take a file of integer type and then sort the file
	write to an output file It iwll use the ReadFileToSparseArray function to get
    a sparse array of the original file content.
	Once the sparse array is read then it is fed to the output file by simply iterating the array
	index to the output file.
*/
func SortLargeFile(inputPath string, outputPath string) error{
	bitArray, err := ReadFileToSparseArray(inputPath)
	if err != nil{
		return err;
	}

	outFile, outErr := os.OpenFile(outputPath,os.O_WRONLY | os.O_CREATE,0666)
    if outErr != nil{
    	return outErr
	}

	defer outFile.Close()
	outWriter := bufio.NewWriter(outFile)
	for index, val := range bitArray {
		if val != 0 {
			outWriter.WriteString(fmt.Sprintf("%d\n", index))
		}
	}
	outWriter.Flush()
	return nil
}

