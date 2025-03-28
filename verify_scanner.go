package gmp

import (
	"encoding/xml"
)

type VerifyScannersCommand struct {
	XMLName   xml.Name `xml:"verify_scanner"`
	ScannerID string   `xml:"scanner_id,attr"`
}

type VerifyScannersResponse struct {
	XMLName    xml.Name `xml:"verify_scanner_response"`
	Status     string   `xml:"status,attr"`
	StatusText string   `xml:"status_text,attr"`
}
