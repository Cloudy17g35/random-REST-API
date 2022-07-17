package internal_requests


import (
	"net/http"
	"time"
	"io/ioutil"
	"random-api/logger"
)


const apiURL = "http://localhost:8080/random/mean"


func GetNumbersAndStd(
	requests string,
	length string) (int, string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	logs := logger.GetLogger()
	req, _ := http.NewRequest(
		http.MethodGet, 
		apiURL, 
		nil)
	var empty_res string
	q := req.URL.Query()
	q.Add("requests", requests)
	q.Add("length", length)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req,)
	if err != nil {
		format := "error occured while connecting to internl service: %v"
		logs.Errorf(
			format,
			err)
		return 0,empty_res,err
	}
	
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		format := "error occured while connecting to internal service: %v"
		logs.Errorf(
			format,
			err)
		return 0,empty_res,err
	}

	return resp.StatusCode, string(responseBody), nil

}