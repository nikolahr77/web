package api_test

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/web"
	"github.com/web/api"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockCampaignRepository struct {
	mock.Mock
}

func (m *MockCampaignRepository) Get(id string) (web.Campaign, error) {
	args := m.Called(id)
	return args.Get(0).(web.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) Create(cam web.RequestCampaign) (web.Campaign, error) {
	args := m.Called(cam)
	return args.Get(0).(web.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCampaignRepository) Update(id string, cam web.RequestCampaign) (web.Campaign, error) {
	args := m.Called(id, cam)
	return args.Get(0).(web.Campaign), args.Error(1)
}

func adaptToSegmentation(s web.RequestSegmentation) web.Segmentation {
	return web.Segmentation{
		Address: s.Address,
		Age:     s.Age,
	}
}

func TestCreateCampaign(t *testing.T) {
	campaign := `{"name":"Test Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("POST", "/campaign", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	cr := web.RequestCampaign{
		Name: "Test Campaign",
		Segmentation: web.Segmentation{
			Address: "Sofia 1512",
			Age:     12,
		},
	}

	c := web.Campaign{
		GUID:      "334",
		Name:      "Test Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
		Segmentation: web.Segmentation{
			Address: "Sofia 1512",
			Age:     12,
		},
	}

	testObj := new(MockCampaignRepository)

	testObj.On("Create", cr).Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/campaign", api.CreateCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := api.CampaignDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.CampaignDTO{
		GUID:      "334",
		Name:      "Test Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
		Segmentation: api.SegmentationDTO{
			Address: "Sofia 1512",
			Age:     12,
		},
	}
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}
