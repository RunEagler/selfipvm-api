package controller

import (
	"testing"
	"selfipvm-api/app"
	"net/http"
	"selfipvm-api/common"
	"selfipvm-api/app/test"
	"selfipvm-api/repository"
)

func TestActivityController(t *testing.T) {

	type request struct {
		payload app.ActivityPayload
	}
	type response struct {
		statusCode int
	}
	type testCase struct {
		testMessage string
		request     request
		response    response
	}
	var testCases []testCase

	activityRepository := repository.NewActivityRepository(postgresqlx)
	ctrl := NewActivityController(service, activityRepository)

	createTestData := func(testMessage string, statusCode int, payload app.ActivityPayload) testCase {
		return testCase{
			testMessage: testMessage,
			request: request{
				payload: payload,
			},
			response: response{
				statusCode: statusCode,
			},
		}
	}

	testCases = append(testCases, createTestData("normal", http.StatusOK, app.ActivityPayload{
		Type:    1,
		Content: common.ToPtr("serverside programing"),
		Minutes: 60,
		Date:    "2018/03/05",
	}))
	testCases = append(testCases, createTestData("required_normal", http.StatusOK, app.ActivityPayload{
		Type:    2,
		Minutes: 30,
		Date:    "2018/03/05",
	}))
	testCases = append(testCases, createTestData("Date Format Error1", http.StatusBadRequest, app.ActivityPayload{
		Type:    1,
		Content: common.ToPtr("サーバサイドプログラミング"),
		Minutes: 140,
		Date:    "2018/3/05",
	}))
	testCases = append(testCases, createTestData("Date Format Error2", http.StatusBadRequest, app.ActivityPayload{
		Type:    1,
		Content: common.ToPtr("サーバサイドプログラミング"),
		Minutes: 400,
		Date:    "2018/03/5",
	}))
	testCases = append(testCases, createTestData("Date Format Error3", http.StatusBadRequest, app.ActivityPayload{
		Type:    1,
		Content: common.ToPtr("サーバサイドプログラミング"),
		Minutes: 400,
		Date:    "2018-03-05",
	}))

	for _, testCase := range testCases {
		switch testCase.response.statusCode {
		case http.StatusOK:
			test.EntryActivityOK(t, nil, service, ctrl, &testCase.request.payload)
		case http.StatusBadRequest:
			test.EntryActivityBadRequest(t, nil, service, ctrl, &testCase.request.payload)
		case http.StatusInternalServerError:
			test.EntryActivityInternalServerError(t, nil, service, ctrl, &testCase.request.payload)
		}
	}

}
