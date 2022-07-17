package handlers

import (
	"net/http"
	"random-api/api_response"
	"random-api/converters"
	"random-api/external_requests"
	"random-api/invalid_responses"
	"random-api/logger"
	"random-api/stat_funcs"
	"random-api/validator"
	"github.com/gin-gonic/gin"
)


func getRandomNumbersAndStd(
	context *gin.Context, 
	signal chan<- struct{}) {
	logs := logger.GetLoger()
	validation_errors := 
	validator.ValidateRequest(context)
	if len(validation_errors) > 0 {
		logs.Errorf(
		"validation errors: %v", 
		validation_errors)
		invalid_responses.AbortValidationError(
			context, 
			validation_errors)
		signal <- struct{}{}
		return
		}
	
	r, l := external_requests.QueryRequest(context)
	rInt, err := converters.StrToInt(r)
	if err != nil{
		logs.Errorf(
			"Error occured while converting string: %v", 
			err,
		)
		invalid_responses.AbortInternalServerError(
			context,
		)
		signal <- struct{}{}
	}

	all_results := api_resp.New()
	var all_nums []float64
	for i := 0; i < rInt; i++ {
		status_code, 
		resp_body, 
		err := 
		external_requests.GetRandomNumbers(l)

		if err != nil {
			logs.Errorf(
			"Error occured while formating random numbers: %v", 
			err,
		)
			invalid_responses.AbortExternalServerError(
				context,
				err,
				)
			signal <- struct{}{}
			return
		}

		if status_code != 200 {
			logs.Errorf("Invalid response code from random site")
			invalid_responses.AbortExternalServerBadStatus(
				context,
				status_code,
				resp_body,
				)
			signal <- struct{}{}
			return
		}
		
		random_nums := resp_body
		trimed := converters.TrimString(random_nums, "\n")
		splited := converters.SplitString(trimed, "\n")
		numbers, err := converters.ConvertArrayOfStringsToFloats(splited)
		if err != nil{
			invalid_responses.AbortInternalServerError(
				context,
			)
			signal <- struct{}{}
			return
		}

		stddev, err := stat_funcs.StdDev(numbers)
		if err != nil {
			invalid_responses.AbortInternalServerError(
			context,
			)
			signal <- struct{}{}
			return
			
		}

		single_res := api_resp.SingleResult{
			Stddev: stddev,
			Data: numbers,
		}
		all_results.Add(single_res)
		all_nums = append(all_nums, numbers...)
	}

	if rInt > 1 {
		stddev, err := stat_funcs.StdDev(all_nums)
		if err != nil {
			logs.Errorf(
			"Error occured while calculating std: %v",
			err)
			invalid_responses.AbortInternalServerError(
			context,
			)
			signal <- struct{}{}
			return
		}
		res_for_all := api_resp.SingleResult {
			Stddev: stddev,
			Data: all_nums,
		}
		all_results.Add(res_for_all)
	}
	
	context.IndentedJSON(http.StatusOK, all_results.StDev)
	signal <- struct{}{}
	}


func RandomMeanHandler(context *gin.Context) {
	logs := logger.GetLoger()
	signal := make(chan struct{}, 1)

	go getRandomNumbersAndStd(context, signal)

	select {
		case <-signal:
		close(signal)

	case <-context.Request.Context().Done():
		logs.Errorf("request cancelation")
		return
	}

}
