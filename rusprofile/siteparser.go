package rusprofile

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"parser/grpcgen"
	"regexp"
)

const q = "https://www.rusprofile.ru/search?query=%s&search_inactive=2"

// SiteParser ...
type SiteParser struct {
}

// NewSiteParser ...
func NewSiteParser() *SiteParser {
	return &SiteParser{}
}

// GetDataFromSite ...
func (s *SiteParser) GetDataFromSite(req *grpcgen.ParserRequest) (*grpcgen.ParserResponse, error) {

	url := fmt.Sprintf(q, req.INN)

	res, err := http.Get(url)
	if err != nil {
		return &grpcgen.ParserResponse{}, err
	}
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return &grpcgen.ParserResponse{}, err
	}

	INN := req.INN
	KPP := ""
	Company := ""
	Manager := ""

	rxINN := regexp.MustCompile(`clip_inn">(\d*)`)
	rxKPP := regexp.MustCompile(`clip_kpp">(\d*)`)
	rxCompany := regexp.MustCompile(`legalName">(.*?)<`)
	rxManager := regexp.MustCompile(`(?s)<a href="/person.*?([А-Яа-я\s\w.,]+)</span></a>`)

	if s := rxINN.FindStringSubmatch(string(b)); len(s) == 2 {
		INN = s[1]
	}
	if s := rxKPP.FindStringSubmatch(string(b)); len(s) == 2 {
		KPP = s[1]
	}
	if s := rxCompany.FindStringSubmatch(string(b)); len(s) == 2 {
		Company = html.UnescapeString(s[1])
	}
	if s := rxManager.FindStringSubmatch(string(b)); len(s) == 2 {
		Manager = s[1]
	}

	return &grpcgen.ParserResponse{
		INN:     INN,
		KPP:     KPP,
		Company: Company,
		Manager: Manager,
	}, nil
}
