package redirect_handler_test

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	redirect_handler "github.com/RozmiDan/url_shortener/internal/http-server/handlers/redirect"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockURLGetter struct {
	mock.Mock
}

func (m *MockURLGetter) GetURL(alias string) (string, error) {
	args := m.Called(alias)
	return args.Get(0).(string), args.Error(1)
}

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func TestGetHandler(t *testing.T) {

	testCases := []struct {
		name           string
		alias          string
		url            string
		mockErr        error
		expectedStatus int
	}{
		{
			name:           "success",
			alias:          "test_alias",
			url:            "https://google.com",
			mockErr:        nil,
			expectedStatus: http.StatusFound,
		},
		{
			name:           "not found alias",
			alias:          "unknown_alias",
			url:            "",
			mockErr:        errors.New("url not found"),
			expectedStatus: http.StatusOK,
		},
	}

	logger := slog.New(slog.NewTextHandler(bytes.NewBuffer(nil), nil))
	mockGetter := new(MockURLGetter)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGetter.ExpectedCalls = nil

			mockGetter.On("GetURL", tc.alias).Return(tc.url, tc.mockErr).Once()

			r := chi.NewRouter()
			r.Get("/{alias}", redirect_handler.RedirectHandlerConstructor(logger, mockGetter))

			ts := httptest.NewServer(r)
			defer ts.Close()

			client := http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			resp, err := client.Get(ts.URL + "/" + tc.alias)
			require.NoError(t, err)

			require.Equal(t, tc.expectedStatus, resp.StatusCode)

			if tc.expectedStatus == http.StatusFound {
				locationURL := resp.Header.Get("Location")
				require.Equal(t, tc.url, locationURL)
			}
		})
	}
}
