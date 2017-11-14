package structures

import (
	"errors"
	"fmt"
)

func PrintBanner(letter string) ([]int, error){
	if len(letter) != 1{
		return nil, errors.New("Cannot work with multiple letters.")
	}
    if val, ok := letterMap[letter]; ok {
		printLetterBanner(val)
		return val, nil
	}else{
		return nil, errors.New("letter is not supported")
	}
	return nil, nil
}

func printLetterBanner(set []int){

	if set[0] == 1{
		PrintPos0_6_3()
	}
	if set[5] == 1 && set[1] == 1{
		PrintPos5_1_4_2()
	}else if set[5] == 1 && set[1] == 0{
		PrintPos5_4()
	}
	if set[6] == 1{
		PrintPos0_6_3()
	}
	if set[4] == 1 && set[2] ==1 {
		PrintPos5_1_4_2()
	}else if set[4] == 1 && set[2] == 0{
		PrintPos5_4()
	}

	if set[3] == 1{
		PrintPos0_6_3()
	}
}

func PrintPos0_6_3(){
	fmt.Println(" "+"-"+" ")
}

func PrintPos5_4(){
	fmt.Println("|")
}

func PrintPos5_1_4_2(){
	fmt.Println("| |")
}




var letterMap = map[string][]int{
		"a": []int{1,1,1,0,1,1,1},
		"A": []int{1,1,1,0,1,1,1},
		"b": []int{1,1,1,1,1,1,1},
	  	"B": []int{1,1,1,1,1,1,1},
	  	"c": []int{1,0,0,1,1,1,0},
	  	"C": []int{1,0,0,1,1,1,0},
		"d": []int{1,1,1,1,1,1,0},
		"D": []int{1,1,1,1,1,1,0},
}