package html

import (
	"github.com/elliotchance/gedcom/v39"
)

// PublishShowOptions contains the options for the publish command.
type PublishShowOptions struct {
	LivingVisibility LivingVisibility
	GoogleAnalyticsID string
}

// Publisher is the main object used to generate a website.
type Publisher struct {
	document *gedcom.Document
	options  *PublishShowOptions
}

func NewPublisher(document *gedcom.Document, options *PublishShowOptions) *Publisher {
	return &Publisher{
		document: document,
		options:  options,
	}
}

// ... rest of the file logic for publishing ...
