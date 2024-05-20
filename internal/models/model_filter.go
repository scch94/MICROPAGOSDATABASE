package models

import (
	"fmt"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/request"
)

type FilterModel struct {
	Id                   string
	Added                string
	Comment              string
	Direction            string
	MobileCountryisocode string
	MobileNumber         string
	Removed              string
	ShortNumber          string
	Useradded            string
	Userremoved          string
	Result               bool //este result solo ayuda a guarda el resultado de la base no hace parte de los paramtros
}

func (m *FilterModel) Name() string {
	return "FilterModel"
}

func (m FilterModel) ToString() string {
	return fmt.Sprintf("%v", m)
}

// utilzamos esta interface que tendra los metos que luego tendremos que utilizar en los distintos motores de base de datos
type FilterStorage interface {
	IsFilter(request.IsFiler) (*FilterModel, error)
}

// el servicio sera el encargado de tener toda la logica del modelo de message
type FilterService struct {
	filterStorage FilterStorage
}

// retorna un puntero de servicio
func NewFilterService(f FilterStorage) *FilterService {
	return &FilterService{filterStorage: f}
}

func (s *FilterService) IsFilter(r request.IsFiler) (*FilterModel, error) {
	return s.filterStorage.IsFilter(r)
}
