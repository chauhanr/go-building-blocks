package filehandling

import (
	"crypto/sha1"
	"fmt"
	"bytes"
	"crypto/md5"
)

func CalculateSHA(secret string)(string, error){
    hasher := sha1.New()
    b := []byte{}
    data := []byte(secret)

    n, err := hasher.Write(data)
    if n != len(data) || err != nil{
    	return "", err
	}
	checksum := hasher.Sum(b)
	checkSumBuf := bytes.NewBufferString("");
	fmt.Fprintf(checkSumBuf,"%x" ,checksum)
	return checkSumBuf.String(), nil
}

func CalculateMD5 (secret string) (string, error){
	md5Hasher := md5.New()
	b := []byte{}

	data := []byte(secret)
	n, err := md5Hasher.Write(data)

	if n != len(data) || err != nil{
		return "", err
	}
	checksum := md5Hasher.Sum(b)

	checkSumBuf := bytes.NewBufferString("");
	fmt.Fprintf(checkSumBuf,"%x" ,checksum)
	return checkSumBuf.String(), nil
}
