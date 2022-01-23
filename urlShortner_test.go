package main

import (
	"github.com/speps/go-hashids"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)


type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
	expectedUrl        string
}
type TestStruct1 struct {
	requestParam        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}
const (
	BadRequestCode     = 400
	SuccessRequestCode = 200
	StatusCreated = 201
)

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {
	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \nexpectedurl: %s\n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode, test.expectedUrl)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \nexpectedurl: %s \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode, test.expectedUrl)
		}
	}
}
func DisplayTestCaseResults1(functionalityName string, tests []TestStruct1, t *testing.T) {
	for _, test := range tests {
		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestParam, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestParam, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}

func TestCreateShorturl(t *testing.T) {
	url := "http://localhost:3000/"
	var id string
	tests := []TestStruct{
		{``, BadRequestCode, "", 0, ""},
		{`{"originalUrl":""}`, BadRequestCode, "", 0, ""},
		{`{"originalUrl":"https://www.youtube.com/watch?v=CBVJTplw4cE"}`, StatusCreated, "", 0,""},
		{`{"originalUrl":"https://www.youtube.com/watch?v=CBVJTplw4cE"}`, SuccessRequestCode, "", 0,""},
	}
	tests1 :=[]TestStruct1{
		{"", SuccessRequestCode, "", 0},
		{"ABDCE", BadRequestCode, "", 0},
	}
	for i, testCase := range tests {
		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
		request, err := http.NewRequest("POST", url, reader)
		res, err := http.DefaultClient.Do(request)
		if i ==2 {
			hd := hashids.NewData()
			h, _ := hashids.NewWithData(hd)
			now := time.Now()
			id, _ = h.Encode([]int{int(now.Unix())})
			ShortenedUrl := "http://localhost:3000/" + id
			tests[2].expectedUrl = ShortenedUrl
		}
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode
	}
	for j, testcase1:= range tests1{
		var reader io.Reader
		if j ==0{
			testcase1.requestParam  = id
		}
		url2 := "http://localhost:3000/url/" + testcase1.requestParam
		request, err := http.NewRequest("GET", url2, reader)
		res, err := http.DefaultClient.Do(request)
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		tests1[j].responseBody = strings.TrimSpace(string(body))
		tests1[0].requestParam = id
		tests1[j].observedStatusCode = res.StatusCode
	}
	DisplayTestCaseResults("ShortUrl", tests, t)
	DisplayTestCaseResults1("getByShortUrl", tests1, t)
}
