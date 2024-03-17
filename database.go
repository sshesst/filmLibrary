package database

import (
	"context"
	"filmLibrary/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func GetPool() (*pgxpool.Pool, error) {
	configPath := "config"
	pool, err := NewDBPool(configPath)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

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
			description TEXT CHECK (LENGTH(description) <= 1000),
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

	_, err = pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE,
			is_auth BOOLEAN NOT NULL DEFAULT FALSE
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
