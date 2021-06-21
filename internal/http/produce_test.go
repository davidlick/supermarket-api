package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/davidlick/supermarket-api/internal/produce"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServer_handleAddProduce(t *testing.T) {
	tests := []struct {
		test       string
		body       string
		expectFunc func(mockProduceSvc *MockProduceService)
		assertFunc func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			test: "it should successfully add valid produce",
			body: `[{"code":"test","Name":"test","price":{"amount":101,"currency":"USD"}}]`,
			expectFunc: func(mockProduceSvc *MockProduceService) {
				mockProduceSvc.EXPECT().Add([]produce.Item{
					{
						Code:  "test",
						Name:  "test",
						Price: money.New(101, "USD"),
					},
				}).Return(nil)
			},
			assertFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, w.Code)
			},
		},
		{
			test:       "it should respond bad request if body isn't an array of produce items",
			body:       ``,
			expectFunc: func(mockProduceSvc *MockProduceService) {},
			assertFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
		{
			test: "it should respond internal server error if adding to service fails",
			body: `[{"code":"test","Name":"test","price":{"amount":101,"currency":"USD"}}]`,
			expectFunc: func(mockProduceSvc *MockProduceService) {
				mockProduceSvc.EXPECT().Add(gomock.Any()).Return(errors.New("test error"))
			},
			assertFunc: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := httptest.NewRequest(http.MethodPost, "/v1/produce", strings.NewReader(tc.body))
			w := httptest.NewRecorder()

			noopLogger := logrus.New()
			noopLogger.SetOutput(ioutil.Discard)

			mockProduceSvc := NewMockProduceService(ctrl)
			tc.expectFunc(mockProduceSvc)

			s := NewServer(3000, noopLogger, "test", mockProduceSvc)

			handler := http.HandlerFunc(s.handleAddProduce)
			handler.ServeHTTP(w, r)

			tc.assertFunc(t, w)
		})
	}
}
