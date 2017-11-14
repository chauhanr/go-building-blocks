package concurrency_patterns

import (
	"net/http"
)

type Result struct{
	Error error
	Response *http.Response
}


func CheckUrlsValidity(urls []string) []Result {
    var testResults []Result

    checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result{
		results := make(chan Result)
		go func(){
			defer close(results)
			for _, url := range urls{
				var result Result
				resp, err := http.Get(url)
				result = Result{Error : err, Response: resp}
				select{
				    case <- done:
				    	return
				    	case results <-result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	for result := range checkStatus(done, urls...){
		testResults = append(testResults, Result{Error: result.Error, Response: result.Response})
	}

	return testResults
}
