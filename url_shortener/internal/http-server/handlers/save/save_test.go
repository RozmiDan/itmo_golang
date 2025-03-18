package save_handler_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	save_handler "github.com/RozmiDan/url_shortener/internal/http-server/handlers/save"
	"github.com/RozmiDan/url_shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock для интерфейса URLSaver
type MockURLSaver struct {
	mock.Mock
}

func (m *MockURLSaver) SaveURL(urlToSave string, alias string) (int64, error) {
	args := m.Called(urlToSave, alias)
	return args.Get(0).(int64), args.Error(1)
}

func TestSaveHandler(t *testing.T) {

	testCases := []struct {
		name           string
		alias          string
		url            string
		mockErr        error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "successful save with provided alias",
			alias:          "custom",
			url:            "https://example.com",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"OK","alias":"custom"}`,
		},
		{
			name:           "successful save with generated alias",
			url:            "https://google.com",
			alias:          "",
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `"status":"OK"`,
		},
		{
			name:           "validation error",
			url:            "not-url",
			alias:          "fasd",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"Error","error":"invalid request parameters"}`,
		},
		{
			name:           "URL already exists",
			url:            "https://example.com",
			alias:          "exists",
			mockErr:        storage.ErrURLExists,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"Error","error":"URL already exists"}`,
		},
	}

	// Отключаем вывод логов
	logger := slog.New(slog.NewTextHandler(bytes.NewBuffer(nil), nil))

	mockSaver := new(MockURLSaver)
	handler := save_handler.NewSaveHandler(logger, mockSaver)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSaver.ExpectedCalls = nil // Сбрасываем ожидания между кейсами

			if tc.mockErr == nil && tc.url != "not-url" {
				mockSaver.On("SaveURL", tc.url, mock.AnythingOfType("string")).Return(int64(1), nil)
			} else if tc.mockErr != nil {
				mockSaver.On("SaveURL", tc.url, tc.alias).Return(int64(0), tc.mockErr)
			}

			input := fmt.Sprintf(`{"url": "%s", "alias": "%s"}`, tc.url, tc.alias)

			req := httptest.NewRequest(http.MethodPost, "/save", bytes.NewReader([]byte(input)))
			rec := httptest.NewRecorder()

			handler(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), tc.expectedBody)

			mockSaver.AssertExpectations(t)
		})
	}
}
