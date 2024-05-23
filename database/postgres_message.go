package database

import (
	"context"
	"database/sql"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/ins_log"
)

//aqui guardaremos los strings de las consultas

const (
	psqlInsertMessage = `INSERT INTO message 
		(type, content, mobile_number, mobile_countryisocode, shortnumber, telco, created, routingtype, 
		matchedpattern, serviceid, telcoid, sessionaction, sessionparameters_map, sessiontimeoutseconds, 
		priority, clientid, url, accesstimeoutseconds, request_id, defaultaction_id, application_id, 
		session_id, processed, millissincerequest, sessionapplicationname, sendafter, sendbefore, sent, 
		status, accesstimeouthandlerqueuename, useunsupportedmobilesregistry, originname) 
	VALUES 
		($?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?, $?)
		RETURNING id
	`
)

// mysql usado para trabajar con postgres y meesage
type PostgresMessage struct {
	db *sql.DB
}

func NewPostgresMessage(db *sql.DB) *PostgresMessage {
	return &PostgresMessage{db}
}
func (p *PostgresMessage) InsertMessage(m *models.MessageModel, ctx context.Context) error {
	ins_log.Tracef(ctx, "se tratara de insertar el mensaje en la base de datos")
	err := p.db.QueryRow(psqlInsertMessage,
		m.Type, m.Content, m.MobileNumber, m.MobileCountryISOCode, m.ShortNumber, m.Telco, m.Created, m.RoutingType,
		m.MatchedPattern, m.ServiceID, m.TelcoID, m.SessionAction, m.SessionParametersMap, m.SessionTimeoutSeconds,
		m.Priority, m.ClientID, m.URL, m.AccessTimeoutSeconds, m.RequestID, m.DefaultActionID, m.ApplicationID, m.SessionID,
		m.Processed, m.MillisSinceRequest, m.SessionApplicationName, m.Sendafter, m.Sendbefore, m.Sent,
		m.Status, m.AccessTimeoutHandlerQueuename, m.UseUnsupportedMobilesRegistry, m.OriginName,
	).Scan(&m.Id)
	if err != nil {
		return err
	}
	ins_log.Infof(ctx, "se inserto el mensaje de manera correcta el mensaje se guardo con el id %s", m.Id)
	return nil
}
