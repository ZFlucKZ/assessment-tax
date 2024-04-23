package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCalculateTax(t *testing.T) {
	// INFO: Below Test case for personal deduction is default value(60,000)
	// tests_progressive_rate := []struct {
	// 	name string
	// 	totalIncome float64
	// 	want        float64
	// 	statusCode int
	// }{
	// 	{"Error when Income -10",-10.0  , 0.0 ,http.StatusBadRequest},
	// 	{"Tax 0 when Income 0",0.0  ,0.0, http.StatusOK},
	// 	{"Tax 0 when Income 70,000",60000.0  ,0.0, http.StatusOK},
	// 	{"Tax 0 when Income 150,000",150000.0  ,0.0, http.StatusOK},
	// 	{"Tax 15,000 when Income 300,000",300000.0  ,9000.0, http.StatusOK},
	// 	{"Tax 35,000 when Income 500,000",500000.0  ,29000.0, http.StatusOK},
	// 	{"Tax 72,500 when Income 750,000",750000.0  ,63500.0, http.StatusOK},
	// 	{"Tax 11,000 when Income 1,000,000",1000000.0  ,101000.0, http.StatusOK},
	// 	{"Tax 210,000 when Income 1,500,000",1500000.0  ,198000.0, http.StatusOK},
	// 	{"Tax 310,000 when Income 2,000,000",2000000.0  ,298000.0, http.StatusOK},
	// 	{"Tax 485,000 when Income 2,500,000",2500000.0  ,464000.0, http.StatusOK},
	// }

	// INFO: Below Test case for personal deduction is not default value(60,000)
	tests_progressive_rate := []struct {
		name string
		totalIncome float64
		statusCode int
	}{
		{"StatusBadRequest when Income -10",-10.0, http.StatusBadRequest},
		{"StatusOK when Income 0",0.0, http.StatusOK},
		{"StatusOK when Income 70,000",60000.0, http.StatusOK},
		{"StatusOK when Income 150,000",150000.0, http.StatusOK},
		{"StatusOK when Income 300,000",300000.0, http.StatusOK},
		{"StatusOK when Income 500,000",500000.0, http.StatusOK},
		{"StatusOK when Income 750,000",750000.0, http.StatusOK},
		{"StatusOK when Income 1,000,000",1000000.0, http.StatusOK},
		{"StatusOK when Income 1,500,000",1500000.0, http.StatusOK},
		{"StatusOK when Income 2,000,000",2000000.0, http.StatusOK},
		{"StatusOK when Income 2,500,000",2500000.0, http.StatusOK},
	}

	for _, tt := range tests_progressive_rate {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(fmt.Sprintf(`{"totalIncome": %v}`, tt.totalIncome)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// testing CalculateTaxHandler for correct status code
			tCl := TaxController{}
			err := tCl.CalculateTaxHandler(c);
			if err != nil {
				t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, tt.statusCode)
			}

			if rec.Code != tt.statusCode {
				t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, rec.Code, tt.statusCode)
			}

			// testing CalculateTaxHandler for correct tax calculation
			// if rec.Code == http.StatusOK {
			// 	var taxResponse dto.TaxResponse
			// 	err := json.Unmarshal(rec.Body.Bytes(), &taxResponse)
		
			// 	if err != nil {
			// 		t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, nil)
			// 	}

			// 	if taxResponse.Tax != tt.want {
			// 		t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, taxResponse.Tax, tt.want)
			// 	}
			// }
		})
	}
}