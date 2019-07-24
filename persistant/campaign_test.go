package persistant_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestCampaignRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rowsCam := sqlmock.NewRows([]string{"guid", "name", "status", "created_on", "updated_on"}).
		AddRow("8fsa-321nfuv", "New Campaign", "draft", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())

	rowsSeg := sqlmock.NewRows([]string{"address", "age", "campaign_id", "id"}).
		AddRow("Sofia", 3, "8fsa-321nfuv", "4")

	mock.ExpectQuery("SELECT \\* FROM users CROSS JOIN segmentation").
		WithArgs("8fsa-321nfuv").
		WillReturnRows(rowsCam, rowsSeg)

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)
	myDB := persistant.NewUserRepository(db, clock)

	actual, err := myDB.Get("15")

	expectedCam := web.Campaign{
		GUID:      "8fsa-321nfuv",
		Name:      "Ivo",
		Status:    "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(10, 0).UTC(),
	}

	//expectedSeg := web.Segmentation{
	//	GUID:      "4",
	//	CampaignID:    "8fsa-321nfuv",
	//	Age:       3,
	//	Address: "Sofia",
	//}
	assert.Equal(t, expectedCam, actual)
}

func TestCampaignRepositoryGetReturnQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM users").
		WithArgs("32ff-sad2-fg5").
		WillReturnError(SQLerror{"SQL Error"})

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)
	myDB := persistant.NewUserRepository(db, clock)

	_, err = myDB.Get("32ff-sad2-fg5")
	expectedError := SQLerror{"SQL Error"}

	assert.Equal(t, expectedError, err)
}

//func TestCampaignRepository_UpdateReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("UPDATE campaign").WillReturnError(SQLerror{"ERROR"})
//	mock.ExpectQuery("SELECT \\* FROM campaign").
//		WithArgs("15").
//		WillReturnError(SQLerror{"ERROR"})
//
//	myDB := persistant.NewCampaignRepository(db)
//
//	_ , err = myDB.Update("15", web.RequestCampaign{Segmentation:web.RequestSegmentation{}})
//
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
