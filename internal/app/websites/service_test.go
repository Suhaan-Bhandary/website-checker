package websites

import (
	"testing"

	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestInsertWebsites(t *testing.T) {
	websitesRepo := mocks.NewWebsitesStorer(t)
	service := NewService(websitesRepo)

	tests := []struct {
		name  string
		input []string
		setup func(websitesMock *mocks.WebsitesStorer)
	}{
		{
			name:  "Success insert website",
			input: []string{"test.com"},
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("InsertWebsite", mock.Anything).Return().Once()
			},
		},
		{
			name:  "Success 2 insert website",
			input: []string{"test.com", "test2.com"},
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("InsertWebsite", mock.Anything).Return().Times(2)
			},
		},
		{
			name:  "Success for empty insert website",
			input: []string{},
			setup: func(websitesMock *mocks.WebsitesStorer) {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup(websitesRepo)
			service.InsertWebsites(test.input)
		})
	}
}

func TestGetWebsites(t *testing.T) {
	websitesRepo := mocks.NewWebsitesStorer(t)
	service := NewService(websitesRepo)

	tests := []struct {
		name  string
		setup func(websitesMock *mocks.WebsitesStorer)
	}{
		{
			name: "Success Get website",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("GetWebsites").Return([]string{}).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup(websitesRepo)
			service.GetWebsites()
		})
	}
}

func TestDeleteWebsite(t *testing.T) {
	websitesRepo := mocks.NewWebsitesStorer(t)
	service := NewService(websitesRepo)

	tests := []struct {
		name            string
		input           string
		setup           func(websitesMock *mocks.WebsitesStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success Get all website",
			input: "all",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("DeleteAllWebsites").Return().Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Success Get website",
			input: "name",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("IsWebsitePresent", mock.Anything).Return(true).Once()
				websitesMock.On("DeleteWebsite", mock.Anything).Return().Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Success Get website",
			input: "not-found",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("IsWebsitePresent", mock.Anything).Return(false).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup(websitesRepo)
			err := service.DeleteWebsite(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetAllStatus(t *testing.T) {
	websitesRepo := mocks.NewWebsitesStorer(t)
	service := NewService(websitesRepo)

	tests := []struct {
		name  string
		input string
		setup func(websitesMock *mocks.WebsitesStorer)
	}{
		{
			name:  "Success",
			input: "all",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("GetAllWebsiteStatus").Return(map[string]repository.WebsitesStatus{}).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup(websitesRepo)
			res := service.GetAllStatus()

			if res == nil {
				t.Errorf("No response found")
			}
		})
	}
}

func TestGetWebsiteStatus(t *testing.T) {
	websitesRepo := mocks.NewWebsitesStorer(t)
	service := NewService(websitesRepo)

	tests := []struct {
		name            string
		input           string
		setup           func(websitesMock *mocks.WebsitesStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success website status",
			input: "test",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("IsWebsitePresent", mock.Anything).Return(true).Once()
				websitesMock.On("GetWebsiteStatus", mock.Anything).Return(repository.WebsitesStatus{}).Once()
			},
			isErrorExpected: false,
		},
		{
			name:  "Fail website not found",
			input: "test",
			setup: func(websitesMock *mocks.WebsitesStorer) {
				websitesMock.On("IsWebsitePresent", mock.Anything).Return(false).Once()
			},
			isErrorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup(websitesRepo)
			_, err := service.GetWebsiteStatus(test.input)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
