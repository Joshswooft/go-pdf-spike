package template

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type AppointmentDetails struct {
	Heading string
	Value   string
}

type LetterAddress struct {
	Name     string `fake:"{name}"`
	Line1    string `fake:"{street}"`
	Line2    string
	City     string `fake:"{city}"`
	Postcode string `fake:"{zip}"`
}

type PaymentDetails struct {
	ServiceName string
	Fee         string
	Discount    string
	VATNumber   string
	Currency    string
}

type ReceiptTemplate struct {
	*Template
	LogoPath        string
	LetterAddress   LetterAddress
	AppointmentDate time.Time
	ServiceName     string
	Location        string
	PaymentDetails  PaymentDetails
	Email           string
	LegalFootNote   string
}

func NewReceiptTemplate() *ReceiptTemplate {
	return &ReceiptTemplate{
		Template: &Template{Pdf: pdf.NewMaroto(consts.Portrait, consts.A4), ITmpl: nil},
	}
}

func (b *ReceiptTemplate) setPageMargins() {
	b.Template.Pdf.SetPageMargins(20, 10, 20)
}

func (b *ReceiptTemplate) buildHeading() {
	b.Template.Pdf.RegisterHeader(func() {
		b.Template.Pdf.Row(50, func() {
			b.Template.Pdf.ColSpace(10)
			b.Template.Pdf.Col(2, func() {
				err := b.Template.Pdf.FileImage(b.LogoPath, props.Rect{
					Percent: 100,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}
			})
		})
	})

	// Create address fields
	b.Template.Pdf.Row(50, func() {
		b.Template.Pdf.Col(4, func() {
			address := []string{b.LetterAddress.Name, b.LetterAddress.Line1, b.LetterAddress.Line2, b.LetterAddress.City, b.LetterAddress.Postcode}

			for _, line := range address {
				b.Template.Pdf.Row(5, func() {
					b.Template.Pdf.Col(12, func() {
						b.Template.Pdf.Text(line, props.Text{
							Align: consts.Left,
						})
					})
				})

			}

		})
	})
}

func (b *ReceiptTemplate) buildMainSection() {
	b.buildTitle()
	b.buildAppointmentSection()
	b.buildPaymentSection()
	b.buildQuestionSection()
}

func (b *ReceiptTemplate) buildFooter() {
	b.Template.Pdf.RegisterFooter(func() {
		b.Template.Pdf.Row(10, func() {
			b.Template.Pdf.Col(12, func() {
				b.Template.Pdf.Text(b.LegalFootNote, props.Text{
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

func (b *ReceiptTemplate) buildTitle() {
	b.Template.Pdf.Row(30, func() {
		b.Template.Pdf.Col(3, func() {
			b.Template.Pdf.Text("Appointment receipt", props.Text{
				Size:            20,
				VerticalPadding: 2,
			})
		})
		b.Template.Pdf.Col(9, func() {
			// TODO: replace with appointment date and format
			b.Template.Pdf.Text("3rd June 2021", props.Text{
				Align: consts.Right,
			})
		})
	})
}
func (b *ReceiptTemplate) buildAppointmentSection() {
	b.Template.Pdf.Row(1, func() {
		b.Template.Pdf.Col(4, func() {
			b.Template.Pdf.Text("Appointment details", props.Text{
				Style: consts.Bold,
				Size:  10,
			})
		})
	})
	b.Template.Pdf.Line(10)

	appointmentDetails := []AppointmentDetails{
		{
			Heading: "Date:",
			Value:   "3rd June 2021",
		},
		{
			Heading: "Service:",
			Value:   b.ServiceName,
		},
		{
			Heading: "Location:",
			Value:   b.Location,
		},
	}

	for _, detail := range appointmentDetails {

		b.Template.Pdf.Row(7, func() {
			b.Template.Pdf.Col(9, func() {
				b.Template.Pdf.Text(detail.Heading, props.Text{
					Style: consts.Bold,
					Size:  9,
				})
			})
			b.Template.Pdf.Col(3, func() {
				b.Template.Pdf.Text(detail.Value, props.Text{
					Size: 9,
				})
			})
		})
	}
}
func (b *ReceiptTemplate) buildPaymentSection() {
	b.Template.Pdf.Row(10, func() {})
	b.Template.Pdf.Row(1, func() {
		b.Template.Pdf.Col(4, func() {
			b.Template.Pdf.Text("Payment details", props.Text{
				Style: consts.Bold,
				Size:  10,
			})
		})
	})
	b.Template.Pdf.Line(10)

	b.Template.Pdf.Row(7, func() {
		b.Template.Pdf.Col(4, func() {
			b.Template.Pdf.Text(b.PaymentDetails.ServiceName, props.Text{
				Style: consts.Bold,
			})
		})
		b.Template.Pdf.Col(8, func() {
			// TODO: formatting of currency
			b.Template.Pdf.Text("£"+b.PaymentDetails.Fee, props.Text{
				Align: consts.Right,
				Style: consts.Bold,
			})
		})
	})
	// TODO: exclude this section if no discount exists
	b.Template.Pdf.Row(10, func() {
		b.Template.Pdf.Col(4, func() {
			b.Template.Pdf.Text("NHS exemption confirmed")
		})
		b.Template.Pdf.Col(8, func() {
			b.Template.Pdf.Text("-£"+b.PaymentDetails.Discount, props.Text{
				Align: consts.Right,
			})
		})
	})
	b.Template.Pdf.Row(10, func() {
		b.Template.Pdf.Col(4, func() {
			b.Template.Pdf.Text("Total paid:")
		})
		b.Template.Pdf.Col(8, func() {
			// TODO: fee - discount > 0 ?? 0
			b.Template.Pdf.Text("£0.00", props.Text{
				Align: consts.Right,
			})
		})
	})
	b.Template.Pdf.Row(7, func() {
		b.Template.Pdf.Col(12, func() {
			b.Template.Pdf.Text("VAT number: " + b.PaymentDetails.VATNumber)
		})

	})
}
func (b *ReceiptTemplate) buildQuestionSection() {
	b.Template.Pdf.Row(10, func() {})
	b.Template.Pdf.Row(10, func() {
		b.Template.Pdf.Col(12, func() {
			b.Template.Pdf.Text("Got a question?", props.Text{
				Size: 10,
			})
		})
	})
	b.Template.Pdf.Row(30, func() {
		b.Template.Pdf.Col(12, func() {
			b.Template.Pdf.Text("Email: "+b.Email, props.Text{
				Style: consts.Bold,
			})
		})
	})
}
