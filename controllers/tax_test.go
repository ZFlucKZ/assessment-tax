package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZFlucKZ/assessment-tax/dto"
	"github.com/labstack/echo/v4"
)

// INFO: Calculate without personal deduction
// personal deduction changes together with Database before testing
var personalDeduction float64 = 60000.0
var wht float64 = 25000.0
var donation float64 = 10000.0

func TestCalculateTax(t *testing.T) {
	tests_progressive_rate := []struct {
		name string
		totalIncome float64
		want        float64
		statusCode int
	}{
		{"Error when Income -10",-10000000.0 , 0.0 ,http.StatusBadRequest},
		// {"Tax 0 when Income 60,000",60000.0 ,0.0, http.StatusOK},
		// {"Tax 0 when Income 150,000",150000.0 ,0.0, http.StatusOK},
		// {"Tax 15,000 when Income 300,000",300000.0 ,15000.0, http.StatusOK},
		// {"Tax 35,000 when Income 500,000",500000.0 ,35000.0, http.StatusOK},
		// {"Tax 72,500 when Income 750,000",750000.0 ,72500.0, http.StatusOK},
		// {"Tax 110,000 when Income 1,000,000",1000000.0 ,110000.0, http.StatusOK},
		// {"Tax 210,000 when Income 1,500,000",1500000.0 ,210000.0, http.StatusOK},
		// {"Tax 310,000 when Income 2,000,000",2000000.0 ,310000.0, http.StatusOK},
		// {"Tax 485,000 when Income 2,500,000",2500000.0 ,485000.0, http.StatusOK},
	}

	for _, tt := range tests_progressive_rate {
		t.Run(tt.name, func(t *testing.T) {
			tt.totalIncome += personalDeduction
			if donation > 100000 {
				donation = 100000
			}
			tt.totalIncome += donation

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(fmt.Sprintf(`{"totalIncome": %v, "allowances": [{"allowanceType": "donation", "amount": %v}]}`, tt.totalIncome, donation)))
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
				t.Errorf("CalculateTaxHandler(%v) StatusCode = %v; want %v", tt.totalIncome, rec.Code, tt.statusCode)
			}

			if rec.Code == http.StatusOK {
				var taxResponse dto.TaxResponse
				err := json.Unmarshal(rec.Body.Bytes(), &taxResponse)
	
				if err != nil {
					t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, nil)
				}

				if taxResponse.Tax != tt.want {
					t.Errorf("CalculateTaxHandler(%v) Tax = %v; want %v", tt.totalIncome, taxResponse.Tax, tt.want)
				}
			}
		})
	}
}

func TestCalculateTaxWithWht(t *testing.T) {
	tests_with_wht := []struct {
		name string
		totalIncome float64
		wht 			 float64
		want        float64
		statusCode int
	}{
		{"Error when Income -10",-10000000.0, wht, math.Abs(0.0 - wht), http.StatusBadRequest},
		// {"Income 60,000",60000.0, wht,math.Abs(0.0 - wht), http.StatusOK},
		// {"Income 150,000",150000.0, wht,math.Abs(0.0 - wht), http.StatusOK},
		// {"Income 300,000",300000.0, wht,math.Abs(15000.0 - wht), http.StatusOK},
		// {"Income 500,000",500000.0, wht,math.Abs(35000.0 - wht), http.StatusOK},
		// {"Income 750,000",750000.0, wht,math.Abs(72500.0 - wht), http.StatusOK},
		// {"Income 1,000,000",1000000.0, wht,math.Abs(110000.0 - wht), http.StatusOK},
		// {"Income 1,500,000",1500000.0, wht,math.Abs(210000.0 - wht), http.StatusOK},
		// {"Income 2,000,000",2000000.0, wht,math.Abs(310000.0 - wht), http.StatusOK},
		// {"Income 2,500,000",2500000.0, wht,math.Abs(485000.0 - wht), http.StatusOK},
	}

	for _, tt := range tests_with_wht {
		t.Run(tt.name, func(t *testing.T) {
			tt.totalIncome += personalDeduction
			if donation > 100000 {
				donation = 100000
			}
			tt.totalIncome += donation

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(fmt.Sprintf(`{"totalIncome": %v, "wht": %v, "allowances": [{"allowanceType": "donation", "amount": %v}]}`, tt.totalIncome, wht, donation)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// testing CalculateTaxHandler for correct status code
			tCl := TaxController{}
			err := tCl.CalculateTaxHandler(c);
			if err != nil {
				t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, tt.statusCode)
			}

			if tt.wht < 0 || tt.wht > tt.totalIncome {
				tt.statusCode = http.StatusBadRequest
			}

			if rec.Code != tt.statusCode {
				t.Errorf("CalculateTaxHandler(%v) StatusCode = %v; want %v", tt.totalIncome, rec.Code, tt.statusCode)
			}

			if rec.Code == http.StatusOK {
				var taxResponse dto.TaxResponse
				var TaxRefundResponse dto.TaxRefundResponse
				err := json.Unmarshal(rec.Body.Bytes(), &taxResponse)
				if err != nil {
					t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, nil)
				}

				err = json.Unmarshal(rec.Body.Bytes(), &TaxRefundResponse)
				if err != nil {
					t.Errorf("CalculateTaxHandler(%v) = %v; want %v", tt.totalIncome, err, nil)
				}

				if taxResponse.Tax != 0.0 && TaxRefundResponse.TaxRefund != 0.0 {
					t.Errorf("Tax and TaxRefund cannot be both not 0.0")
				}

				if taxResponse.Tax != 0.0 && taxResponse.Tax != tt.want {
					t.Errorf("CalculateTaxHandler(%v) Tax = %v; want %v", tt.totalIncome, taxResponse.Tax, tt.want)
				}

				if TaxRefundResponse.TaxRefund != 0.0 && TaxRefundResponse.TaxRefund != tt.want {
					t.Errorf("CalculateTaxHandler(%v) TaxRefund = %v; want %v", tt.totalIncome, TaxRefundResponse.TaxRefund, tt.want)
				}
			}
		})
	}
}