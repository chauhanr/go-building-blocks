package sorting

import "errors"

func TransposeMatrix(metrix [][]int) error{
	rowCount := len(metrix)
	if rowCount == 0{
		return errors.New("metrix cannot be empty")
	}
	columnCount := len(metrix[0])
	if columnCount != rowCount{
		return errors.New("metrix need to be a square metrix")
	}

	var i, j int
	for i=0; i<rowCount; i++{
		for j=i; j<columnCount; j++{
			if i == j {
				// do nothing
			}else{
				temp := metrix[j][i]
				metrix[j][i] = metrix[i][j]
				metrix[i][j] = temp
			}
		}
	}

	return nil
}