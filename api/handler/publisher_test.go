package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/entity"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/clean-architecture-go-v2/usecase/publisher/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listPublisher(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePublisherHandlers(r, *n, service)
	path, err := r.GetRoute("listPublisher").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/publisher", path)
	b := &entity.Publisher{}
	service.EXPECT().
		ListPublishers().
		Return([]*entity.Publisher{b}, nil)
	ts := httptest.NewServer(listPublishers(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listPublisher_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listPublishers(service))
	defer ts.Close()
	service.EXPECT().
		SearchPublishers("publisher of publisher").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=publisher+of+publisher")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listPublisher_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	b := &entity.Publisher{}
	service.EXPECT().
		SearchPublishers("ozzy").
		Return([]*entity.Publisher{b}, nil)
	ts := httptest.NewServer(listPublishers(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createPublisher(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePublisherHandlers(r, *n, service)
	path, err := r.GetRoute("createPublisher").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/publisher", path)

	service.EXPECT().
		CreatePublisher(gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createPublisher(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
		"name": "I Am Ozzy",
		"address": "Ozzy Osbourne"
		}`)
	resp, _ := http.Post(ts.URL+"/v1/publisher", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var b *entity.Publisher
	json.NewDecoder(resp.Body).Decode(&b)
	assert.Equal(t, "I Am Ozzy", b.Name)
}

func Test_getPublisher(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePublisherHandlers(r, *n, service)
	path, err := r.GetRoute("getPublisher").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/publisher/{id}", path)
	b := &entity.Publisher{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetPublisher(b.ID).
		Return(b, nil)
	handler := getPublisher(service)
	r.Handle("/v1/publisher/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/publisher/" + strconv.FormatInt(b.ID, 10))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.Publisher
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}

func Test_deletePublisher(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakePublisherHandlers(r, *n, service)
	path, err := r.GetRoute("deletePublisher").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/publisher/{id}", path)
	b := &entity.Publisher{
		ID: entity.NewID(),
	}
	service.EXPECT().DeletePublisher(b.ID).Return(nil)
	handler := deletePublisher(service)
	req, _ := http.NewRequest("DELETE", "/v1/publisher/"+strconv.FormatInt(b.ID, 10), nil)
	r.Handle("/v1/publishermark/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
