package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio-backend/internal/models"
	"time"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `josn:"version"`
	}{
		Status:  "active",
		Message: "Go Moies up and ",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *Application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	releaseData, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: releaseData,
		MPAARating:  "R",
		RunTime:     116,
		Description: "Really nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	raidersOfTheLostArk := models.Movie{
		ID:          2,
		Title:       "Raiders of the Lost Ark",
		ReleaseDate: releaseData,
		MPAARating:  "R",
		RunTime:     220,
		Description: "Another really nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, raidersOfTheLostArk)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
