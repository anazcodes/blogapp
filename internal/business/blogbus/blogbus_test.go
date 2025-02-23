package blogbus_test

import (
	"testing"
	"time"

	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
	mockblogbus "github.com/anazcodes/blog-crud-api/internal/mock/business/blogbus"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo/cache"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddBlogPost(t *testing.T) {
	input := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	testCases := []struct {
		name        string
		setupExpect func(repo *mockblogbus.MockRepo)
		output      blogbus.ID
		expectedErr error
	}{

		{
			name: "Success",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().AddBlogPost(gomock.Any(), input).Return(uint64(1), nil).AnyTimes()
			},
			output:      blogbus.ToID(1),
			expectedErr: nil,
		},
		{
			name: "Failure",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().AddBlogPost(gomock.Any(), input).Return(uint64(0), cache.ErrCacheInMaxCap).AnyTimes()
			},
			output:      blogbus.ToID(0),
			expectedErr: cache.ErrCacheInMaxCap,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockblogbus.NewMockRepo(ctrl)
			tc.setupExpect(repo)
			bus := blogbus.NewBusiness(repo)

			output, err := bus.AddBlogPost(t.Context(), input)

			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.output.ID(), output.ID())
		})
	}
}

func TestBlogPosts(t *testing.T) {
	bp := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	testCases := []struct {
		name        string
		setupExpect func(repo *mockblogbus.MockRepo)
		output      []blogbus.BlogPost
		expectedLen int
		expectedErr error
	}{

		{
			name: "Success",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().BlogPosts(gomock.Any()).Return([]blogbus.BlogPost{
					{
						ID:          1,
						Title:       bp.Title,
						Description: bp.Description,
						Body:        bp.Body,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				}).AnyTimes()
			},
			output: []blogbus.BlogPost{
				{
					ID:          1,
					Title:       bp.Title,
					Description: bp.Description,
					Body:        bp.Body,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			expectedLen: 1,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mockblogbus.NewMockRepo(ctrl)

			tc.setupExpect(repo)
			bus := blogbus.NewBusiness(repo)

			output := bus.BlogPosts(t.Context())

			len := len(output)
			assert.Equal(t, tc.expectedLen, len)

			if len > 0 {
				assert.Equal(t, tc.output[0].Title, output[0].Title)
				assert.Equal(t, tc.output[0].Description, output[0].Description)
				assert.Equal(t, tc.output[0].Body, output[0].Body)
			}
		})
	}
}

func TestBlogPost(t *testing.T) {
	input := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	testCases := []struct {
		name        string
		setupExpect func(repo *mockblogbus.MockRepo)
		id          uint64
		output      blogbus.BlogPost
		expectedErr error
	}{

		{
			name: "Success",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().BlogPost(gomock.Any(), uint64(1)).Return(blogbus.BlogPost{
					ID:          1,
					Title:       input.Title,
					Description: input.Description,
					Body:        input.Body,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil).AnyTimes()
			},
			id: 1,
			output: blogbus.BlogPost{
				ID:          1,
				Title:       input.Title,
				Description: input.Description,
				Body:        input.Body,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			expectedErr: nil,
		},
		{
			name: "Failure",
			id:   2,
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().BlogPost(gomock.Any(), uint64(2)).Return(blogbus.BlogPost{}, cache.ErrItemNotFound).AnyTimes()
			},

			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockblogbus.NewMockRepo(ctrl)
			tc.setupExpect(repo)
			bus := blogbus.NewBusiness(repo)

			output, err := bus.BlogPost(t.Context(), tc.id)

			assert.ErrorIs(t, err, tc.expectedErr)
			if err == nil {
				assert.Equal(t, tc.output.ID, output.ID)
				assert.Equal(t, tc.output.Title, output.Title)
				assert.Equal(t, tc.output.Description, output.Description)
				assert.Equal(t, tc.output.Body, output.Body)
			}
		})
	}
}

func TestDeleteBlogPost(t *testing.T) {
	testCases := []struct {
		name        string
		setupExpect func(repo *mockblogbus.MockRepo)
		id          uint64
		output      blogbus.ID
		expectedErr error
	}{

		{
			name: "Success",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().DeleteBlogPost(gomock.Any(), uint64(1)).Return(uint64(1), nil).AnyTimes()
			},
			id:          1,
			output:      blogbus.ToID(1),
			expectedErr: nil,
		},
		{
			name: "Failure",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().DeleteBlogPost(gomock.Any(), uint64(2)).Return(uint64(0), cache.ErrItemNotFound).AnyTimes()
			},
			id:          2,
			output:      nil,
			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockblogbus.NewMockRepo(ctrl)
			tc.setupExpect(repo)
			bus := blogbus.NewBusiness(repo)

			output, err := bus.DeleteBlogPost(t.Context(), tc.id)
			assert.ErrorIs(t, err, tc.expectedErr)
			if err == nil {
				assert.Equal(t, tc.output.ID(), output.ID())
			}
		})
	}
}
func TestUpdateBlogPost(t *testing.T) {
	testCases := []struct {
		name        string
		setupExpect func(repo *mockblogbus.MockRepo)
		id          uint64
		input       blogbus.UpdateBlogPost
		output      blogbus.ID
		expectedErr error
	}{

		{
			name: "Success",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().UpdateBlogPost(gomock.Any(), uint64(1), blogbus.UpdateBlogPost{
					Title:       "Title",
					Description: "Description",
					Body:        "Body",
				}).Return(uint64(1), nil).AnyTimes()
			},
			input: blogbus.UpdateBlogPost{
				Title:       "Title",
				Description: "Description",
				Body:        "Body",
			},
			id:          1,
			output:      blogbus.ToID(1),
			expectedErr: nil,
		},
		{
			name: "Failure",
			setupExpect: func(repo *mockblogbus.MockRepo) {
				repo.EXPECT().UpdateBlogPost(gomock.Any(), uint64(1), blogbus.UpdateBlogPost{
					Title:       "Title",
					Description: "Description",
					Body:        "Body",
				}).Return(uint64(0), cache.ErrItemNotFound).AnyTimes()
			},
			input: blogbus.UpdateBlogPost{
				Title:       "Title",
				Description: "Description",
				Body:        "Body",
			},
			id:          1,
			output:      blogbus.ToID(0),
			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mockblogbus.NewMockRepo(ctrl)
			tc.setupExpect(repo)
			bus := blogbus.NewBusiness(repo)

			output, err := bus.UpdateBlogPost(t.Context(), tc.id, tc.input)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.output.ID(), output.ID())
		})
	}
}
