package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/cbodonnell/api-template/cache"
	"github.com/cbodonnell/api-template/config"
	"github.com/cbodonnell/api-template/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PSQLRepository struct {
	conf  config.Configuration
	cache cache.Cache
	db    *pgxpool.Pool
}

func NewPSQLRepository(_conf config.Configuration, _cache cache.Cache) Repository {
	return &PSQLRepository{
		conf:  _conf,
		cache: _cache,
		db:    connectDb(_conf.Db),
	}
}

func connectDb(s config.DataSource) *pgxpool.Pool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.User, s.Password, s.Dbname)
	db, err := pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	log.Printf("Connected to %s as %s\n", s.Dbname, s.User)
	return db
}

func (r *PSQLRepository) Close() {
	r.db.Close()
}

func (r *PSQLRepository) QueryUserByName(name string) (models.User, error) {
	sql := `SELECT * FROM users
	WHERE name = $1
	LIMIT 1`

	var user models.User
	err := r.db.QueryRow(context.Background(), sql, name).Scan(
		&user.ID,
		&user.Name,
		&user.Discoverable,
		&user.IRI,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
