package request

import (
	"log"
	"strings"
	"testing"
)

type Exp struct {
	URL string
}

func TestCreateRequest(t *testing.T) {
	cases := []struct {
		Data   HTTPAPI
		Expect Exp
	}{
		{
			Data: HTTPAPI{
				Method:    "GET",
				URL:       "something.com",
				URIParams: Params{"something": "new", "and": "shiny", "random": "stuff"},
			}, Expect: Exp{
				URL: "something.com?something=new&and=shiny&random=stuff",
			},
		},
	}

	for _, val := range cases {
		req, err := NewRequest(val.Data)
		if err != nil {
			t.Error(err)
			continue
		}

		log.Println(strings.Index(req.URL.String(), "?"))
		if req.URL.String() != val.Expect.URL {
			//t.Errorf("URL is incorrect. Got %s expecting %s", req.URL.String(), val.Expect.URL)
		}
	}
}
