package filehandling

import (
	"testing"
	"io"
	"os"
)


func TestBasicReadFunc(t *testing.T) {
    err :=  ReadFile("snippet-number.txt")
	if err != nil && err != io.EOF{
		t.Errorf("Error in reading file %s \n",err.Error())
	}else{
		t.Log("Test passed successfully")
	}
}

func TestReadFileToSparseArray(t *testing.T) {
	bitArray, err := ReadFileToSparseArray("snippet-number.txt")
	if err != nil{
		t.Errorf("Error reading file : %s", err.Error())
	}else{
		for index, value := range bitArray{
			if value != 0{
				t.Logf("%d", index);
			}
		}
	}
}


func TestFileSortFunction(t *testing.T){
	inputPath := "number.txt"
	outputPath := "sorted-file.txt"
	// first delete the outputFile
	os.Remove(outputPath)

	err := SortLargeFile(inputPath,outputPath)
	if err != nil {
		t.Errorf("Error while sorting large file : %s\n", err.Error())
	}else {
		isSorted, err := IsFileSorted(outputPath)
		if err != nil{
			t.Errorf("Error while reading output sorted file : %s\n", err.Error())
		}else if !isSorted {
			t.Errorf("File %s should be sorted but was found to be unsorted")
		}
	}
}

func BenchmarkSortLargeFile(b *testing.B){

	inputPath := "number.txt"
	outputPath := "sorted-file.txt"
	// first delete the outputFile

	for n := 0; n <b.N; n++ {
		os.Remove(outputPath)
		err := SortLargeFile(inputPath, outputPath)
		if err != nil {
			b.Errorf("Error while sorting large file : %s\n", err.Error())
		} else {
			isSorted, err := IsFileSorted(outputPath)
			if err != nil {
				b.Errorf("Error while reading output sorted file : %s\n", err.Error())
			} else if !isSorted {
				b.Errorf("File %s should be sorted but was found to be unsorted")
			}
		}
	}
}