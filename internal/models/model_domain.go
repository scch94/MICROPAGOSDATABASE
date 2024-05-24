package models

import (
	"context"
	"fmt"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/request"
)

type DomainModel struct {
	Domainname string
	Username   string
	Password   string
	Result     string
}

func (m *DomainModel) Name() string {
	return "FilterModel"
}

func (m DomainModel) ToString() string {
	return fmt.Sprintf("%v", m)
}

// utilzamos esta interface que tendra los metos que luego tendremos que utilizar en los distintos motores de base de datos
type DomainStorage interface {
	GetUserDomain(request.GetUserDomain, context.Context) (*DomainModel, error)
}

// el servicio sera el encargado de tener toda la logica del modelo de domain
type DomainService struct {
	domainStorage DomainStorage
}

// retorna un puntero de servicio
func NewDomainService(d DomainStorage) *DomainService {
	return &DomainService{domainStorage: d}
}

func (s *DomainService) GetUserDomain(r request.GetUserDomain, ctx context.Context) (*DomainModel, error) {
	return s.domainStorage.GetUserDomain(r, ctx)
}
