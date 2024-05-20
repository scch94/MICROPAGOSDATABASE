package models

import (
	"fmt"
	"time"
)

// modelo de la base de message ademas del json
type MessageModel struct {
	Utfi                          string    `json:"utfi"`
	Id                            uint64    `json:"id"`
	Type                          string    `json:"type"` // Nombre del campo en MySQL
	Content                       string    `json:"content"`
	MobileNumber                  string    `json:"mobile_number"`
	MobileCountryISOCode          string    `json:"mobile_country_iso_code"`
	ShortNumber                   string    `json:"short_number"`
	Telco                         string    `json:"telco"`
	Created                       time.Time `json:"created"`
	RoutingType                   string    `json:"routing_type"`
	MatchedPattern                string    `json:"matched_pattern"`
	ServiceID                     string    `json:"service_id"`
	TelcoID                       string    `json:"telco_id"`
	SessionAction                 string    `json:"session_action"`
	SessionParametersMap          string    `json:"session_parameters_map"`
	SessionTimeoutSeconds         uint64    `json:"session_timeout_seconds"`
	Priority                      uint64    `json:"priority"`
	ClientID                      string    `json:"client_id"`
	URL                           string    `json:"url"`
	AccessTimeoutSeconds          uint64    `json:"access_timeout_seconds"`
	RequestID                     uint64    `json:"request_id"`
	DefaultActionID               uint64    `json:"default_action_id"`
	ApplicationID                 uint64    `json:"application_id"`
	SessionID                     uint64    `json:"session_id"`
	Processed                     time.Time `json:"processed"`
	MillisSinceRequest            uint64    `json:"millis_since_request"`
	SessionApplicationName        string    `json:"session_application_name"`
	Sendafter                     string    `json:"sendafter"`
	Sendbefore                    string    `json:"sendbefore"`
	Sent                          time.Time `json:"sent"`
	Status                        string    `json:"status"`
	AccessTimeoutHandlerQueuename string    `json:"access_timeout_handler_queuename"`
	UseUnsupportedMobilesRegistry uint64    `json:"use_unsupported_mobiles_registry"`
	OriginName                    string    `json:"origin_name"`
}

func (m MessageModel) ToString() string {
	data := ""
	data += fmt.Sprintf("Type: %s ", m.Type)
	data += fmt.Sprintf("Content: %s ", m.Content)
	data += fmt.Sprintf("MobileNumber: %s ", m.MobileNumber)
	data += fmt.Sprintf("MobileCountryISOCode: %s ", m.MobileCountryISOCode)
	data += fmt.Sprintf("ShortNumber: %s ", m.ShortNumber)
	data += fmt.Sprintf("Telco: %s ", m.Telco)
	data += fmt.Sprintf("Created: %s ", m.Created)
	data += fmt.Sprintf("RoutingType: %s ", m.RoutingType)
	data += fmt.Sprintf("MatchedPattern: %s ", m.MatchedPattern)
	data += fmt.Sprintf("ServiceID: %s ", m.ServiceID)
	data += fmt.Sprintf("TelcoID: %s ", m.TelcoID)
	data += fmt.Sprintf("SessionAction: %s ", m.SessionAction)
	data += fmt.Sprintf("SessionParametersMap: %s ", m.SessionParametersMap)
	data += fmt.Sprintf("SessionTimeoutSeconds: %d", m.SessionTimeoutSeconds)
	data += fmt.Sprintf("Priority: %d", m.Priority)
	data += fmt.Sprintf("ClientID: %s ", m.ClientID)
	data += fmt.Sprintf("URL: %s ", m.URL)
	data += fmt.Sprintf("AccessTimeoutSeconds: %d", m.AccessTimeoutSeconds)
	data += fmt.Sprintf("RequestID: %d", m.RequestID)
	data += fmt.Sprintf("DefaultActionID: %d", m.DefaultActionID)
	data += fmt.Sprintf("ApplicationID: %d", m.ApplicationID)
	data += fmt.Sprintf("SessionID: %d", m.SessionID)
	data += fmt.Sprintf("Processed: %s ", m.Processed)
	data += fmt.Sprintf("MillisSinceRequest: %d", m.MillisSinceRequest)
	data += fmt.Sprintf("SessionApplicationName: %s ", m.SessionApplicationName)
	data += fmt.Sprintf("Sendafter: %s ", m.Sendafter)
	data += fmt.Sprintf("Sendbefore: %s ", m.Sendbefore)
	data += fmt.Sprintf("Sent: %s ", m.Sent)
	data += fmt.Sprintf("Status: %s ", m.Status)
	data += fmt.Sprintf("AccessTimeoutHandlerQueuename: %s ", m.AccessTimeoutHandlerQueuename)
	data += fmt.Sprintf("UseUnsupportedMobilesRegistry: %d", m.UseUnsupportedMobilesRegistry)
	data += fmt.Sprintf("OriginName: %s ", m.OriginName)
	return data
}

func (m *MessageModel) Name() string {
	return "message-model"
}

// utilzamos esta interface que tendra los metos que luego tendremos que utilizar en los distintos motores de base de datos
type Storage interface {
	InsertMessage(*MessageModel) error
	GetByID(uint64) (*MessageModel, error)
}

// el servicio sera el encargado de tener toda la logica del modelo de message
type Service struct {
	storage Storage
}

// retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// aqui deberas crear los metodos para que el main o desde donde quieras llamarlo pueda llamar al metodo que quieras utilizar
// insert permite insertar en la base un mensaje
func (s *Service) InsertMessage(m *MessageModel) error {
	return s.storage.InsertMessage(m)
}

// para traer un unico valor de la tabla message
func (s *Service) GetByID(id uint64) (*MessageModel, error) {
	return s.storage.GetByID(id)
}
