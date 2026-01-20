package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

// IndividualPage is the page that shows detailed information about an
// individual.
type IndividualPage struct {
	document          *gedcom.Document
	individual        *gedcom.IndividualNode
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewIndividualPage(document *gedcom.Document, individual *gedcom.IndividualNode, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *IndividualPage {
	return &IndividualPage{
		document:          document,
		individual:        individual,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
	}
}

func (c *IndividualPage) WriteHTMLTo(w io.Writer) (int64, error) {
	// Ensure a name exists before accessing index 0.
	nameString := "Unknown"
	if len(c.individual.Names()) > 0 {
		nameString = c.individual.Names()[0].String()
	}

	// Cast the forced visibility to the specific type required by the functions.
	visibility := LivingVisibility(LivingVisibilityShow)

	individualName := NewIndividualName(c.individual, visibility, UnknownEmphasis)
	individualDates := NewIndividualDates(c.individual, visibility)

	return core.NewPage(
		nameString,
		core.NewComponents(
			NewPublishHeader(c.document, nameString, selectedExtraTab,
				c.options, c.indexLetters, c.placesMap),
			NewAllParentButtons(c.document, c.individual,
				visibility, c.placesMap),
			core.NewBigTitle(1, individualName),
			core.NewBigTitle(3, individualDates),
			core.NewHorizontalRuleRow(),
			core.NewRow(
				core.NewColumn(core.HalfRow, NewIndividualNameAndSex(c.individual)),
				core.NewColumn(core.HalfRow, NewIndividualAdditionalNames(c.individual)),
			),
			core.NewSpace(),
			NewIndividualEvents(c.document, c.individual,
				visibility, c.placesMap),
			core.NewSpace(),
			NewPartnersAndChildren(c.document, c.individual,
				visibility, c.placesMap),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
