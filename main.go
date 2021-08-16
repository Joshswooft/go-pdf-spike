package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"

	"github.com/JoshSwooft/go-pdf-spike/ordinal"
	"github.com/JoshSwooft/go-pdf-spike/template"
	"github.com/brianvoe/gofakeit/v6"
)

type AppointmentDetails struct {
	Heading string
	Value   string
}

func main() {

	now := time.Now()

	parsed := now.Format("02 January 2006")
	parsedParts := strings.Split(parsed, " ")
	day, _ := strconv.Atoi(parsedParts[0])
	ordinalDay := ordinal.Ordinalize(day)
	fmt.Println(parsed)
	parsedParts[0] = ordinalDay
	parsed = strings.Join(parsedParts, " ")
	fmt.Println(parsed)
	// os.Exit(1)

	var fakeAddress template.LetterAddress
	gofakeit.Struct(&fakeAddress)

	pdfTemplate := &template.ReceiptTemplate{
		Template: template.Template{Pdf: pdf.NewMaroto(consts.Portrait, consts.A4)},
		LogoPath: "assets/logo.png",
		LetterAddress: template.LetterAddress{
			Line1:    fakeAddress.Line1,
			Line2:    "Lake Batz",
			City:     fakeAddress.City,
			Name:     fakeAddress.Name,
			Postcode: fakeAddress.Postcode,
		},
		AppointmentDate: time.Now(),
		ServiceName:     "Flu Vaccination",
		Location:        "Some Pharmacy, Manchester",
		Email:           gofakeit.Email(),
		LegalFootNote:   gofakeit.LoremIpsumSentence(60),
		PaymentDetails: template.PaymentDetails{
			Currency:    gofakeit.CurrencyShort(),
			ServiceName: "Flu",
			Fee:         "12.99",
			Discount:    "3.99",
			VATNumber:   "1234567",
		},
	}

	template.Generate(pdfTemplate)
	pdfTemplate.Save()

}
