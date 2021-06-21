package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davidlick/supermarket-api/internal/produce"
	"github.com/go-chi/chi"
)

func (s *server) produceGroup(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/produce", func(r chi.Router) {
			r.Get("/", s.handleGetAllProduce)
			r.Post("/", s.handleAddProduce)
			r.Route("/{produceCode}", func(r chi.Router) {
				r.Delete("/", s.handleDeleteProduce)
			})
		})
	})
}

func (s *server) handleAddProduce(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.writeError(ctx, w, err, http.StatusBadRequest)
		return
	}

	var items []produce.Item
	err = json.Unmarshal(body, &items)
	if err != nil {
		s.writeError(ctx, w, err, http.StatusBadRequest)
		return
	}

	err = s.produceSvc.Add(items)
	if err != nil {
		s.writeError(ctx, w, err, http.StatusInternalServerError)
		return
	}

	s.writeSuccess(ctx, w, nil, http.StatusCreated)
	return
}

func (s *server) handleGetAllProduce(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	items, err := s.produceSvc.All()
	if err != nil {
		s.writeError(ctx, w, err, http.StatusInternalServerError)
		return
	}

	s.writeSuccess(ctx, w, items, http.StatusOK)
	return
}

func (s *server) handleDeleteProduce(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	produceCode := chi.URLParam(r, "produceCode")
	if produceCode == "" {
		s.writeError(ctx, w, ErrUnrecognizedCode, http.StatusBadRequest)
		return
	}

	err := s.produceSvc.Remove(produce.Item{
		Code: produceCode,
	})
	if err != nil {
		s.writeError(ctx, w, err, http.StatusInternalServerError)
		return
	}

	s.writeSuccess(ctx, w, nil, http.StatusNoContent)
	return
}
