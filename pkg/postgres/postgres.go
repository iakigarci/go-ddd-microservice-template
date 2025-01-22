package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	_ "github.com/lib/pq"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	DB *sql.DB
}

var pg *Postgres
var hdlOnce sync.Once

func NewOrGetSingleton(config *config.Config) *Postgres {
	hdlOnce.Do(func() {
		postgres, err := initPg(config)
		if err != nil {
			panic(err)
		}
		pg = postgres
	})
	return pg
}

func initPg(config *config.Config) (*Postgres, error) {
	pg = &Postgres{
		maxPoolSize:  config.Postgres.PoolMax,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.DB, err = sql.Open("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				config.Postgres.Host,
				config.Postgres.Port,
				config.Postgres.User,
				config.Postgres.Password,
				config.Postgres.Name,
				config.Postgres.SSLMode,
			),
		)
		if err == nil {
			pg.DB.SetMaxOpenConns(pg.maxPoolSize)
			pg.DB.SetMaxIdleConns(pg.maxPoolSize)

			if err = pg.DB.Ping(); err == nil {
				break
			}
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.DB.PingContext(ctx)
}

func (p *Postgres) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return p.DB.BeginTx(ctx, &sql.TxOptions{})
}
