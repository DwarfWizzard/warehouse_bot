package repository

import (
	"fmt"

	"github.com/DwarfWizzard/warehouse_bot/pkg/models"
	"github.com/jmoiron/sqlx"
)

type SubsidiaryPostgres struct {
	db *sqlx.DB
}

func NewSubsidiaryPostgres(db *sqlx.DB) *SubsidiaryPostgres {
	return &SubsidiaryPostgres{
		db: db,
	}
}

func (r *SubsidiaryPostgres) GetSubsidiary(cityName string) (models.Subsidiary, error) {
	var subsidiary models.Subsidiary

	query := fmt.Sprintf("SELECT * FROM %s WHERE city=$1", subsidiaryTable)
	err := r.db.Get(&subsidiary, query, cityName)
	if err != nil {
		return subsidiary, fmt.Errorf("repository/GetSubsidiary: [cityName %s] : error %s", cityName, err.Error())
	}

	return subsidiary, nil
} 

func (r *SubsidiaryPostgres) GetSubsidiarys() ([]models.Subsidiary, error) {
	var subsidiarys []models.Subsidiary

	query := fmt.Sprintf("SELECT * FROM %s", subsidiaryTable)
	err := r.db.Select(&subsidiarys, query)
	if err != nil {
		return subsidiarys, fmt.Errorf("repository/GetSubsidiarys: error %s", err.Error())
	}

	return subsidiarys, nil
}