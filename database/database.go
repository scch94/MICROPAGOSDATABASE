package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/scch94/MICROPAGOSDATABASE.git/config"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/ins_log"
)

// DB Y ONCE once nos permite hacer el singleton
var (
	dbMessage     *sql.DB
	dbUsers       *sql.DB
	dbMessageOnce sync.Once
	dbUsersOnce   sync.Once
)

//driver of database

type Driver string

const (
	MySQL      Driver = "MYSQL"
	MySQLUsers Driver = "MYSQL_USERS"
)

// con new creamos la coneccion a la base
func New(d Driver, ctx context.Context) {
	//traemos el contexto y le setiamos el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")
	switch d {
	case MySQL:
		newMySQLDB(ctx)
	case MySQLUsers:
		newMySQLDbUser(ctx)
	default:

		ins_log.Info(ctx, "not implemented")
	}
}

// func newPostgresDb() {
// 	dbMessageOnce.Do(func() {

//			var err error
//			dbMessage, err = sql.Open("postgres", postgresConnectionString.ConnectionString)
//			if err != nil {
//				ins_log.Fatalf(ctx, "cant open postgres database with string connection :%s , with error :%s", postgresConnectionString, err)
//				return
//			}
//			if err = dbMessage.Ping(); err != nil {
//				ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
//				return
//			}
//			ins_log.Info(ctx, "conected to the postgres message database")
//		})
//	}
//
//	func newPostgresDbUser() {
//		dbUsersOnce.Do(func() {
//			var err error
//			dbUsers, err = sql.Open("mysql", mysqlConnectionSrtingtousersdatabase.ConnectionString)
//			if err != nil {
//				ins_log.Fatalf(ctx, "cant open myssql database with string connection :%s , with error :%s", postgresConnectionString, err)
//			}
//			if err = dbUsers.Ping(); err != nil {
//				ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
//			}
//			ins_log.Info(ctx, "conected to postgres user database")
//		})
//	}
func newMySQLDB(ctx context.Context) {
	dbMessageOnce.Do(func() {
		ctx = ins_log.SetPackageNameInContext(ctx, "database")
		var err error
		dbMessage, err = sql.Open("mysql", config.Config.MySQLConnection.Raven.ConnectionString)
		if err != nil {
			ins_log.Fatalf(ctx, "cant open myssql database with string connection :%s , with error :%s", config.Config.MySQLConnection.Raven.ConnectionString, err)
		}
		if err = dbMessage.Ping(); err != nil {
			ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
		}
		ins_log.Info(ctx, "conected to the mysql message database")

		// dbMessage.SetMaxOpenConns(config.Config.MySQLConnection.Raven.MaxOpenConns)
		// dbMessage.SetMaxIdleConns(config.Config.MySQLConnection.Raven.MaxIdleConns)
		// dbMessage.SetConnMaxLifetime(time.Duration(config.Config.MySQLConnection.Raven.ConnMaxLifeTime) * time.Millisecond)
	})
}
func newMySQLDbUser(ctx context.Context) {
	dbUsersOnce.Do(func() {
		ctx = ins_log.SetPackageNameInContext(ctx, "database")

		var err error
		dbUsers, err = sql.Open("mysql", config.Config.MySQLConnection.Weaver.ConnectionString)
		if err != nil {
			ins_log.Fatalf(ctx, "cant open myssql database with string connection :%s , with error :%s", config.Config.MySQLConnection.Weaver.ConnectionString, err)
		}
		if err = dbUsers.Ping(); err != nil {
			ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
		}
		ins_log.Info(ctx, "conected to the mysql users database")

		///setiamos los valores de config
		// dbUsers.SetMaxOpenConns(config.Config.MySQLConnection.Weaver.MaxOpenConns)
		// dbUsers.SetMaxIdleConns(config.Config.MySQLConnection.Weaver.MaxIdleConns)
		// dbUsers.SetConnMaxLifetime(time.Duration(config.Config.MySQLConnection.Weaver.ConnMaxLifeTime) * time.Millisecond)

	})
}

// Pool returna una unica isntancia del pool de conexiones de la base de datos q se creo en el metodo newMysqlDb
func PoolMessage() *sql.DB {
	return dbMessage
}
func PoolUsers() *sql.DB {
	return dbUsers
}

func StringToNull(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func TimeToNull(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Time: t, Valid: false}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func Uint64ToNull(i uint64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: int64(i), Valid: true}
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func ScanRowMessage(s scanner, ctx context.Context) (*models.MessageModel, error) {

	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	ins_log.Debug(ctx, "starting to scan message models")
	m := &models.MessageModel{}
	contentNull := sql.NullString{}
	err := s.Scan(
		&contentNull,
	)
	if err != nil {
		ins_log.Infof(ctx, "SQL response"+err.Error())
		return &models.MessageModel{}, err
	}
	// Assign null variables to struct fields
	ins_log.Trace(ctx, "scanning process end")
	// Assign values from Null variables to struct fields, extracting String, Int, and Time values properly
	m.Content = contentNull.String
	return m, nil
}
