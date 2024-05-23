package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/request"
	"github.com/scch94/ins_log"
)

type MysqlDomain struct {
	db *sql.DB
}

func NewMysqlDomain(db *sql.DB) *MysqlDomain {
	return &MysqlDomain{db}
}

const (
	mySQLGetDomain = `SELECT d.name FROM user u JOIN domain d ON u.domain_id = d.id where u.username=?`
)

func (p *MysqlDomain) GetUserDomain(r request.GetUserDomain, ctx context.Context) (*models.DomainModel, error) {

	//traemos el contexto y le setiamos el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	//creamos la variable que guardara el dominio del usuario
	var domainModel models.DomainModel

	ins_log.Tracef(ctx, "starting to get the domain for the user :%s", r.UserName)

	//creamos variable que Captura el tiempo de inicio de la operación
	startTime := time.Now()

	//realizamos la consula
	ins_log.Tracef(ctx, "this is the QUERY: %s and the params: Username=%s,", mySQLGetDomain, r.UserName)
	err := p.db.QueryRow(mySQLGetDomain, r.UserName).Scan(&domainModel.Domainname)
	switch {
	case err == sql.ErrNoRows:
		ins_log.Infof(ctx, "didnt find domain for the user %s", r.UserName)
		domainModel.Domainname = ""
		domainModel.Result = err.Error()
		err = nil
	case err != nil:
		ins_log.Fatalf(ctx, "query error %v", err)
	default:
		domainModel.Result = "the domain name is: " + domainModel.Domainname
	}

	// Calcula la duración de la operacion
	duration := time.Since(startTime)
	ins_log.Infof(ctx, "the query in the database tooks: %v", duration)
	ins_log.Infof(ctx, "and this is the domain: %v", domainModel.Domainname)

	return &domainModel, err
}
