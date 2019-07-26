package persistant

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/web"
	"github.com/web/convert"
	"time"
)

//Get is used to return a campaign from the DB by a given ID.
func (c campaignRepository) Get(id string, userID string) (web.Campaign, error) {
	var cam campaignEntity
	getCampaign := `
	SELECT * FROM campaign 
	CROSS JOIN segmentation 
	WHERE guid = $1 AND userid = $2
     `

	rows, err := c.db.Query(getCampaign, id, userID)
	if err != nil {
		return web.Campaign{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cam.GUID, &cam.Name, &cam.Status, &cam.CreatedOn, &cam.UpdatedOn, &cam.MessageGUID, &cam.UserID,
			&cam.Segmentation.Address, &cam.Segmentation.Age, &cam.Segmentation.CampaignID,
			&cam.Segmentation.GUID)
		if err != nil {
			return web.Campaign{}, err
		}
	}
	result := web.Campaign{}
	convert.SourceToDestination(cam, &result)
	return result, err
}

//Delete is used to remove a campaign from the DB by a given ID.
func (c campaignRepository) Delete(id string, userID string) error {
	deleteCampaign := `
	DELETE FROM campaign WHERE guid=$1`

	deleteSegmentation := `
	DELETE FROM segmentation where campaign_id=$1`

	tx, _ := c.db.Begin()
	_, err := c.db.Exec(deleteSegmentation, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = c.db.Exec(deleteCampaign, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

//Update searches the DB for a campaign by a given
// ID and updates the campaign with the given RequestCampaign
func (c campaignRepository) Update(id string, m web.RequestCampaign, userID string) (web.Campaign, error) {
	if m.Status != "draft" {
		return web.Campaign{}, errors.New("Sent or delivered campaign can't be edited")
	}
	updateCampaign := `
	UPDATE campaign 
	SET name=$1, status=$2, updated_on=$3, message_guid=$4
	WHERE guid = $5 AND userID = $6;`

	updateSegmentation := `
	UPDATE segmentation
	SET address=$1,age = $2
	WHERE campaign_id = $3;`

	updatedOn := c.clock.Now().UTC()

	tx, _ := c.db.Begin()
	_, err := c.db.Exec(updateCampaign, m.Name, "draft", updatedOn, m.MessageGUID, id, userID)
	if err != nil {
		tx.Rollback()
		return web.Campaign{}, err
	}
	_, err = c.db.Exec(updateSegmentation, m.Segmentation.Address, m.Segmentation.Age, id)
	if err != nil {
		tx.Rollback()
		return web.Campaign{}, err
	}
	tx.Commit()

	return web.Campaign{
		GUID:   id,
		Name:   m.Name,
		Status: m.Status,
		Segmentation: web.Segmentation{
			Address: m.Segmentation.Address,
			Age:     m.Segmentation.Age,
		},
		UpdatedOn:   updatedOn,
		MessageGUID: m.MessageGUID,
	}, err
}

//SentStatus changes the satatus of the campaign to sent
func (c campaignRepository) SentStatus(id string) (web.Campaign, error) {
	updateStatus := `UPDATE campaign 
	SET status=$1
	WHERE guid = $2
	`
	_, err := c.db.Exec(updateStatus, "sent", id)
	if err != nil {
		return web.Campaign{}, err
	}
	return web.Campaign{
		Status: "sent",
	}, err
}

//Create adds a new campaign to the DB
func (c campaignRepository) Create(m web.RequestCampaign, userID string) (web.Campaign, error) {
	uuid := uuid.New()
	inseretCampaign := `
	INSERT INTO campaign (guid, name, status, created_on, updated_on, message_guid, userID)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	inseretSegmentation := `
	INSERT INTO segmentation (address, age, campaign_id)
	VALUES ($1, $2, $3);`

	createdOn := c.clock.Now().UTC()
	tx, _ := c.db.Begin()
	_, err := tx.Exec(inseretCampaign, uuid, m.Name, "draft", createdOn, createdOn, m.MessageGUID, userID)
	if err != nil {
		tx.Rollback()
		return web.Campaign{}, err
	}
	_, err = tx.Exec(inseretSegmentation, m.Segmentation.Address, m.Segmentation.Age, uuid)
	if err != nil {
		tx.Rollback()
		return web.Campaign{}, err
	}

	tx.Commit()

	return web.Campaign{
		GUID:   uuid.String(),
		Name:   m.Name,
		Status: "draft",
		Segmentation: web.Segmentation{
			Address: m.Segmentation.Address,
			Age:     m.Segmentation.Age,
		},
		CreatedOn:   createdOn,
		UpdatedOn:   createdOn,
		MessageGUID: m.MessageGUID,
		UserID:      userID,
	}, err
}

type campaignEntity struct {
	GUID         string `db:"uuid"`
	Name         string `db:"name"`
	Segmentation segmentationEntity
	Status       string    `db": "status"`
	CreatedOn    time.Time `db: "created_on"`
	UpdatedOn    time.Time `db: "updated_on"`
	MessageGUID  string    `db: "message_guid"`
	UserID       string    `db: "userID"`
}

type segmentationEntity struct {
	GUID       string `db:"uuid"`
	Address    string `db:"address"`
	Age        int    `db:"age"`
	CampaignID string `db:"campaign_id"`
}

func NewCampaignRepository(db *sql.DB, clock Clock) web.CampaignRepository {
	return campaignRepository{db: db, clock: clock}
}

type campaignRepository struct {
	db    *sql.DB
	clock Clock
}
