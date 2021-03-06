package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testingTask/pkg/error"
	"testingTask/pkg/links"
)

type LinksHandler struct {
	LinksRepo links.LinksRepo
}

func (lh *LinksHandler) FromLongToShort(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: "data have not arrived"}
		log.Printf("read body error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorMessage.Error())
	}

	link := &links.Link{}

	err = json.Unmarshal(body, link)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: "invalid input"}
		log.Printf("json unmarshal error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorMessage.Error())
	}

	_, err = url.ParseRequestURI(link.Data)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: "invalid input"}
		log.Printf("parse uri error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorMessage.Error())
	}

	shortURL, err := lh.LinksRepo.Add(link.Data)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: err.Error()}
		log.Printf("add new link error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorMessage.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(shortURL))
	w.WriteHeader(http.StatusOK)
}

func (lh *LinksHandler) FromShortToLong(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: "data have not arrived"}

		log.Printf("read body error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorMessage.Error())
	}

	link := &links.Link{}

	err = json.Unmarshal(body, link)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: "invalid input"}

		log.Printf("json unmarshal error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorMessage.Error())
	}

	longURL, err := lh.LinksRepo.Get(link.Data)

	if err != nil {
		ErrorMessage := error.Error{
			ErrMsg: err.Error()}

		log.Printf("get link error: %s", err.Error())
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorMessage.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(longURL))

}
