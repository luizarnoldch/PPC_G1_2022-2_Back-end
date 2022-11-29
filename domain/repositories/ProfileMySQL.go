package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type ProfileRepository interface {
	FindAllProfiles() ([]entities.Profile, error)
	FinProfileById(id int64) (*entities.Profile, error)
	SaveProfile(req dto.ProfileRequest) (*entities.Profile, error)
	UpdateProfile(id int64, req dto.ProfileRequest) (*entities.Profile, error)
	DeleteProfile(id int64) (*entities.Profile, error)
}

type ProfileDatabaseMySQL struct {
	client *sqlx.DB
}

func NewProfileDataMySQL(db *sqlx.DB) ProfileDatabaseMySQL {
	return ProfileDatabaseMySQL{db}
}

func (db ProfileDatabaseMySQL) FindAllProfiles() ([]entities.Profile, error) {
	profiles := make([]entities.Profile, 0)
	profileQuery := "SELECT * FROM posta_ppc.perfil"
	err := db.client.Select(&profiles, profileQuery)
	if err != nil {
		return nil, errors.New("error while getting all areas")
	}
	return profiles, nil
}

func (db ProfileDatabaseMySQL) FinProfileById(id int64) (*entities.Profile, error) {
	var profile entities.Profile
	profileQuery := "SELECT * FROM posta_ppc.area WHERE ID_Area = ?"
	err := db.client.Get(&profile, profileQuery, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("no profile found")
	}
	if err != nil {
		return nil, errors.New("unexpected database error")
	}
	return &profile, nil
}

func (db ProfileDatabaseMySQL) SaveProfile(req dto.ProfileRequest) (*entities.Profile, error) {
	areaQuery := `INSERT INTO
		posta_ppc.perfil (Nombre_Perfil,
		                Estado_Perfil,
		                Atributos_Perfil)
		VALUES(?,?,?)`
	res, err := db.client.Exec(
		areaQuery,
		req.Nombre_Perfil,
		req.Estado_Perfil,
		req.Atributos_Perfil)
	if err != nil {
		return nil, errors.New("error while saving profile")
	}
	id, errId := res.LastInsertId()
	if errId != nil {
		return nil, errors.New("error while getting id from profile")
	}
	profile := entities.Profile{
		id,
		req.Nombre_Perfil,
		req.Estado_Perfil,
		req.Atributos_Perfil}
	return &profile, nil
}

func (db ProfileDatabaseMySQL) UpdateProfile(id int64, req dto.ProfileRequest) (*entities.Profile, error) {
	profileQuery := `UPDATE posta_ppc.perfil 
		SET Nombre_Perfil = ?,
		    Estado_Perfil = ?,
		    Atributos_Perfil = ?
		WHERE ID_Perfil = ?`

	profileUpdate, err := db.client.Exec(
		profileQuery,
		req.Nombre_Perfil,
		req.Estado_Perfil,
		req.Atributos_Perfil,
		id)
	if err != nil {
		return nil, errors.New("error while updating profile")
	}
	rows, errNoRow := profileUpdate.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting updated rows")
	}
	if rows == 0 {
		return nil, errors.New("no area updated")
	}
	profileID, errID := profileUpdate.LastInsertId()
	if errID != nil {
		return nil, errors.New("error while getting id from profile")
	}
	profile := entities.NewProfile(
		profileID,
		req.Nombre_Perfil,
		req.Estado_Perfil,
		req.Atributos_Perfil)
	return &profile, nil
}

func (db ProfileDatabaseMySQL) DeleteProfile(id int64) (*entities.Profile, error) {
	var profile entities.Profile
	profileQuery := "DELETE FROM posta_ppc.perfil WHERE ID_Perfil = ?"
	profileDelete, errDelete := db.client.Exec(profileQuery, id)
	if errDelete != nil {
		return nil, errors.New("error while deleting the profile")
	}
	rows, errNoRow := profileDelete.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting deleted rows")
	}
	if rows == 0 {
		return nil, errors.New("no profile deleted")
	}
	profile.ID_Perfil = id
	return &profile, nil
}
