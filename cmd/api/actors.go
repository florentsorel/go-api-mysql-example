package main

import (
	"api-mysql-example/data"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (app *application) showActorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	actor, err := app.models.Actor.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		case errors.Is(err, context.DeadlineExceeded):
			app.requestTimeoutResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, actor, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.logger.Info("actor info", "actor", actor)
}

func (app *application) createActorHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	actor := &data.Actor{
		Name:         input.Name,
		CreationDate: time.Now().UTC(),
	}

	id, err := app.models.Actor.Insert(actor)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			app.requestTimeoutResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	actor.ID = id

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/actors/%d", actor.ID))

	err = app.writeJSON(w, http.StatusCreated, actor, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listActorHandler(w http.ResponseWriter, r *http.Request) {
	actor, err := app.models.Actor.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			app.requestTimeoutResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, actor, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
