package database

import (
	"context"
	"filmLibrary/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func NewDBPool(configPath string) (*pgxpool.Pool, error) {
	configFile := configPath + string(os.PathSeparator) + "server.toml"
	config, err := config.ReadConfig(configFile)
	if err != nil {
		return nil, err
	}

	dsn := "user=" + config.Database.Username +
		" password=" + config.Database.Password +
		" host=" + config.Database.Host +
		" port=" + config.Database.Port +
		" dbname=" + config.Database.Database

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func CreateTables(pool *pgxpool.Pool) error {
	_, err := pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS actors (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			gender VARCHAR(255) NOT NULL,
			birthdate DATE NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS movies (
			id SERIAL PRIMARY KEY,
			title VARCHAR(150) NOT NULL,
			description TEXT,
			release_date DATE NOT NULL,
			rating FLOAT CHECK (rating >= 0 AND rating <= 10) NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS movie_actors (
			movie_id SERIAL REFERENCES movies(id),
			actor_id SERIAL REFERENCES actors(id),
			PRIMARY KEY (movie_id, actor_id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
