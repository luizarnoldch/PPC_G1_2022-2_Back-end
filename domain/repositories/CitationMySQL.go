package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type CitationRepository interface {
	FindAllCitation() ([]entities.Citation, error)
	FinCitationById(id int64) (*entities.Citation, error)
	SaveCitation(req dto.CitationRequest) (*entities.Citation, error)
	UpdateCitation(id int64, req dto.CitationRequest) (*entities.Citation, error)
	DeleteCitation(id int64) (*entities.Citation, error)
}

type CitationDatabaseMySQL struct {
	client *sqlx.DB
}

func NewCitationDataMySQL(db *sqlx.DB) CitationDatabaseMySQL {
	return CitationDatabaseMySQL{db}
}

func (db CitationDatabaseMySQL) FindAllCitation() ([]entities.Citation, error) {
	citations := make([]entities.Citation, 0)
	citationQuery := "SELECT * FROM posta_ppc.cita"
	err := db.client.Select(&citations, citationQuery)
	if err != nil {
		return nil, errors.New("error while getting all citations")
	}
	return citations, nil
}

func (db CitationDatabaseMySQL) FinCitationById(id int64) (*entities.Citation, error) {
	var citation entities.Citation
	citationQuery := "SELECT * FROM posta_ppc.cita WHERE ID_Cita = ?"
	err := db.client.Get(&citation, citationQuery, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("no citation found")
	}
	if err != nil {
		return nil, errors.New("unexpected database error")
	}
	return &citation, nil
}

func (db CitationDatabaseMySQL) SaveCitation(req dto.CitationRequest) (*entities.Citation, error) {
	citationQuery := `INSERT INTO
		posta_ppc.cita (ID_Usuario,
		                ID_Paciente,
		                Titulo_Cita,
		                DATE_Inicio_Cita,
		                DATE_Fin_Cita,
		                Hora_Inicio_Cita,
		                Hora_Fin_Cita,
		                Descripcion_Cita,
		                Identificador_Cita,
		                Hora_Llegada_Paciente_Cita,
		                Hora_Salida_Paciente_cita,
		                Motivo_Anulacion_Cita,
		                Estado_Cita)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`
	res, err := db.client.Exec(
		citationQuery,
		req.ID_Usuario,
		req.ID_Paciente,
		req.Titulo_Cita,
		req.DATE_Inicio_Cita,
		req.DATE_Fin_Cita,
		req.Hora_Inicio_Cita,
		req.Hora_Fin_Cita,
		req.Descripcion_Cita,
		req.Identificador_Cita,
		req.Hora_Llegada_Paciente_Cita,
		req.Hora_Salida_Paciente_cita,
		req.Motivo_Anulacion_Cita,
		req.Estado_Cita)

	if err != nil {
		return nil, errors.New("error while saving citation")
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		return nil, errors.New("error while getting id from citation")
	}

	citation := entities.Citation{
		id,
		req.ID_Usuario,
		req.ID_Paciente,
		req.Titulo_Cita,
		req.DATE_Inicio_Cita,
		req.DATE_Fin_Cita,
		req.Hora_Inicio_Cita,
		req.Hora_Fin_Cita,
		req.Descripcion_Cita,
		req.Identificador_Cita,
		req.Hora_Llegada_Paciente_Cita,
		req.Hora_Salida_Paciente_cita,
		req.Motivo_Anulacion_Cita,
		req.Estado_Cita}
	return &citation, nil
}

func (db CitationDatabaseMySQL) UpdateCitation(id int64, req dto.CitationRequest) (*entities.Citation, error) {
	citationQuery := `UPDATE posta_ppc.cita 
		SET ID_Usuario = ?,
		    ID_Paciente = ?,
		    Titulo_Cita = ?,
		    DATE_Inicio_Cita = ?,
		    DATE_Fin_Cita = ?,
		    Hora_Inicio_Cita = ?,
		    Hora_Fin_Cita = ?,
		    Descripcion_Cita = ?,
		    Identificador_Cita = ?,
		    Hora_Llegada_Paciente_Cita = ?,
		    Hora_Salida_Paciente_cita = ?,
		    Motivo_Anulacion_Cita = ?,
		    Estado_Cita = ?
		WHERE ID_Cita = ?`

	citationUpdate, err := db.client.Exec(
		citationQuery,
		req.ID_Usuario,
		req.ID_Paciente,
		req.Titulo_Cita,
		req.DATE_Inicio_Cita,
		req.DATE_Fin_Cita,
		req.Hora_Inicio_Cita,
		req.Hora_Fin_Cita,
		req.Descripcion_Cita,
		req.Identificador_Cita,
		req.Hora_Llegada_Paciente_Cita,
		req.Hora_Salida_Paciente_cita,
		req.Motivo_Anulacion_Cita,
		req.Estado_Cita,
		id)
	if err != nil {
		return nil, errors.New("error while updating citation")
	}
	rows, errNoRow := citationUpdate.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting updated rows")
	}
	if rows == 0 {
		return nil, errors.New("no area updated")
	}
	citationID, errID := citationUpdate.LastInsertId()
	if errID != nil {
		return nil, errors.New("error while getting id from citation")
	}
	citation := entities.NewCitation(
		citationID,
		req.ID_Usuario,
		req.ID_Paciente,
		req.Titulo_Cita,
		req.DATE_Inicio_Cita,
		req.DATE_Fin_Cita,
		req.Hora_Inicio_Cita,
		req.Hora_Fin_Cita,
		req.Descripcion_Cita,
		req.Identificador_Cita,
		req.Hora_Llegada_Paciente_Cita,
		req.Hora_Salida_Paciente_cita,
		req.Motivo_Anulacion_Cita,
		req.Estado_Cita)
	return &citation, nil
}

func (db CitationDatabaseMySQL) DeleteCitation(id int64) (*entities.Citation, error) {
	var citation entities.Citation
	citationQuery := "DELETE FROM posta_ppc.cita WHERE ID_Cita = ?"
	ctitationDelete, errDelete := db.client.Exec(citationQuery, id)
	if errDelete != nil {
		return nil, errors.New("error while deleting the citation")
	}
	rows, errNoRow := ctitationDelete.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting deleted rows")
	}
	if rows == 0 {
		return nil, errors.New("no citation deleted")
	}
	citation.ID_Cita = id
	return &citation, nil
}
