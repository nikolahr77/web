package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func adaptToSegmentation(c segmentationEntity) web.Segmentation {
	return web.Segmentation{
		Address: c.Address,
		Age:     c.Age,
	}
}

func adaptToCampaign(c campaignEntity) web.Campaign {
	return web.Campaign{
		GUID:         c.GUID,
		Name:         c.Name,
		Status:       c.Status,
		Segmentation: adaptToSegmentation(c.Segmentation),
		CreatedOn:    c.CreatedOn,
		UpdatedOn:    c.UpdatedOn,
	}
}

func (c campaignRepository) Get(id string) (web.Campaign, error) {
	var cam campaignEntity
	getCampaign := `
	SELECT * FROM campaign 
	CROSS JOIN segmentation 
	WHERE guid = $1
     `

	rows, err := c.db.Query(getCampaign, id)
	if err != nil {
		return web.Campaign{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&cam.GUID, &cam.Name, &cam.Status, &cam.CreatedOn, &cam.UpdatedOn, &cam.Segmentation.Address, &cam.Segmentation.Age, &cam.Segmentation.CampaignID, &cam.Segmentation.GUID)
		if err != nil {
			return web.Campaign{}, err
		}
	}
	result := adaptToCampaign(cam)
	return result, err
}

func (c campaignRepository) Delete(id string) error {
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

func (c campaignRepository) Update(id string, m web.RequestCampaign) (web.Campaign, error) {
	updateCampaign := `
	UPDATE campaign 
	SET name=$1, status=$2, updated_on=$3
	WHERE guid = $4;`

	updateSegmentation := `
	UPDATE segmentation
	SET address=$1,age = $2
	WHERE campaign_id = $3;`

	updatedOn := time.Now().UTC()

	tx, _ := c.db.Begin()
	_, err := c.db.Exec(updateCampaign, m.Name, "draft", updatedOn, id)
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
		UpdatedOn: updatedOn,
	}, err
}

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
