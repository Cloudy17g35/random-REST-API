package main

import (
	"random-api/internal_requests"
	"testing"
	"fmt"
)


func TestValidationErrorBadTypes(t *testing.T) {
	requests := "r"
	length := "l"
	
	status_code, 
	_, 
	_ := 
	internal_requests.GetNumbersAndStd(
		requests,
		length)
	
	expexted_status_code := 400

	if status_code != expexted_status_code {
		m := fmt.Sprintf("expected status code: %d, but get %d", 
		expexted_status_code, 
		status_code)
		t.Errorf(m)
	}
}


func TestValidationErrorInvalidRequests(t *testing.T) {
	requests := "1321"
	length := "2"
	
	status_code, 
	_, 
	_ := 
	internal_requests.GetNumbersAndStd(
		requests,
		length)
	
	expexted_status_code := 400

	if status_code != expexted_status_code {
		m := fmt.Sprintf("expected status code: %d, but get %d", 
		expexted_status_code, 
		status_code)
		t.Errorf(m)
	}
}


func TestValidationErrorInvalidLength(t *testing.T) {
	requests := "1"
	length := "1"
	
	status_code, 
	_, 
	_ := 
	internal_requests.GetNumbersAndStd(
		requests,
		length)
	
	expexted_status_code := 400

	if status_code != expexted_status_code {
		m := fmt.Sprintf("expected status code: %d, but get %d", 
		expexted_status_code, 
		status_code)
		t.Errorf(m)
	}
}


func TestValidRequest(t *testing.T) {
	requests := "1"
	length := "2"
	
	status_code, 
	_, 
	_ := 
	internal_requests.GetNumbersAndStd(
		requests,
		length)
	
	expexted_status_code := 200

	if status_code != expexted_status_code {
		m := fmt.Sprintf("expected status code: %d, but get %d", 
		expexted_status_code, 
		status_code)
		t.Errorf(m)
	}
}