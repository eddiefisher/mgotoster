// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Tue, 27 Feb 2018]

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eddiefisher/mgotoster/src/dao"
	"github.com/eddiefisher/mgotoster/src/models"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func init() {
	dao.Connect()
}

func AllMovies(w http.ResponseWriter, r *http.Request) {
	movies := models.Movies{}
	err := movies.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func FindMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie := models.Movie{}
	err := movie.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	fmt.Printf("%#v\n", movie)
	respondWithJson(w, http.StatusOK, movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	movie := models.Movie{}
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := movie.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	movie := models.Movie{}
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := movie.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	movie := models.Movie{}
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := movie.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovie).Methods("GET")

	if err := http.ListenAndServe(":3030", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
