package database

import (
	"database/sql"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/ins_log"
)

//aqui guardaremos los strings de las consultas

const (
	mySQLInsertMessage = `INSERT INTO message ` +
		`(type, content, mobile_number, mobile_countryisocode, shortnumber, telco, created, routingtype, ` +
		`matchedpattern, serviceid, telcoid, sessionaction, sessionparameters_map, sessiontimeoutseconds, ` +
		`priority, clientid, url, accesstimeoutseconds, request_id, defaultaction_id, application_id, ` +
		`session_id, processed, millissincerequest, sessionapplicationname, sendafter, sendbefore, sent, ` +
		`status, accesstimeouthandlerqueuename, useunsupportedmobilesregistry, originname) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	mySQLGetMessageById = `select content from message where id = ?`
)

// mysql usado para trabajar con postgres y meesage
type MysqlMessage struct {
	db *sql.DB
}

func NewMysqlMessage(db *sql.DB) *MysqlMessage {
	return &MysqlMessage{db}
}
func (p *MysqlMessage) InsertMessage(m *models.MessageModel) error {
	ins_log.Tracef(ctx, "se tratara de insertar el mensaje en la base de datos")
	stmt, err := p.db.Prepare(mySQLInsertMessage)
	ins_log.Tracef(ctx, "esta es la consulta que intentaremos insertar: %s", mySQLInsertMessage)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		(m.Type),
		(m.Content),
		StringToNull(m.MobileNumber),
		StringToNull(m.MobileCountryISOCode),
		StringToNull(m.ShortNumber),
		StringToNull(m.Telco),
		TimeToNull(m.Created),
		StringToNull(m.RoutingType),
		StringToNull(m.MatchedPattern),
		StringToNull(m.ServiceID),
		StringToNull(m.TelcoID),
		StringToNull(m.SessionAction),
		StringToNull(m.SessionParametersMap),
		Uint64ToNull(m.SessionTimeoutSeconds),
		Uint64ToNull(m.Priority),
		StringToNull(m.ClientID),
		StringToNull(m.URL),
		Uint64ToNull(m.AccessTimeoutSeconds),
		Uint64ToNull(m.RequestID),
		Uint64ToNull(m.DefaultActionID),
		Uint64ToNull(m.ApplicationID),
		Uint64ToNull(m.SessionID),
		TimeToNull(m.Processed),
		Uint64ToNull(m.MillisSinceRequest),
		StringToNull(m.SessionApplicationName),
		TimeToNull(m.Sendafter),
		TimeToNull(m.Sendbefore),
		TimeToNull(m.Sent),
		StringToNull(m.Status),
		StringToNull(m.AccessTimeoutHandlerQueuename),
		Uint64ToNull(m.UseUnsupportedMobilesRegistry),
		StringToNull(m.OriginName),
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = uint64(id)
	ins_log.Infof(ctx, "se inserto el mensaje de manera correcta el mensaje se guardo con el id %d ", m.Id)
	return nil
}

func (p *MysqlMessage) GetByID(id uint64) (*models.MessageModel, error) {
	stmt, err := p.db.Prepare(mySQLGetMessageById)
	if err != nil {
		return &models.MessageModel{}, err
	}
	defer stmt.Close()

	return ScanRowMessage(stmt.QueryRow(id))
}
