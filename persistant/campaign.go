package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func (c campaignRepository) Create(m web.RequestCampaign) (web.Campaign, error) {
	uuid := uuid.New()
	inseretCampaign := `
	INSERT INTO campaign (guid, name, status, created_on, updated_on)
	VALUES ($1, $2, $3, $4, $5);`

	inseretSegmentation := `
	INSERT INTO segmentation (address, age, campaign_id)
	VALUES ($1, $2, $3);`

	createdOn := time.Now().UTC()
	tx, _ := c.db.Begin()
	_, err := tx.Exec(inseretCampaign, uuid, m.Name, "draft", createdOn, createdOn)
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
		Status: m.Status,
		Segmentation: web.Segmentation{
			Address: m.Segmentation.Address,
			Age:     m.Segmentation.Age,
		},
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
	}, err
}

type campaignEntity struct {
	GUID         string `db:"uuid"`
	Name         string `db:"name"`
	Segmentation segmentationEntity
	Status       string    `db": "status"`
	CreatedOn    time.Time `db: "created_on"`
	UpdatedOn    time.Time `db: "updated_on"`
}

type segmentationEntity struct {
	GUID       string `db:"uuid"`
	Address    string `db:"address"`
	Age        int    `db:"age"`
	CampaignID string `db:"campaign_id"`
}

func NewCampaignRepository(db *sql.DB) web.CampaignRepository {
	return campaignRepository{db: db}
}

type campaignRepository struct {
	db *sql.DB
}
