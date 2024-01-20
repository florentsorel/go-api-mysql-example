package data

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"time"
)

type Actor struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	CreationDate   time.Time  `json:"creation_date"`
	LastUpdateDate *time.Time `json:"last_update_date,omitempty"`
}

type ActorModel struct {
	DB *sql.DB
}

func (a ActorModel) Get(id int64) (*Actor, error) {
	if id < 1 {
		return nil, ErrorRecordNotFound
	}

	timeoutEnabled, err := strconv.ParseBool(os.Getenv("REQUEST_TIMEOUT_ENABLED"))
	if err != nil {
		timeoutEnabled = false
	}

	var query string
	if timeoutEnabled == true {
		query = `
		SELECT SLEEP(5) idActor, name, creationDate, lastUpdateDate
		FROM Actor
		WHERE idActor = ?`
	} else {
		query = `
		SELECT idActor, name, creationDate, lastUpdateDate
		FROM Actor
		WHERE idActor = ?`
	}

	var actor Actor

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// time.Sleep(2 * time.Second)

	err = a.DB.QueryRowContext(ctx, query, id).Scan(
		&actor.ID,
		&actor.Name,
		&actor.CreationDate,
		&actor.LastUpdateDate,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}

	return &actor, nil
}

func (a ActorModel) Insert(actor *Actor) (int64, error) {
	query := `INSERT INTO Actor (name, creationDate) VALUES(?, ?)`

	args := []any{actor.Name, actor.CreationDate}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := a.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func (a ActorModel) GetAll() ([]*Actor, error) {
	query := "SELECT idActor, name, creationDate, lastUpdateDate FROM Actor"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actors := []*Actor{}

	for rows.Next() {
		var actor Actor

		err := rows.Scan(
			&actor.ID,
			&actor.Name,
			&actor.CreationDate,
			&actor.LastUpdateDate,
		)
		if err != nil {
			return nil, err
		}

		actors = append(actors, &actor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}
