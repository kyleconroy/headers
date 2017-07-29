package headers

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// The HTTP Public-Key-Pins response header associates a specific cryptographic
// public key with a certain web server to decrease the risk of MITM attacks with
// forged certificates. If one or several keys are pinned and none of them are
// used by the server, the browser will not accept the response as legitimate, and
// will not display it.
//
// For more information, see https://developer.mozilla.org/en-US/docs/Web/HTTP/Public_Key_Pinning
type PublicKeyPins struct {
	Certificates      []*x509.Certificate
	MaxAge            time.Duration
	IncludeSubdomains bool
	ReportURL         *url.URL
	ReportOnly        bool
}

func (h PublicKeyPins) Name() string {
	return "Public-Key-Pins"
}

func (h PublicKeyPins) Value() string {
	pairs := []string{}
	for _, cert := range h.Certificates {
		digest := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
		pairs = append(pairs, fmt.Sprintf("pin-sha256=\"%s\"", base64.StdEncoding.EncodeToString(digest[:])))
	}
	pairs = append(pairs, fmt.Sprintf("max-age=%d", int(h.MaxAge.Seconds())))
	if h.IncludeSubdomains {
		pairs = append(pairs, "includeSubDomains")
	}
	if h.ReportURL != nil {
		pairs = append(pairs, "report-uri=\"%s\"", h.ReportURL.String())
	}
	return strings.Join(pairs, "; ")
}

func (h *PublicKeyPins) Parse(hdr string) error {
	return errors.New("NO")
}

var _ Header = &PublicKeyPins{}

// The HTTP Public-Key-Pins-Report-Only response header sends reports of pinning
// violation to the report-uri specified in the header but, unlike Public-Key-Pins
// still allows browsers to connect to the server if the pinning is violated.
type PublicKeyPinsReportOnly struct {
	*PublicKeyPins
}

func (h PublicKeyPinsReportOnly) Name() string {
	return "Public-Key-Pins-Report-Only"
}

var _ Header = &PublicKeyPinsReportOnly{}
