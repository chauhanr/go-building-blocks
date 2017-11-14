package concurrency_patterns

import (
	"testing"
)

var urlTestCase = []struct{
	Urls []string
	valid []bool
}{
	{
		[]string{"http://www.google.com","http://block"},
		[]bool{true, false},
	},
}

func TestCheckUrlsValidity(t *testing.T) {

	for _, urlCase := range urlTestCase {
			urls := urlCase.Urls
			results := CheckUrlsValidity(urls)
			for index, result := range results{
				 if (result.Error == nil) !=  urlCase.valid[index]{
				 	t.Errorf("Url %s is supposed to be valid : %s but was found to be %s", urls[index], "true", urlCase.valid[index])
				 }
				 if result.Error != nil{
				 	t.Logf("Error in accessing url: %s \n", result.Error.Error())
				 }
			}
	}

}
