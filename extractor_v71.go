package ogame

import (
	"bytes"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ExtractorV71 ...
type ExtractorV71 struct {
	ExtractorV7
}

// NewExtractorV71 ...
func NewExtractorV71() *ExtractorV71 {
	return &ExtractorV71{}
}

func (e ExtractorV71) ExtractResourcesDetails(pageHTML []byte) (out ResourcesDetails, err error) {
	return extractResourcesDetailsV71(pageHTML)
}

// ExtractEspionageReport ...
func (e ExtractorV71) ExtractEspionageReport(pageHTML []byte, location *time.Location) (EspionageReport, error) {
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(pageHTML))
	return e.ExtractEspionageReportFromDoc(doc, location)
}

// ExtractEspionageReportFromDoc ...
func (e ExtractorV71) ExtractEspionageReportFromDoc(doc *goquery.Document, location *time.Location) (EspionageReport, error) {
	return extractEspionageReportFromDocV71(doc, location)
}

// ExtractProduction extracts ships/defenses production from the shipyard page
func (e ExtractorV71) ExtractProduction(pageHTML []byte) ([]Quantifiable, int64, error) {
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(pageHTML))
	shipSumCountdown := e.ExtractOverviewShipSumCountdownFromBytes(pageHTML)
	production, err := e.ExtractProductionFromDoc(doc)
	return production, shipSumCountdown, err
}

// ExtractProductionFromDoc extracts ships/defenses production from the shipyard page
func (e ExtractorV71) ExtractProductionFromDoc(doc *goquery.Document) ([]Quantifiable, error) {
	return extractProductionFromDocV71(doc)
}
