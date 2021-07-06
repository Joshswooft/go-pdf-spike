package template

/**
This package uses an adapted pattern from the Template design pattern.

The idea is that we have a set of steps we need to perform but the implementation of those steps may be handled by
specific templates how they see fit. E.g. The receipt pdf template may have the logo on the right side whereas the GP
pdf may have a centered logo and some text explaining patient details in a table. - This is typical template pattern.

The adapted part comes from the need for common methods to be reflected in a "Base" class.
E.g. building the footer maybe the same for EVERY template barring one so it makes sense to NOT have to duplicate the same
code for each derived class and instead only the different template can define this method.

Golang doesn't have the normal inheritance support from more mature OOP languages so we attempt to accomplish this
by using composition.

*/
import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/pdf"
)

type iTemplate interface {
	setPageMargins()
	buildHeading()
	buildMainSection()
	buildFooter()
}

type Template struct {
	ITmpl iTemplate
	Pdf   pdf.Maroto
}

func (t *Template) Save() {
	err := t.Pdf.OutputFileAndClose("pdfs/test.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func (t *Template) setPageMargins() {
	fmt.Println("Base class calls setPageMargins()")
}

func (t *Template) buildHeading() {
	fmt.Println("Base class calls buildHeading()")
}

func (t *Template) buildMainSection() {
	fmt.Println("Base class calls buildMainSection()")
}

func (t *Template) buildFooter() {
	fmt.Println("Base class calls buildFooter()")
}

/**
Every template will have a header, main and footer
The implementation for this is handled by the specific template
but if the implementation is not handled in the derived class then the base method is used instead
*/
func (t *Template) Generate() {
	t.ITmpl.setPageMargins()
	t.ITmpl.buildHeading()
	t.ITmpl.buildMainSection()
	t.ITmpl.buildFooter()
}
