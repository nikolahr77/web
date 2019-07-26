package persistant_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
)

func TestCreateCampaignRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewCampaignRepository(DB, clock)
	newSeg := web.Segmentation{
		Address: "Sofia 1515",
		Age:     30,
	}
	msgID := uuid.New()

	newCam := web.RequestCampaign{
		Name:         "TestCampaign",
		Segmentation: newSeg,
		MessageGUID:  msgID.String(),
	}

	userID := uuid.New()
	actual, err := mr.Create(newCam, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Campaign{
		GUID: actual.GUID,
		Name: "TestCampaign",
		Segmentation: web.Segmentation{
			GUID:    actual.Segmentation.GUID,
			Address: "Sofia 1515",
			Age:     30,
		},
		Status:    "draft",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,

		MessageGUID: msgID.String(),
		UserID:      userID.String(),
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "campaign")
}

func TestUpdateCampaignRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewCampaignRepository(DB, clock)

	oldSeg := web.Segmentation{
		Address: "Sofia 1515",
		Age:     30,
	}
	msgID := uuid.New()

	oldCam := web.RequestCampaign{
		Name:         "TestCampaign",
		Segmentation: oldSeg,
		MessageGUID:  msgID.String(),
	}

	newSeg := web.Segmentation{
		Address: "Plovdiv 32515",
		Age:     42,
	}

	NewMsgID := uuid.New()

	newCam := web.RequestCampaign{
		Name:         "NeWTestCampaign",
		Segmentation: newSeg,
		MessageGUID:  NewMsgID.String(),
	}
	userID := uuid.New()

	old, err := mr.Create(oldCam, userID.String())

	actual, err := mr.Update(old.GUID, newCam, userID.String())

	if err != nil {
		panic(err)
	}

	expected := web.Campaign{
		GUID: actual.GUID,
		Name: "NeWTestCampaign",
		Segmentation: web.Segmentation{
			GUID:    actual.Segmentation.GUID,
			Address: "Plovdiv 32515",
			Age:     42,
		},
		Status:    "draft",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,

		MessageGUID: NewMsgID.String(),
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "campaign")
}

func TestDeleteCampaignRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewCampaignRepository(DB, clock)

	oldSeg := web.Segmentation{
		Address: "Sofia 1515",
		Age:     30,
	}
	msgID := uuid.New()

	oldCam := web.RequestCampaign{
		Name:         "TestCampaign",
		Segmentation: oldSeg,
		MessageGUID:  msgID.String(),
	}

	userID := uuid.New()

	old, err := mr.Create(oldCam, userID.String())

	err = mr.Delete(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
}

func TestSentStatusCampaignRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewCampaignRepository(DB, clock)

	oldSeg := web.Segmentation{
		Address: "Sofia 1515",
		Age:     30,
	}
	msgID := uuid.New()

	oldCam := web.RequestCampaign{
		Name:         "TestCampaign",
		Segmentation: oldSeg,
		MessageGUID:  msgID.String(),
	}

	userID := uuid.New()

	old, err := mr.Create(oldCam, userID.String())

	actual, err := mr.SentStatus(old.GUID)
	if err != nil {
		panic(err)
	}

	expected := web.Campaign{
		Status: "sent",
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "campaign")
}

func TestGetCampaignRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewCampaignRepository(DB, clock)

	newSeg := web.Segmentation{
		Address: "Plovdiv 32515",
		Age:     42,
	}

	NewMsgID := uuid.New()

	newCam := web.RequestCampaign{
		Name:         "NeWTestCampaign",
		Segmentation: newSeg,
		MessageGUID:  NewMsgID.String(),
	}
	userID := uuid.New()

	campaign, err := mr.Create(newCam, userID.String())

	actual, err := mr.Get(campaign.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Campaign{
		GUID: actual.GUID,
		Name: "NeWTestCampaign",
		Segmentation: web.Segmentation{
			GUID:       actual.Segmentation.GUID,
			Address:    "Plovdiv 32515",
			Age:        42,
			CampaignID: campaign.GUID,
		},
		Status:      "draft",
		CreatedOn:   actual.CreatedOn, //I should't do this
		UpdatedOn:   actual.UpdatedOn,
		UserID:      userID.String(),
		MessageGUID: NewMsgID.String(),
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "campaign")
	DBCleaner(DB, "segmentation")

}

//
//func TestCampaignRepository_Get(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	rowsCam := sqlmock.NewRows([]string{"guid", "name", "status", "created_on", "updated_on"}).
//		AddRow("8fsa-321nfuv", "New Campaign", "draft", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())
//
//	rowsSeg := sqlmock.NewRows([]string{"address", "age", "campaign_id", "id"}).
//		AddRow("Sofia", 3, "8fsa-321nfuv", "4")
//
//	mock.ExpectQuery("SELECT \\* FROM users CROSS JOIN segmentation").
//		WithArgs("8fsa-321nfuv").
//		WillReturnRows(rowsCam, rowsSeg)
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//	myDB := persistant.NewUserRepository(db, clock)
//
//	actual, err := myDB.Get("15")
//
//	expectedCam := web.Campaign{
//		GUID:      "8fsa-321nfuv",
//		Name:      "Ivo",
//		Status:    "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK",
//		CreatedOn: time.Unix(10, 0).UTC(),
//		UpdatedOn: time.Unix(10, 0).UTC(),
//	}
//
//	//expectedSeg := web.Segmentation{
//	//	GUID:      "4",
//	//	CampaignID:    "8fsa-321nfuv",
//	//	Age:       3,
//	//	Address: "Sofia",
//	//}
//	assert.Equal(t, expectedCam, actual)
//}
//
//func TestCampaignRepositoryGetReturnQueryError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectQuery("SELECT \\* FROM users").
//		WithArgs("32ff-sad2-fg5").
//		WillReturnError(SQLerror{"SQL Error"})
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//	myDB := persistant.NewUserRepository(db, clock)
//
//	_, err = myDB.Get("32ff-sad2-fg5")
//	expectedError := SQLerror{"SQL Error"}
//
//	assert.Equal(t, expectedError, err)
//}
//
////func TestCampaignRepository_UpdateReturnError(t *testing.T) {
////	db, mock, err := sqlmock.New()
////	if err != nil {
////		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
////	}
////	defer db.Close()
////
////	mock.ExpectExec("UPDATE campaign").WillReturnError(SQLerror{"ERROR"})
////	mock.ExpectQuery("SELECT \\* FROM campaign").
////		WithArgs("15").
////		WillReturnError(SQLerror{"ERROR"})
////
////	myDB := persistant.NewCampaignRepository(db)
////
////	_ , err = myDB.Update("15", web.RequestCampaign{Segmentation:web.RequestSegmentation{}})
////
////	expected := SQLerror{"ERROR"}
////
////	assert.Equal(t, expected, err)
////}
