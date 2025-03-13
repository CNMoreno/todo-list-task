package http_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-task/internal/app"
	"todo-list-task/internal/domain"
	httpHandler "todo-list-task/internal/infrastructure/http"
	"todo-list-task/mocks"
)

const (
	routeUser = "/users"
)

type valuesTestCasesUser struct {
	name         string
	body         *domain.UserRequest
	token        string
	isError      bool
	isErrorBody  bool
	userResponse *domain.UserResponse
	err          error
	statusCode   int
}

var userRequest = &domain.UserRequest{
	Username: "cristianm",
	Password: "cristianm",
}

var userResponse = &domain.UserResponse{
	Token: "123DSASDEFRGR",
}

func TestUserHandler_RegisterUser(t *testing.T) {
	testCases := []valuesTestCasesUser{
		{
			name:         "Should create user when body is correct",
			body:         userRequest,
			userResponse: userResponse,
			statusCode:   http.StatusCreated,
			token:        "123DSASDEFRGR",
		},
		{
			name:        "Should throw an error when body is incorrect",
			statusCode:  http.StatusBadRequest,
			isError:     true,
			isErrorBody: true,
		},
		{
			name:       "Should throw an error when service return an error",
			statusCode: http.StatusInternalServerError,
			isError:    true,
			body:       userRequest,
			err:        assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo, handler, router := configurationUser()
			router.POST("/users", handler.RegisterUser)
			mockRepo.On("Create", mock.Anything).Return(tc.token, tc.err)
			bodyBytes, _ := json.Marshal(tc.body)
			req, _ := mockRequestEndPoint(tc.isErrorBody, "POST", routeUser, bytes.NewBuffer(bodyBytes))

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if tc.isError {
				assert.Equal(t, tc.statusCode, resp.Code)
			} else {
				var response *domain.UserResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tc.userResponse, response)
				assert.Equal(t, tc.statusCode, resp.Code)
			}

		})
	}
}

func TestUserHandler_LoginUser(t *testing.T) {
	testCases := []valuesTestCasesUser{
		{
			name:         "Should login user when body is correct",
			body:         userRequest,
			userResponse: userResponse,
			statusCode:   http.StatusOK,
			token:        "123DSASDEFRGR",
		},
		{
			name:        "Should throw an error when body is incorrect",
			statusCode:  http.StatusBadRequest,
			isError:     true,
			isErrorBody: true,
		},
		{
			name:       "Should throw an error when service return an error",
			statusCode: http.StatusInternalServerError,
			isError:    true,
			body:       userRequest,
			err:        assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo, handler, router := configurationUser()
			router.POST("/users/login", handler.LoginUser)
			mockRepo.On("Login", mock.Anything).Return(tc.token, tc.err)
			bodyBytes, _ := json.Marshal(tc.body)
			req, _ := mockRequestEndPoint(tc.isErrorBody, "POST", routeUser+"/login", bytes.NewBuffer(bodyBytes))

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if tc.isError {
				assert.Equal(t, tc.statusCode, resp.Code)
			} else {
				var response *domain.UserResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tc.userResponse, response)
				assert.Equal(t, tc.statusCode, resp.Code)
			}

		})
	}
}

func configurationUser() (*mocks.UserRepository, *httpHandler.UserHandler, *gin.Engine) {
	mockRepo := new(mocks.UserRepository)
	userService := app.NewUserService(mockRepo)
	handler := httpHandler.NewUserHandler(userService)

	router := gin.Default()

	return mockRepo, handler, router
}
