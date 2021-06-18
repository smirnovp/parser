package rusprofile

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
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

	url, err := url.Parse(fmt.Sprintf(q, req.INN))
	if err != nil {
		return &grpcgen.ParserResponse{}, err
	}
	res, err := http.Get(url.String())
	if err != nil {
		return &grpcgen.ParserResponse{}, err
	}
	defer res.Body.Close()
	return parseData(res.Body, req.INN)
}

func parseData(r io.Reader, inn string) (*grpcgen.ParserResponse, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return &grpcgen.ParserResponse{}, err
	}

	INN := inn
	KPP := ""
	Company := ""
	Manager := ""

	rxINN := regexp.MustCompile(`clip_inn">(\d*)`)
	rxKPP := regexp.MustCompile(`clip_kpp">(\d*)`)
	rxCompany := regexp.MustCompile(`legalName">(.*?)<`)
	// rxManager := regexp.MustCompile(`(?s)<a href="/person.*?([А-Яа-я\s\w.,]+)</span></a>`)
	rxManager := regexp.MustCompile(`(?s)<span class="chief-title">Генеральный директор</span>.*?>([А-Яа-я]+\s[А-Яа-я]+\s[А-Яа-я]+)<`)

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
