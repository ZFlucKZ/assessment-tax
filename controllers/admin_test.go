package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUpdatePersonalDeduction(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		wantStatusCode int
	}{
		{"Error when Amount -10", -10.0, http.StatusBadRequest},
		{"Error when Amount 0", 0.0, http.StatusBadRequest},
		{"Error when Amount 9,999", 9999.0, http.StatusBadRequest},
		{"Error when Amount 100,001", 100001.0, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/admin/deductions/personal", strings.NewReader(fmt.Sprintf(`{"amount": %v}`, tt.amount)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			aCl := AdminController{}
			err := aCl.UpdatePersonalDeduction(c)
			if err != nil {
				t.Errorf("UpdatePersonalDeduction(%v) = %v; want %v", tt.amount, err, tt.wantStatusCode)
			}
		})
	}
}