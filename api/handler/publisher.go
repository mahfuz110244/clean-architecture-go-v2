package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eminetto/clean-architecture-go-v2/api/presenter"
	"github.com/eminetto/clean-architecture-go-v2/usecase/publisher"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/entity"
	"github.com/gorilla/mux"
)

func listPublishers(service publisher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Publishers"
		errorMessageNoData := "No data found"
		var data []*entity.Publisher
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListPublishers()
		default:
			data, err = service.SearchPublishers(name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessageNoData))
			return
		}
		var toJ []*presenter.Publisher
		for _, d := range data {
			toJ = append(toJ, &presenter.Publisher{
				ID:      d.ID,
				Name:    d.Name,
				Address: d.Address,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createPublisher(service publisher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding Publisher"
		var input struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreatePublisher(input.Name, input.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Publisher{
			ID:      id,
			Name:    input.Name,
			Address: input.Address,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getPublisher(service publisher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Publisher"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetPublisher(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Publisher{
			ID:      data.ID,
			Name:    data.Name,
			Address: data.Address,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deletePublisher(service publisher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing Publishermark"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeletePublisher(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakePublisherHandlers make url handlers
func MakePublisherHandlers(r *mux.Router, n negroni.Negroni, service publisher.UseCase) {
	r.Handle("/v1/publisher", n.With(
		negroni.Wrap(listPublishers(service)),
	)).Methods("GET", "OPTIONS").Name("listPublisher")

	r.Handle("/v1/publisher", n.With(
		negroni.Wrap(createPublisher(service)),
	)).Methods("POST", "OPTIONS").Name("createPublisher")

	r.Handle("/v1/publisher/{id}", n.With(
		negroni.Wrap(getPublisher(service)),
	)).Methods("GET", "OPTIONS").Name("getPublisher")

	r.Handle("/v1/publisher/{id}", n.With(
		negroni.Wrap(deletePublisher(service)),
	)).Methods("DELETE", "OPTIONS").Name("deletePublisher")
}
