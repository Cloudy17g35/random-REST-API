package external_requests

import (
	"net/http"
	"time"
	"io/ioutil"
	"random-api/logger"
)


const RandomURL = "https://www.random.org/integers/?min=1&max=100&col=1&base=10&format=plain&rnd=new"


func GetRandomNumbers(length string) (int, string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	logs := logger.GetLogger()
	req, _ := http.NewRequest(http.MethodGet, RandomURL, nil)
	var empty_res string
	q := req.URL.Query()
	q.Add("num", length)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req,)
	if err != nil {
		format := "error occured while connecting to external service: %v"
		logs.Errorf(
			format,
			err)
		return 0,empty_res,err
	}
	
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		format := "error occured while connecting to external service: %v"
		logs.Errorf(
			format,
			err)
		return 0,empty_res,err
	}

	return resp.StatusCode, string(responseBody), nil

}