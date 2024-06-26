package handlers

import "github.com/ZFlucKZ/assessment-tax/db"

func calculatePersonalDeductionTax(income float64, personal float64) (float64, error) {
	personalDeductionDetail, err := db.GetDeductionAmountByDeductionType("Personal")
	if err != nil {
		return 0, err
	}

	return income - personalDeductionDetail.Amount, nil
}

func calculateDonationDeductionTax(income float64, donation float64) (float64, error) {
	donationDeductionDetail, err := db.GetDeductionAmountByDeductionType("Donation")
	if err != nil {
		return 0, err
	}

	donationAmount := donationDeductionDetail.Amount

	if donation > donationAmount {
		donation = donationAmount
	}

	return income - donation, nil
}

func calculateKReceiptDeductionTax(income float64, kReceipt float64) (float64, error) {
	kReceiptDeductionDetail, err := db.GetDeductionAmountByDeductionType("KReceipt")
	if err != nil {
		return 0, err
	}

	kReceiptAmount := kReceiptDeductionDetail.Amount

	if kReceipt > kReceiptAmount {
		kReceipt = kReceiptAmount
	}

	return income - kReceipt, nil
}