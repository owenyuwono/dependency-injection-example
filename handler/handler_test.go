package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"dependency-injection-example/handler/mock"
	"dependency-injection-example/model"

	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
)

func TestHandler_InsertData(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mock           func(*mock.Mockrepository)
		want           string
		wantStatusCode int
	}{
		{"success case", `{"name":"123"}`, func(m *mock.Mockrepository) {
			m.EXPECT().Insert(model.Data{
				Name: "123",
			}).Return("randomstring", nil)
		}, `{"id":"randomstring","status":"OK"}`, http.StatusOK},
		{"failure case - dependency", `{"name":"123"}`, func(m *mock.Mockrepository) {
			m.EXPECT().Insert(model.Data{
				Name: "123",
			}).Return("", errors.New("unknown error"))
		}, `{"status":"internal server error"}`, http.StatusInternalServerError},
		{"failure case - dependency", `}`, func(m *mock.Mockrepository) {}, `{"status":"bad request"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockrepository(ctrl)
			tt.mock(m)

			h := New(m)
			// create a mock http object
			res := httptest.NewRecorder()
			b := bytes.NewBufferString(tt.body)
			req, err := http.NewRequest(http.MethodPost, "/insert", b)
			if err != nil {
				panic(err)
			}
			_, r := gin.CreateTestContext(res)
			r.POST("/insert", h.InsertData)
			r.ServeHTTP(res, req)
			if strings.Compare(tt.want, res.Body.String()) != 0 {
				t.Errorf("want: %s, got: %s", tt.want, res.Body.String())
			}
			if tt.wantStatusCode != res.Code {
				t.Errorf("want: %d, got: %d", tt.wantStatusCode, res.Code)
			}
			// h.InsertData(tt.args)
		})
	}
}
