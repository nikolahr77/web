package persistant_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestCreateUpdateGetCampaignRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}
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
	if err != nil {
		panic(err)
	}

	_, err = mr.Update(old.GUID, newCam, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := mr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

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
			CampaignID: old.GUID,
		},
		Status:      "draft",
		CreatedOn:   time.Unix(25000, 0).UTC(),
		UpdatedOn:   time.Unix(25000, 0).UTC(),
		UserID:      userID.String(),
		MessageGUID: NewMsgID.String(),
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "campaign")
	dbCleaner(DB, "segmentation")
}

func TestCreateDeleteGetCampaignRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}
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

	actual, err := mr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, web.Campaign{})

}

func TestSentStatusCampaignRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

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

	_, err = mr.SentStatus(old.GUID)
	if err != nil {
		panic(err)
	}

	actual, err := mr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}
	expected := web.Campaign{
		Status: "sent",
	}

	assert.Equal(t, expected.Status, actual.Status)

	dbCleaner(DB, "campaign")
	dbCleaner(DB, "segmentation")
}
