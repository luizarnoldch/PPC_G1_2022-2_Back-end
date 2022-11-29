package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type UserRepository interface {
	FindAllUser() ([]entities.User, error)
	FinUserById(id int64) (*entities.User, error)
	SaveUser(req dto.UserRequest) (*entities.User, error)
	UpdateUser(id int64, req dto.UserRequest) (*entities.User, error)
	DeleteUser(id int64) (*entities.User, error)
}

type UserDatabaseMySQL struct {
	client *sqlx.DB
}

func NewUserDataMySQL(db *sqlx.DB) AreaDatabaseMySQL {
	return AreaDatabaseMySQL{db}
}

func (db AreaDatabaseMySQL) FindAllUser() ([]entities.User, error) {
	users := make([]entities.User, 0)
	areaQuery := "SELECT * FROM posta_ppc.usuario"
	err := db.client.Select(&users, areaQuery)
	if err != nil {
		return nil, errors.New("error while getting all users")
	}
	return users, nil
}

func (db AreaDatabaseMySQL) FinUserById(id int64) (*entities.User, error) {
	var user entities.User
	userQuery := "SELECT * FROM posta_ppc.usuario WHERE ID_Usuario = ?"
	err := db.client.Get(&user, userQuery, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user found")
	}
	if err != nil {
		return nil, errors.New("unexpected database error")
	}
	return &user, nil
}

func (db AreaDatabaseMySQL) SaveUser(req dto.UserRequest) (*entities.User, error) {
	areaQuery := `INSERT INTO
		posta_ppc.usuario(ID_Area,
		                ID_Perfil,
		                Nombre_Usuario,
		                Apellido_Usuairo,
		                Nick_Usuario,
		                Clave_Usuario,
		                Mail_Usuario,
		                Especialidad_Usuario,
		                Estado_Usuario)
		VALUES(?,?,?,?,?,?,?,?,?)`
	res, err := db.client.Exec(
		areaQuery,
		req.ID_Usuario,
		req.ID_Area,
		req.ID_Perfil,
		req.Nombre_Usuario,
		req.Apellido_Usuairo,
		req.Nick_Usuario,
		req.Clave_Usuario,
		req.Mail_Usuario,
		req.Especialidad_Usuario,
		req.Estado_Usuario)

	if err != nil {
		return nil, errors.New("error while saving user")
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		return nil, errors.New("error while getting id from user")
	}

	user := entities.User{
		id,
		req.ID_Area,
		req.ID_Perfil,
		req.Nombre_Usuario,
		req.Apellido_Usuairo,
		req.Nick_Usuario,
		req.Clave_Usuario,
		req.Mail_Usuario,
		req.Especialidad_Usuario,
		req.Estado_Usuario}

	return &user, nil
}

func (db AreaDatabaseMySQL) UpdateUser(id int64, req dto.UserRequest) (*entities.User, error) {
	userQuery := `UPDATE posta_ppc.usuario 
		SET ID_Area = ?,
		    ID_Perfil = ?,
		    Nombre_Usuario = ?,
		    Apellido_Usuairo = ?,
		    Nick_Usuario = ?,
		    Clave_Usuario = ?,
		    Mail_Usuario = ?,
		    Especialidad_Usuario = ?,
		    Estado_Usuario = ?
		WHERE ID_Usuario = ?`

	userUpdate, err := db.client.Exec(
		userQuery,
		req.ID_Area,
		req.ID_Perfil,
		req.Nombre_Usuario,
		req.Apellido_Usuairo,
		req.Nick_Usuario,
		req.Clave_Usuario,
		req.Mail_Usuario,
		req.Especialidad_Usuario,
		req.Estado_Usuario,
		id)
	if err != nil {
		return nil, errors.New("error while updating user")
	}
	rows, errNoRow := userUpdate.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting updated rows")
	}
	if rows == 0 {
		return nil, errors.New("no user updated")
	}
	userID, errID := userUpdate.LastInsertId()
	if errID != nil {
		return nil, errors.New("error while getting id from user")
	}
	user := entities.NewUser(
		userID,
		req.ID_Area,
		req.ID_Perfil,
		req.Nombre_Usuario,
		req.Apellido_Usuairo,
		req.Nick_Usuario,
		req.Clave_Usuario,
		req.Mail_Usuario,
		req.Especialidad_Usuario,
		req.Estado_Usuario)
	return &user, nil
}

func (db AreaDatabaseMySQL) DeleteUser(id int64) (*entities.User, error) {
	var user entities.User
	userQuery := "DELETE FROM posta_ppc.usuario WHERE ID_Usuario = ?"
	userDelete, errDelete := db.client.Exec(userQuery, id)
	if errDelete != nil {
		return nil, errors.New("error while deleting the user")
	}
	rows, errNoRow := userDelete.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting deleted rows")
	}
	if rows == 0 {
		return nil, errors.New("no user deleted")
	}
	user.ID_Usuario = id
	return &user, nil
}
