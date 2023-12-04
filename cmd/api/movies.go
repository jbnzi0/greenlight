package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jbnzi0/greenlight/internal/data"
)

func (app *Application) getMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        1,
		CreatedAt: time.Now(),
		Title:     "Day Before tomorrow",
		Year:      2010,
		Runtime:   2000,
		Genres:    []string{"action"},
		Version:   1,
	}

	envelope := Envelope{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusOK, envelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) createMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
