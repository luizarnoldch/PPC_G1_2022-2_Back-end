package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type AreaRepository interface {
	FindAllAreas() ([]entities.Area, error)
	FinAreaById(id int64) (*entities.Area, error)
	SaveArea(req dto.AreaRequest) (*entities.Area, error)
	UpdateArea(id int64, req dto.AreaRequest) (*entities.Area, error)
	DeleteArea(id int64) (*entities.Area, error)
}

type AreaDatabaseMySQL struct {
	client *sqlx.DB
}

func NewAreaDataMySQL(db *sqlx.DB) AreaDatabaseMySQL {
	return AreaDatabaseMySQL{db}
}

func (db AreaDatabaseMySQL) FindAllAreas() ([]entities.Area, error) {
	areas := make([]entities.Area, 0)
	areaQuery := "SELECT * FROM posta_ppc.area"
	err := db.client.Select(&areas, areaQuery)
	if err != nil {
		return nil, errors.New("error while getting all areas")
	}
	return areas, nil
}

func (db AreaDatabaseMySQL) FinAreaById(id int64) (*entities.Area, error) {
	var area entities.Area
	areaQuery := "SELECT * FROM posta_ppc.area WHERE ID_Area = ?"
	err := db.client.Get(&area, areaQuery, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("no area found")
	}
	if err != nil {
		return nil, errors.New("unexpected database error")
	}
	return &area, nil
}

func (db AreaDatabaseMySQL) SaveArea(req dto.AreaRequest) (*entities.Area, error) {
	areaQuery := `INSERT INTO
		posta_ppc.area (Nombre_Area,
		                Ubicacion_Area,
		                Localidades_Area,
		                Ciudad_Area,
		                Pais_Area,
		                Capacidad_Area,
		                Tag_Area,
		                Desc_Area)
		VALUES(?,?,?,?,?,?,?,?)
	`
	res, err := db.client.Exec(
		areaQuery,
		req.AreaName,
		req.AreaAddress,
		req.AreaLocation,
		req.AreaCity,
		req.AreaCountry,
		req.AreaCapacity,
		req.AreaTag,
		req.AreaDiscount)

	if err != nil {
		return nil, errors.New("error while saving area")
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		return nil, errors.New("error while getting id from patient")
	}

	area := entities.Area{
		id,
		req.AreaName,
		req.AreaAddress,
		req.AreaLocation,
		req.AreaCity,
		req.AreaCountry,
		req.AreaCapacity,
		req.AreaTag,
		req.AreaDiscount}

	return &area, nil
}

func (db AreaDatabaseMySQL) UpdateArea(id int64, req dto.AreaRequest) (*entities.Area, error) {
	areaQuery := `UPDATE posta_ppc.area 
		SET Nombre_Area = ?,
		    Ubicacion_Area = ?,
		    Localidades_Area = ?,
		    Ciudad_Area = ?,
		    Pais_Area = ?,
		    Capacidad_Area = ?,
		    Tag_Area = ?,
		    Desc_Area = ?
		WHERE ID_Area = ?`

	areaUpdate, err := db.client.Exec(
		areaQuery,
		req.AreaName,
		req.AreaAddress,
		req.AreaLocation,
		req.AreaCity,
		req.AreaCountry,
		req.AreaCapacity,
		req.AreaTag,
		req.AreaDiscount,
		id)
	if err != nil {
		return nil, errors.New("error while area")
	}
	rows, errNoRow := areaUpdate.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting updated rows")
	}
	if rows == 0 {
		return nil, errors.New("no area updated")
	}
	areaID, errID := areaUpdate.LastInsertId()
	if errID != nil {
		return nil, errors.New("error while getting id from area")
	}
	area := entities.NewArea(
		areaID,
		req.AreaName,
		req.AreaAddress,
		req.AreaLocation,
		req.AreaCity,
		req.AreaCountry,
		req.AreaCapacity,
		req.AreaTag,
		req.AreaDiscount)
	return &area, nil
}

func (db AreaDatabaseMySQL) DeleteArea(id int64) (*entities.Area, error) {
	var area entities.Area
	areaQuery := "DELETE FROM posta_ppc.area WHERE ID_Area = ?"
	areaDelete, errDelete := db.client.Exec(areaQuery, id)
	if errDelete != nil {
		return nil, errors.New("error while deleting the area")
	}
	rows, errNoRow := areaDelete.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting deleted rows")
	}
	if rows == 0 {
		return nil, errors.New("no area deleted")
	}
	area.AreaID = id
	return &area, nil
}
