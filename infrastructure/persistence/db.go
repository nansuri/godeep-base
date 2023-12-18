package persistence

import (
	"os"
	"time"

	baseEntity "bitbucket.org/be-proj/osp-base/domain/entity/base"
	enum "bitbucket.org/be-proj/osp-base/domain/enum/base"
	baseRepo "bitbucket.org/be-proj/osp-base/domain/repository/base"
	"bitbucket.org/be-proj/osp-base/infrastructure/persistence/base"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type Repositories struct {
	db   *gorm.DB
	Base baseRepo.BaseRepository
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	msql := mysql.Config{}
	logLevel := logger.Info

	logrus.Debug(msql)

	// db config
	dsn := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Asia%2FJakarta"

	if os.Getenv("APP_ENV") == enum.EnvDevelopment || os.Getenv("APP_ENV") == enum.EnvStaging {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return nil, err
	}

	db.Use(dbresolver.Register(
		dbresolver.Config{}).
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		SetMaxIdleConns(2).
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		SetMaxOpenConns(2).
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		SetConnMaxLifetime(2 * time.Minute))

	// db.SingularTable(true)

	return &Repositories{
		db:   db,
		Base: base.NewBaseRepository(db),
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDb.Close()
}

// Put your entity here to do auto DDL / auto migrate
func (s *Repositories) Automigrate() error {

	return s.db.AutoMigrate(
		&baseEntity.TmBaseContext{},
	)
}
