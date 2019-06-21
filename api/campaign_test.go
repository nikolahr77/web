package api_test

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/web"
	"github.com/web/api"
	"net/http"
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

func TestCreateCampaignError(t *testing.T) {
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

	testObj := new(MockCampaignRepository)

	testObj.On("Create", cr).Return(web.Campaign{}, errors.New("Test error"))

	r := mux.NewRouter()
	r.Handle("/campaign", api.CreateCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestCreateCampaignMalformedJson(t *testing.T) {
	campaign := `{"name":551,"segmentation":{"address":"Sofia 1512","age":"two"}}`
	req := httptest.NewRequest("POST", "/campaign", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/campaign", api.CreateCampaign(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestDeleteCampaign(t *testing.T) {
	campaign := `{"name":"Delete Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("DELETE", "/campaign/325b", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	testObj := new(MockCampaignRepository)

	testObj.On("Delete", "325b").Return(nil)

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.DeleteCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 200
	assert.Equal(t, expected, actual)
}

func TestDeleteCampaignReturnError(t *testing.T) {
	campaign := `{"name":"Delete Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("DELETE", "/campaign/325b", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	testObj := new(MockCampaignRepository)

	testObj.On("Delete", "325b").Return(errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.DeleteCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
}

func TestUpdateCampaign(t *testing.T) {
	campaign := `{"name":"Update Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("POST", "/campaign/9f-245", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	cr := web.RequestCampaign{
		Name: "Update Campaign",
		Segmentation: web.Segmentation{
			Address: "Sofia 1512",
			Age:     12,
		},
	}

	c := web.Campaign{
		GUID:      "9f-245",
		Name:      "Updated Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(152, 0).UTC(),
		UpdatedOn: time.Unix(440, 0).UTC(),
		Segmentation: web.Segmentation{
			Address: "Vartna 366",
			Age:     71,
		},
	}

	testObj := new(MockCampaignRepository)

	testObj.On("Update", "9f-245", cr).Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.UpdateCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := api.CampaignDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.CampaignDTO{
		GUID:      "9f-245",
		Name:      "Updated Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(152, 0).UTC(),
		UpdatedOn: time.Unix(440, 0).UTC(),
		Segmentation: api.SegmentationDTO{
			Address: "Vartna 366",
			Age:     71,
		},
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)

	testObj.AssertExpectations(t)
}

func TestUpdateCampaignReturnError(t *testing.T) {
	campaign := `{"name":"Update Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("POST", "/campaign/9f-245", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	cr := web.RequestCampaign{
		Name: "Update Campaign",
		Segmentation: web.Segmentation{
			Address: "Sofia 1512",
			Age:     12,
		},
	}

	testObj := new(MockCampaignRepository)

	testObj.On("Update", "9f-245", cr).Return(web.Campaign{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.UpdateCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestUpdateCampaignMalformedJson(t *testing.T) {
	campaign := `{"name":51251,"segmentation":{"address":"Sofia 1512","age":"HEY"}}`
	req := httptest.NewRequest("POST", "/campaign/9bb", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.UpdateCampaign(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestGetCampaign(t *testing.T) {
	campaign := `{"name":"Get Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("GET", "/campaign/9f-245", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	c := web.Campaign{
		GUID:      "9f-245",
		Name:      "Get Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(152, 0).UTC(),
		UpdatedOn: time.Unix(440, 0).UTC(),
		Segmentation: web.Segmentation{
			Address: "Sofia 1512",
			Age:     12,
		},
	}

	testObj := new(MockCampaignRepository)

	testObj.On("Get", "9f-245").Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.GetCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := api.CampaignDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.CampaignDTO{
		GUID:      "9f-245",
		Name:      "Get Campaign",
		Status:    "draft",
		CreatedOn: time.Unix(152, 0).UTC(),
		UpdatedOn: time.Unix(440, 0).UTC(),
		Segmentation: api.SegmentationDTO{
			Address: "Sofia 1512",
			Age:     12,
		},
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)

	testObj.AssertExpectations(t)
}

func TestGetCampaignReturnError(t *testing.T) {
	campaign := `{"name":"Get Campaign","segmentation":{"address":"Sofia 1512","age":12}}`
	req := httptest.NewRequest("GET", "/campaign/9f-245", strings.NewReader(campaign))
	w := httptest.NewRecorder()

	testObj := new(MockCampaignRepository)

	testObj.On("Get", "9f-245").Return(web.Campaign{}, errors.New("Test error"))

	r := mux.NewRouter()
	r.Handle("/campaign/{id}", api.GetCampaign(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}
