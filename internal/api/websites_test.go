package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suhaan-Bhandary/website-checker/internal/app/websites/mocks"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestAddWebsitesHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := AddWebsitesHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success Added website",
			input: `
				{
					"websites": ["hi.com"]
				}
			`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("InsertWebsites", mock.Anything).Return().Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "Fail Invalid Json",
			input:              ``,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Failed because of not websites",
			input:              `{}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/websites", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetWebsitesHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := GetWebsitesHandler(websitesSvc)

	tests := []struct {
		name               string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success Get websites",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetWebsites", mock.Anything).Return([]string{}).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", "/websites", bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestDeleteWebsiteHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	deleteWebsitesHandler := DeleteWebsiteHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success Get websites",
			input: "name",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteWebsite", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Fail as no search query",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "Fail as delete website service failed",
			input: "name",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteWebsite", mock.Anything).Return(errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("/websites?website=%s", test.input), bytes.NewBuffer([]byte("")),
			)

			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(deleteWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetWebsitesStatusHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	getWebsitesStatusHandler := GetWebsiteStatusHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Success Get specific website status",
			input: "name",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetWebsiteStatus", mock.Anything).Return(repository.WebsitesStatus{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "Success Get all website status",
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetAllStatus", mock.Anything).Return(map[string]repository.WebsitesStatus{}).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "Fail as website not found",
			input: "name",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetWebsiteStatus", mock.Anything).Return(repository.WebsitesStatus{}, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("/websites/status?website=%s", test.input), bytes.NewBuffer([]byte("")),
			)

			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getWebsitesStatusHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
