package tavi

import (
	"log"
	"time"

	"github.com/signintech/gopdf"
)

// MakeTavi50 use for generate Tavi50 pdf
func MakeTavi50(data Tavi50) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	addTavi50Form(&pdf)

	if err := setupFont(&pdf); err != nil {
		log.Print(err.Error())
		return
	}
	addEmployerInfo(&pdf, data.Employer)
	addEmployeeInfo(&pdf, data.Employee)

	if err := addIncomeAmount(&pdf, data.Amount); err != nil {
		log.Print(err.Error())
		return
	}
	if err := addTaxAmount(&pdf, data.Amount, data.PercentTax); err != nil {
		log.Print(err.Error())
		return
	}

	addTaxWithheldCheckmark(&pdf)
	addTaxPayDate(&pdf, data.Time)
	addDocumentDate(&pdf, data.Time)

	pdf.WritePdf("tavi.pdf")
}

func addTavi50Form(pdf *gopdf.GoPdf) {
	pdf.Image("github.com/skyfoxs/tavi/images/tavi50.png", 0, 0, gopdf.PageSizeA4)
}

func setupFont(pdf *gopdf.GoPdf) (err error) {
	err = pdf.AddTTFFont("k2d", "github.com/skyfoxs/tavi/fonts/THK2DJuly8.ttf")
	if err != nil {
		return
	}
	err = pdf.AddTTFFont("inconsolata", "github.com/skyfoxs/tavi/fonts/Inconsolata-Regular.ttf")
	if err != nil {
		return
	}
	err = pdf.SetFont("k2d", "", 14)
	return
}

func addEmployerInfo(pdf *gopdf.GoPdf, data Person) {
	add(pdf, 58, 100, data.Name)
	add(pdf, 59, 124, data.Address)
	addTIN(pdf, 83, data.ID)
}

func addEmployeeInfo(pdf *gopdf.GoPdf, data Person) {
	add(pdf, 58, 173, data.Name)
	add(pdf, 59, 198, data.Address)
	addTIN(pdf, 151, data.ID)
}

// addTIN using for add Taxpayer identification number (TIN) 13 digits
func addTIN(pdf *gopdf.GoPdf, top float64, id string) {
	counter := 0
	left := [13]float64{377, 395, 407, 419, 431, 450, 462, 474, 486, 498, 516, 529, 547}
	for _, c := range id {
		add(pdf, left[counter], top, string(c))
		counter++
	}
}

func addIncomeAmount(pdf *gopdf.GoPdf, income float64) (err error) {
	formattedAmountBaht, formattedAmountSatang := makeIncomeAmount(income)

	err = pdf.SetFont("inconsolata", "", 12)
	if err != nil {
		return
	}

	add(pdf, 474-float64(len(formattedAmountBaht))*6, 309, formattedAmountBaht)
	add(pdf, 543, 309, "."+formattedAmountSatang)

	add(pdf, 474-float64(len(formattedAmountBaht))*6, 648, formattedAmountBaht)
	add(pdf, 543, 648, "."+formattedAmountSatang)
	return
}

func addTaxAmount(pdf *gopdf.GoPdf, income, percentTax float64) (err error) {
	formattedTaxBaht, formattedTaxSatang, formattedTaxMessage := makeTaxAmount(income, percentTax)

	err = pdf.SetFont("inconsolata", "", 12)
	if err != nil {
		return
	}

	add(pdf, 545-float64(len(formattedTaxBaht))*6, 309, formattedTaxBaht)
	add(pdf, 543, 309, "."+formattedTaxSatang)

	add(pdf, 545-float64(len(formattedTaxBaht))*6, 648, formattedTaxBaht)
	add(pdf, 543, 648, "."+formattedTaxSatang)

	err = pdf.SetFont("k2d", "", 14)
	if err != nil {
		return
	}
	add(pdf, 183, 670, formattedTaxMessage)
	return
}

func addTaxWithheldCheckmark(pdf *gopdf.GoPdf) {
	add(pdf, 86, 708, "x")
}

func addTaxPayDate(pdf *gopdf.GoPdf, date time.Time) {
	add(pdf, 334, 309, formatDate(date))
}

func addDocumentDate(pdf *gopdf.GoPdf, date time.Time) {
	add(pdf, 368, 754, formatDate(date))
}

func add(pdf *gopdf.GoPdf, x, y float64, text string) {
	pdf.SetX(x)
	pdf.SetY(y)
	pdf.Cell(nil, text)
}

func makeIncomeAmount(income float64) (formattedAmountBaht, formattedAmountSatang string) {
	incomeWithoutFraction := int64(income * 100)
	amountSatang := incomeWithoutFraction % 100
	amountBaht := (incomeWithoutFraction - amountSatang) / 100

	formattedAmountBaht = formatMoney(amountBaht)
	formattedAmountSatang = formatMoney(amountSatang)
	return
}

func makeTaxAmount(income, percentTax float64) (formattedTaxBaht, formattedTaxSatang, formattedTaxMessage string) {
	tax := income * percentTax / 100.0
	taxWithoutFraction := int64(tax * 100)
	taxSatang := taxWithoutFraction % 100
	taxBaht := (taxWithoutFraction - taxSatang) / 100

	formattedTaxBaht = formatMoney(taxBaht)
	formattedTaxSatang = formatMoney(taxSatang)

	formattedTaxMessage = getFormatted(taxBaht) + "บาท"
	if taxSatang > 0 {
		formattedTaxMessage += getFormatted(taxSatang) + "สตางค์"
	}
	formattedTaxMessage += "ถ้วน"
	return
}
