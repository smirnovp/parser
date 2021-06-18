package rusprofile

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"parser/grpcgen"
	"path"
	"testing"
)

type badReader struct {
	r io.Reader
}

func (br badReader) Read(b []byte) (int, error) {
	return 0, errors.New("io.Reader fail")
}

func TestSiteParserCreate(t *testing.T) {
	got := NewSiteParser()
	want := SiteParser{url: q}
	if *got != want {
		t.Errorf("\ngot:%#v\nwanted:%#v\n", *got, want)
	}
}

type testStruct struct {
	file    string
	inn     string
	kpp     string
	company string
	manager string
}

var tests = []testStruct{
	{file: "testpage.html",
		inn:     "7706561875",
		kpp:     "773001001",
		company: `ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ "ЭЛЕМЕНТ ЛИЗИНГ"`,
		manager: "Писаренко Андрей Витальевич",
	},
	{file: "testpage2.html",
		inn:     "7703735562",
		kpp:     "770301001",
		company: `ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ "ИНСТИТУТ АЛЛЕРГОЛОГИИ И КЛИНИЧЕСКОЙ ИММУНОЛОГИИ"`,
		manager: "Золина Людмила Васильевна",
	},
}

func TestParser(t *testing.T) {

	_, err := parseData(&badReader{}, "something")
	if err.Error() != "io.Reader fail" {
		t.Fatal(err)
	}
	for _, test := range tests {

		fname := path.Join("testdata", test.file)
		f, err := os.Open(fname)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		res, err := parseData(f, test.inn)
		if err != nil {
			t.Fatal("parseData() error:", err)
		}
		response := testStruct{file: test.file, inn: res.INN, kpp: res.KPP, company: res.Company, manager: res.Manager}

		if response != test {
			t.Errorf("\ngot:\n%#v,\nwanted:\n%#v\n", response, test)
		}
	}
}

func TestGetDataFromSite(t *testing.T) {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/testpage.html")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			t.Fatal(err)
		}
	}))
	sp := NewSiteParser()
	sp.url = ts.URL + "?query=%s"

	res, err := sp.GetDataFromSite(&grpcgen.ParserRequest{INN: "7706561875"})
	if err == nil {
		t.Error(err)
	}

	ts.Start()
	sp.url = ts.URL + "?query=%s"

	res, err = sp.GetDataFromSite(&grpcgen.ParserRequest{INN: "7706561875"})
	if err != nil {
		t.Error(err)
	}
	response := testStruct{inn: res.INN, kpp: res.KPP, company: res.Company, manager: res.Manager}
	expect := testStruct{inn: "7706561875", kpp: "773001001", company: `ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ "ЭЛЕМЕНТ ЛИЗИНГ"`, manager: "Писаренко Андрей Витальевич"}
	if response != expect {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v\n", response, expect)
	}
}
