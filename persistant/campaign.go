package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func (c campaignRepository)Create(m web.RequestCampaign,s web.RequestSegmentation) (web.Campaign, web.Segmentation, error){
   uuidCam := uuid.New()
   uuidSeg := uuid.New()
   query := `
	INSERT INTO campaign (guid, name, status, created_on, updated_on
	VALUES ($1, $2, $3, $4, $5);`
   createdOn := time.Now().UTC()
   _, err := c.db.Exec(query,uuidCam, m.Name, m.Status,createdOn,createdOn)
   if err != nil {
   	return web.Campaign{}, web.Segmentation{}, err
   }
   query1 := `
	INSERT INTO segmentation (guid, address, age, campaign_id)
	VALUES ($1, $2, $3, $4);`
	_, err1 := c.db.Exec(query1,uuidSeg,s.Address,s.Age,uuidCam)
	if err1 != nil {
		return web.Campaign{}, web.Segmentation{}, err
	}
	return web.Campaign{
		GUID: uuidCam.String(),
		Name: m.Name,
		Status:  m.Status,
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
	},
	web.Segmentation{
		GUID: uuidSeg.String(),
		Address: s.Address,
		Age: s.Age,
		CampaignID: uuidCam.String(),
	}, err
}



type campaignEntity struct {
	GUID         string `db:"uuid"`
	Name         string `db:"name"`
	Segmentation segmentationEntity
	Status       string `db": "status"`
	CreatedOn time.Time `db: "created_on"`
	UpdatedOn time.Time `db: "updated_on"`
}

type segmentationEntity struct {
	GUID string `db:"uuid"`
	Address string `db:"address"`
	Age     int `db:"age"`
	CampaignID string `db:"campaign_id"`
}

func NewCampaignRepository(db *sql.DB) web.CampaignRepository{
	return campaignRepository{db:db}
}

type campaignRepository struct {
	db *sql.DB
}