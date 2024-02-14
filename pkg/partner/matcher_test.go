package partner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bi6o/aroundhome-challenge/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetMatchingPartners(ctx *gin.Context, floorMaterial string, addressLong, addressLat float64) ([]model.Partner, error) {
	args := m.Called(ctx, floorMaterial, addressLong, addressLat)
	return args.Get(0).([]model.Partner), args.Error(1)
}

func (m *MockRepo) Get(ctx *gin.Context, id uuid.UUID) (*model.Partner, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Partner), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestController_MatchSimplified(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	mockRepo := new(MockRepo)
	controller := &Controller{repo: mockRepo, logger: logger}

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    MatcherRequest
		prepareMock    func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid request",
			requestBody: MatcherRequest{
				FloorMaterial: "wood",
				AddressLong:   10,
				AddressLat:    20,
			},
			prepareMock: func() {
				mockRepo.On("GetMatchingPartners", mock.Anything, "wood", 10.0, 20.0).
					Return([]model.Partner{}, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "[]",
		},
		{
			name: "invalid floor material",
			requestBody: MatcherRequest{
				FloorMaterial: "invalid",
				AddressLong:   10,
				AddressLat:    20,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid floor material"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.prepareMock != nil {
				tc.prepareMock()
			}

			bodyBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest("POST", "/match", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			controller.Match(c)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
