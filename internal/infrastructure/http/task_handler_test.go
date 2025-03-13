package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-list-task/internal/app"
	"todo-list-task/internal/domain"
	httpHandler "todo-list-task/internal/infrastructure/http"
	"todo-list-task/mocks"
)

const (
	route = "/tasks"
)

type valuesTestCases struct {
	name         string
	body         *domain.TaskRequest
	id           string
	isError      bool
	isErrorBody  bool
	userResponse *domain.Task
	err          error
	statusCode   int
}

var taskRequest = &domain.TaskRequest{
	Title:       "title",
	Description: "description",
}

var taskResponse = &domain.Task{
	Title:       "title",
	Description: "description",
	ID:          "12334556778",
}

func TestTaskHandler_RegisterTask(t *testing.T) {
	testCases := []valuesTestCases{
		{
			name:         "Create task",
			body:         taskRequest,
			userResponse: taskResponse,
			statusCode:   http.StatusCreated,
		},
		{
			name:        "Should throw an error when body is invalid",
			statusCode:  http.StatusBadRequest,
			isError:     true,
			isErrorBody: true,
		},
		{
			name:       "should return an error when repository return an error",
			err:        errors.New("some error"),
			body:       taskRequest,
			isError:    true,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo, handler, router := configuration()
			router.POST("/tasks", handler.RegisterTask)
			mockRepo.On("CreateTask", mock.Anything).Return(testCase.userResponse, testCase.err)
			bodyBytes, _ := json.Marshal(testCase.body)
			req, _ := mockRequestEndPoint(testCase.isErrorBody, "POST", route, bytes.NewBuffer(bodyBytes))

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if testCase.isError {
				assert.Equal(t, testCase.statusCode, resp.Code)
			} else {
				var response *domain.Task
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.userResponse, response)
				assert.Equal(t, testCase.statusCode, resp.Code)
			}
		})
	}
}

func TestTaskHandler_GetTaskByID(t *testing.T) {
	testCases := []valuesTestCases{
		{
			name:         "Get task by id",
			id:           "12334556778",
			userResponse: taskResponse,
			statusCode:   http.StatusOK,
		},
		{
			name:       "should return an error when repository return an error",
			err:        errors.New("some error"),
			id:         "12334556778",
			isError:    true,
			statusCode: http.StatusInternalServerError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo, handler, router := configuration()
			router.GET("/tasks/:id", handler.GetTaskByID)
			mockRepo.On("GetTask", mock.Anything).Return(testCase.userResponse, testCase.err)
			req, _ := mockRequestEndPoint(testCase.isErrorBody, "GET", route+"/"+testCase.id, nil)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if testCase.isError {
				assert.Equal(t, testCase.statusCode, resp.Code)
			} else {
				var response *domain.Task
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.userResponse, response)
				assert.Equal(t, testCase.statusCode, resp.Code)
			}
		})
	}
}

func TestTaskHandler_GetAllTasks(t *testing.T) {
	testCases := []valuesTestCases{
		{
			name: "Get all tasks",
			userResponse: &domain.Task{
				Title:       "title",
				Description: "description",
				ID:          "12334556778",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "should return an error when repository return an error",
			err:        errors.New("some error"),
			isError:    true,
			statusCode: http.StatusInternalServerError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo, handler, router := configuration()
			router.GET("/tasks", handler.GetAllTask)
			mockRepo.On("GetTasks").Return([]*domain.Task{testCase.userResponse}, testCase.err)
			req, _ := mockRequestEndPoint(testCase.isErrorBody, "GET", route, nil)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if testCase.isError {
				assert.Equal(t, testCase.statusCode, resp.Code)
			} else {
				var response []*domain.Task
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, []*domain.Task{testCase.userResponse}, response)
				assert.Equal(t, testCase.statusCode, resp.Code)
			}
		})
	}
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	testCases := []valuesTestCases{
		{
			name:         "Update task",
			id:           "12334556778",
			body:         taskRequest,
			userResponse: taskResponse,
			statusCode:   http.StatusOK,
		},
		{
			name:        "Should throw an error when body is invalid",
			id:          "12334556778",
			statusCode:  http.StatusBadRequest,
			isError:     true,
			isErrorBody: true,
		},
		{
			name:       "should return an error when repository return an error",
			id:         "12334556778",
			err:        errors.New("some error"),
			body:       taskRequest,
			isError:    true,
			statusCode: http.StatusInternalServerError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo, handler, router := configuration()
			router.PUT("/tasks/:id", handler.UpdateTask)
			mockRepo.On("UpdateTask", mock.Anything, mock.Anything).Return(testCase.userResponse, testCase.err)
			bodyBytes, _ := json.Marshal(testCase.body)
			req, _ := mockRequestEndPoint(testCase.isErrorBody, "PUT", route+"/"+testCase.id, bytes.NewBuffer(bodyBytes))

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if testCase.isError {
				assert.Equal(t, testCase.statusCode, resp.Code)
			} else {
				var response *domain.Task
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.userResponse, response)
				assert.Equal(t, testCase.statusCode, resp.Code)
			}
		})
	}
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	testCases := []valuesTestCases{
		{
			name:       "Delete task",
			id:         "12334556778",
			statusCode: http.StatusOK,
		},
		{
			name:       "should return an error when repository return an error",
			id:         "12334556778",
			err:        errors.New("some error"),
			isError:    true,
			statusCode: http.StatusInternalServerError,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo, handler, router := configuration()
			router.DELETE("/tasks/:id", handler.DeleteTask)
			mockRepo.On("DeleteTask", mock.Anything).Return(testCase.err)
			req, _ := mockRequestEndPoint(testCase.isErrorBody, "DELETE", route+"/"+testCase.id, nil)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if testCase.isError {
				assert.Equal(t, testCase.statusCode, resp.Code)
			} else {
				assert.Equal(t, testCase.statusCode, resp.Code)
			}
		})
	}
}

func mockRequestEndPoint(isError bool, method string, api string, body io.Reader) (*http.Request, error) {
	if isError {
		return http.NewRequest(method, api, strings.NewReader("Invalid Body"))
	}

	req, _ := http.NewRequest(method, api, body)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func configuration() (*mocks.TaskRepository, *httpHandler.TaskHandler, *gin.Engine) {
	mockRepo := new(mocks.TaskRepository)
	taskService := app.NewTaskService(mockRepo)
	handler := httpHandler.NewTaskHandler(taskService)

	router := gin.Default()

	router.Use(MockAuthMiddleware())
	return mockRepo, handler, router
}

func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
