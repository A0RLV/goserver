package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Cleaned_Body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	bad := map[string]bool{"kerfuffle": true, "sharbert": true, "fornax": true}
	words := strings.Split(params.Body, " ")
	for index, word := range words {
		if bad[strings.ToLower(word)] {
			words[index] = "****"
		}
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Cleaned_Body: strings.Join(words, " "),
	})
}
