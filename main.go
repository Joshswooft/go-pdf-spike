package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type AppointmentDetails struct {
	Heading string
	Value   string
}

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildTitle(m)
	buildAppointmentSection(m)
	buildPaymentSection(m)
	buildQuestionSection(m)
	buildFooter(m)

	err := m.OutputFileAndClose("pdfs/test.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildTitle(m pdf.Maroto) {
	m.Row(30, func() {
		m.Col(3, func() {
			m.Text("Appointment receipt", props.Text{
				Size:            20,
				VerticalPadding: 2,
			})
		})
		m.Col(9, func() {
			m.Text("3rd June 2021", props.Text{
				Align: consts.Right,
			})
		})
	})
}

func buildAppointmentSection(m pdf.Maroto) {
	m.Row(1, func() {
		m.Col(4, func() {
			m.Text("Appointment details", props.Text{
				Style: consts.Bold,
				Size:  10,
			})
		})
	})
	m.Line(10)

	appointmentDetails := []AppointmentDetails{
		{
			Heading: "Date:",
			Value:   "3rd June 2021",
		},
		{
			Heading: "Service:",
			Value:   "Flu Vaccination",
		},
		{
			Heading: "Location:",
			Value:   "Some Pharmacy, Manchester",
		},
	}

	for _, detail := range appointmentDetails {

		m.Row(7, func() {
			m.Col(9, func() {
				m.Text(detail.Heading, props.Text{
					Style: consts.Bold,
					Size:  9,
				})
			})
			m.Col(3, func() {
				m.Text(detail.Value, props.Text{
					Size: 9,
				})
			})
		})
	}

}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.ColSpace(10)
			m.Col(2, func() {
				err := m.FileImage("assets/logo.png", props.Rect{
					Percent: 100,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}
			})
		})
	})

	// Create address fields
	m.Row(50, func() {
		m.Col(4, func() {
			address := gofakeit.Address()
			addressString := strings.ReplaceAll(address.Address, ", ", "\n")

			details := []string{gofakeit.Name()}
			addressLines := strings.Split(addressString, "\n")
			details = append(details, addressLines...)

			for _, line := range details {
				m.Row(5, func() {
					m.Col(12, func() {
						m.Text(line, props.Text{
							Align: consts.Left,
						})
					})
				})

			}

		})
	})
}

func buildPaymentSection(m pdf.Maroto) {
	// TODO: refactor this top part
	m.Row(10, func() {})
	m.Row(1, func() {
		m.Col(4, func() {
			m.Text("Payment details", props.Text{
				Style: consts.Bold,
				Size:  10,
			})
		})
	})
	m.Line(10)

	m.Row(7, func() {
		m.Col(4, func() {
			m.Text("Flu", props.Text{
				Style: consts.Bold,
			})
		})
		m.Col(8, func() {
			m.Text("£14.50", props.Text{
				Align: consts.Right,
				Style: consts.Bold,
			})
		})
	})
	m.Row(10, func() {
		m.Col(4, func() {
			m.Text("NHS exemption confirmed")
		})
		m.Col(8, func() {
			m.Text("-£14.50", props.Text{
				Align: consts.Right,
			})
		})
	})
	m.Row(10, func() {
		m.Col(4, func() {
			m.Text("Total paid:")
		})
		m.Col(8, func() {
			m.Text("£0.00", props.Text{
				Align: consts.Right,
			})
		})
	})
	m.Row(7, func() {
		m.Col(12, func() {
			m.Text("VAT number: 1234567")
		})

	})
}

func buildQuestionSection(m pdf.Maroto) {
	m.Row(10, func() {})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Got a question?", props.Text{
				Size: 10,
			})
		})
	})
	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("Email: hello@world.com", props.Text{
				Style: consts.Bold,
			})
		})
	})
}

func buildFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(gofakeit.LoremIpsumSentence(60), props.Text{
					Align:           consts.Center,
					Size:            8,
					VerticalPadding: 2,
					Color: color.Color{
						Red:   153,
						Green: 153,
						Blue:  153,
					},
				})
			})
		})
	})
}
