package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	save_handler "github.com/RozmiDan/url_shortener/internal/http-server/handlers/save"
	"github.com/RozmiDan/url_shortener/internal/usecase/random"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/require"
)

func Test_HappyPath(t *testing.T) {
	u := url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	}

	e := httpexpect.Default(t, u.String())

	e.POST("/url").WithJSON(save_handler.Request{
		URL:   gofakeit.URL(),
		Alias: random.NewAliasForURL(8),
	}).Expect().Status(200).JSON().Object().ContainsKey("alias")
}

func Test_Save_Redirect(t *testing.T) {
	testCases := []struct {
		name  string
		url   string
		alias string
		error string
	}{
		{
			name:  "valid test",
			url:   gofakeit.URL(),
			alias: gofakeit.Word(),
		},
		{
			name:  "invalid test",
			url:   "not-url",
			alias: gofakeit.Word(),
			error: "invalid request parameters",
		},
		{
			name: "empty alias",
			url:  gofakeit.URL(),
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			u := url.URL{
				Host:   "localhost:8080",
				Scheme: "http",
			}
			e := httpexpect.Default(t, u.String())

			response := e.POST("/url").WithJSON(save_handler.Request{
				URL:   tCase.url,
				Alias: tCase.alias,
			}).Expect().Status(200).JSON().Object()

			if tCase.error != "" {
				response.NotContainsKey("alias")
				response.Value("error").String().IsEqual(tCase.error)
				return
			}

			alias := tCase.alias

			if alias != "" {
				response.Value("alias").String().IsEqual(alias)
			} else {
				response.Value("alias").String().NotEmpty()
				alias = response.Value("alias").String().Raw()
			}

			testRedirect(t, alias, tCase.url)

		})
	}
}

func testRedirect(t *testing.T, alias string, urlToRedirect string) {
	u := url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
		Path:   alias,
	}

	redirToURL, err := GetRedirect(u.String())

	require.NoError(t, err)
	require.Equal(t, urlToRedirect, redirToURL)

}

func GetRedirect(url string) (string, error) {

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode != http.StatusFound {
		return "", fmt.Errorf("%s", "test.GetRedirect: invalid status code")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return resp.Header.Get("Location"), nil

}
