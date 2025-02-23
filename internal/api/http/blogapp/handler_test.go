package blogapp_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/anazcodes/blog-crud-api/internal/api/http/blogapp"
	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
	mockblogbus "github.com/anazcodes/blog-crud-api/internal/mock/business/blogbus"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo/cache"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddBlogPost(t *testing.T) {
	port := ":3000"
	endpoint := "/api/blog-post"

	testCases := []struct {
		name           string
		input          blogbus.AddBlogPost
		expectedStatus int
		mockSetup      func(bus *mockblogbus.MockBusiness)
	}{
		{
			name: "Success",
			input: blogbus.AddBlogPost{
				Title:       "Title",
				Description: "Description",
				Body:        "Body",
			},
			expectedStatus: fiber.StatusCreated,
			mockSetup: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().AddBlogPost(gomock.Any(), gomock.Any()).Return(blogbus.ToID(1), nil).AnyTimes()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bus := mockblogbus.NewMockBusiness(ctrl)
			tc.mockSetup(bus)

			app := blogapp.NewApp(port, bus)
			fbr := app.Fiber()

			body, err := json.Marshal(tc.input)
			assert.Nil(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				endpoint,
				bytes.NewBuffer(body),
			)

			assert.Nil(t, err)

			req.Header.Set("Content-Type", "application/json")

			res, err := fbr.Test(req, -1)
			assert.Nil(t, err)

			resBody, err := io.ReadAll(res.Body)
			assert.Nil(t, err)

			defer res.Body.Close()

			log.Printf("Request URL: %s \n Response Body: %s", req.URL.String(), string(resBody))

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestBlogPosts(t *testing.T) {
	port := ":3000"
	endpoint := "/api/blog-post"

	testCases := []struct {
		name           string
		expectedStatus int
		setupExpect    func(bus *mockblogbus.MockBusiness)
	}{
		{
			name:           "Success",
			expectedStatus: fiber.StatusOK,
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().BlogPosts(gomock.Any()).Return([]blogbus.BlogPost{
					{
						ID:          1,
						Title:       "Title",
						Description: "Description",
						Body:        "Body",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				}).AnyTimes()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bus := mockblogbus.NewMockBusiness(ctrl)
			tc.setupExpect(bus)

			app := blogapp.NewApp(port, bus)
			fbr := app.Fiber()

			req, err := http.NewRequest(
				http.MethodGet,
				endpoint,
				nil,
			)

			assert.Nil(t, err)

			res, err := fbr.Test(req, -1)
			assert.Nil(t, err)

			resBody, err := io.ReadAll(res.Body)
			assert.Nil(t, err)

			defer res.Body.Close()

			assert.NotNil(t, resBody)

			log.Printf("Request URL: %s \n Response Body: %s", req.URL.String(), string(resBody))

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestBlogPost(t *testing.T) {
	port := ":3000"
	endpoint := "/api/blog-post/%d"

	testCases := []struct {
		name           string
		id             uint64
		expectedStatus int
		setupExpect    func(bus *mockblogbus.MockBusiness)
	}{
		{
			name: "Success",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().BlogPost(gomock.Any(), uint64(1)).Return(blogbus.BlogPost{
					ID:          1,
					Title:       "Title",
					Description: "Description",
					Body:        "Body",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil).AnyTimes()
			},
			id:             1,
			expectedStatus: fiber.StatusOK,
		},

		{
			name: "Not Found",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().BlogPost(gomock.Any(), uint64(1)).Return(blogbus.BlogPost{}, cache.ErrItemNotFound).AnyTimes()
			},
			id:             1,
			expectedStatus: fiber.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bus := mockblogbus.NewMockBusiness(ctrl)
			tc.setupExpect(bus)

			app := blogapp.NewApp(port, bus)
			fbr := app.Fiber()

			req, err := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf(endpoint, tc.id),
				nil,
			)

			assert.Nil(t, err)

			res, err := fbr.Test(req, -1)
			assert.Nil(t, err)

			resBody, err := io.ReadAll(res.Body)
			assert.Nil(t, err)

			defer res.Body.Close()

			assert.NotNil(t, resBody)

			log.Printf("Request URL: %s \n Response Body: %s", req.URL.String(), string(resBody))

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestDeleteBlogPost(t *testing.T) {
	t.Parallel()

	port := ":3000"
	endpoint := "/api/blog-post/%d"

	testCases := []struct {
		name           string
		setupExpect    func(bus *mockblogbus.MockBusiness)
		id             uint64
		expectedStatus int
	}{
		{
			name: "Success",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().DeleteBlogPost(gomock.Any(), uint64(1)).Return(blogbus.ToID(1), nil).AnyTimes()
			},
			id:             1,
			expectedStatus: fiber.StatusOK,
		},
		{
			name: "Not Found",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().DeleteBlogPost(gomock.Any(), uint64(1)).Return(nil, cache.ErrItemNotFound).AnyTimes()
			},
			id:             1,
			expectedStatus: fiber.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bus := mockblogbus.NewMockBusiness(ctrl)
			tc.setupExpect(bus)

			app := blogapp.NewApp(port, bus)
			fbr := app.Fiber()

			req, err := http.NewRequest(
				http.MethodDelete,
				fmt.Sprintf(endpoint, tc.id),
				nil,
			)

			assert.Nil(t, err)

			res, err := fbr.Test(req, -1)
			assert.Nil(t, err)

			resBody, err := io.ReadAll(res.Body)
			assert.Nil(t, err)

			defer res.Body.Close()

			assert.NotNil(t, resBody)

			log.Printf("Request URL: %s \n Response Body: %s", req.URL.String(), string(resBody))

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestUpdateBlogPost(t *testing.T) {
	t.Parallel()

	port := ":3000"
	endpoint := "/api/blog-post/%d"

	testCases := []struct {
		name           string
		setupExpect    func(bus *mockblogbus.MockBusiness)
		input          blogbus.UpdateBlogPost
		id             uint64
		expectedStatus int
	}{
		{
			name: "Success",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().UpdateBlogPost(gomock.Any(), uint64(1), blogbus.UpdateBlogPost{
					Title:       "Title",
					Description: "Description",
					Body:        "Body",
				}).Return(blogbus.ToID(1), nil).AnyTimes()
			},
			input: blogbus.UpdateBlogPost{
				Title:       "Title",
				Description: "Description",
				Body:        "Body",
			},
			id:             1,
			expectedStatus: fiber.StatusOK,
		},

		{
			name: "Not Found",
			setupExpect: func(bus *mockblogbus.MockBusiness) {
				bus.EXPECT().UpdateBlogPost(gomock.Any(), uint64(1), blogbus.UpdateBlogPost{
					Title:       "Title",
					Description: "Description",
					Body:        "Body",
				}).
					Return(nil, cache.ErrItemNotFound).
					AnyTimes()
			},
			input: blogbus.UpdateBlogPost{
				Title:       "Title",
				Description: "Description",
				Body:        "Body",
			},
			id:             1,
			expectedStatus: fiber.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bus := mockblogbus.NewMockBusiness(ctrl)
			tc.setupExpect(bus)

			app := blogapp.NewApp(port, bus)
			fbr := app.Fiber()

			body, err := json.Marshal(tc.input)
			assert.Nil(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				fmt.Sprintf(endpoint, tc.id),
				bytes.NewBuffer(body),
			)

			assert.Nil(t, err)

			req.Header.Set("Content-Type", "application/json")

			res, err := fbr.Test(req, -1)
			assert.Nil(t, err)

			resBody, err := io.ReadAll(res.Body)
			assert.Nil(t, err)

			defer res.Body.Close()

			assert.NotNil(t, resBody)

			assert.Equal(t, tc.expectedStatus, res.StatusCode, "Unexpected response status")
		})
	}
}
