package filehandling

import "testing"

var checksumSHATests = []struct{
	value string
	checksum string
}{
	{"happy", "3978d009748ef54ad6ef7bf851bd55491b1fe6bb"},
	{"charlie", "d8cd10b920dcbdb5163ca0185e402357bc27c265"},
}

var checksumMD5Tests = []struct{
	value string
	checksum string
}{
	{"happy", "56ab24c15b72a457069c5ea42fcfc640"},
	{"charlie", "bf779e0933a882808585d19455cd7937"},
}


func TestSHAchecksum(t *testing.T){

	for _,checksumTestCase := range checksumSHATests{

		secret := checksumTestCase.value
		checksum := checksumTestCase.checksum

		result, err := CalculateSHA(secret)
		if err != nil{
			t.Errorf("There was an error calcuating the SHA checksum for %s - %s ", secret, err.Error())
		}
		if checksum != result {
			t.Errorf("Secret %s must have a SHA checksum %s but was %s", secret, checksum, result)
		}
		//t.Logf("secret %s testcase checksum %s result checksum %s", secret, checksum, result )
	}

}

func TestMD5checksum(t *testing.T){

	for _,checksumTestCase := range checksumMD5Tests{

		secret := checksumTestCase.value
		checksum := checksumTestCase.checksum

		result, err := CalculateMD5(secret)
		if err != nil{
			t.Errorf("There was an error calcuating the MD5 checksum for %s - %s ", secret, err.Error())
		}
		if checksum != result {
			t.Errorf("Secret %s must have a MD5 checksum %s but was %s", secret, checksum, result)
		}
		//t.Logf("secret %s testcase checksum %s result checksum %s", secret, checksum, result )
	}

}