package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/doniacld/simple-web-api/db"
	"github.com/doniacld/simple-web-api/todo"

	"github.com/gorilla/mux"
)

const (
	contentType  = "Content-Type"
	mimeTypeJSON = "application/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodosRetrieve retrieves a todo
func TodosRetrieve(w http.ResponseWriter, r *http.Request) {
	tds := db.RepoRetrieveTodos()
	w.Header().Set(contentType, mimeTypeJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tds); err != nil {
		panic(err)
	}
}

// TodoShow shows a todo
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]

	todoIDint, err := strconv.Atoi(todoID)
	if err != nil {
		panic(err)
	}
	td, err := db.RepoFindTodo(todoIDint)
	if err != nil {
		panic(err)
	}

	w.Header().Set(contentType, mimeTypeJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(td); err != nil {
		panic(err) //TODO DONIA replace it by a nice catch
	}
}

// TodoCreate creates a todo
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var td todo.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // protect against malicious attacks on your server. Imagine if someone wanted to send you 500GBs of json...
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &td); err != nil {
		w.Header().Set(contentType, "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.RepoCreateTodo(td)
	w.Header().Set(contentType, mimeTypeJSON)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// TodoDelete deletes a todo
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]

	todoIDint, err := strconv.Atoi(todoID)
	if err != nil {
		panic(err)
	}
	err = db.RepoDestroyTodo(todoIDint)
	if err != nil {
		panic(err)
	}

	w.Header().Set(contentType, mimeTypeJSON)
	w.WriteHeader(http.StatusNoContent)
}
