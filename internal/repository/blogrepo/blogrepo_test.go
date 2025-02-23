package blogrepo_test

import (
	"testing"

	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo/cache"
	"github.com/stretchr/testify/assert"
)

func TestAddBlogPost(t *testing.T) {
	input := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	capacity := 1 // cache capacity
	repo := blogrepo.NewRepository(capacity)

	testCases := []struct {
		name        string
		output      uint64
		expectedErr error
	}{
		{
			name:        "Success",
			output:      1,
			expectedErr: nil,
		},
		{
			name:        "Failure",
			output:      0,
			expectedErr: cache.ErrCacheInMaxCap,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := repo.AddBlogPost(t.Context(), input)

			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.output, output)
		})
	}
}

func TestBlogPosts(t *testing.T) {
	capacity := 1 // cache capacity

	input := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	repo := blogrepo.NewRepository(capacity)
	_, err := repo.AddBlogPost(t.Context(), input)
	assert.Nil(t, err)

	testCases := []struct {
		name        string
		outputLen   int
		expectedErr error
	}{
		{
			name:        "Success",
			outputLen:   1,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			output := repo.BlogPosts(t.Context())
			len := len(output)
			assert.Equal(t, tc.outputLen, len)

			if len > 0 {
				out := output[0]
				assert.NotNil(t, out)
				assert.Equal(t, input.Title, out.Title)
				assert.Equal(t, input.Description, out.Description)
				assert.Equal(t, input.Body, out.Body)
			}
		})
	}
}

func TestBlogPost(t *testing.T) {
	capacity := 1 // cache capacity

	input := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	repo := blogrepo.NewRepository(capacity)
	_, err := repo.AddBlogPost(t.Context(), input)

	assert.Nil(t, err)

	testCases := []struct {
		name        string
		id          uint64
		expectedErr error
	}{
		{
			name:        "Success",
			id:          1,
			expectedErr: nil,
		},

		{
			name:        "Failure",
			id:          2,
			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := repo.BlogPost(t.Context(), tc.id)

			assert.ErrorIs(t, err, tc.expectedErr)
			if err == nil {
				assert.NotNil(t, output)
				assert.Equal(t, input.Title, output.Title)
			}
		})
	}
}

func TestDeleteBlogPost(t *testing.T) {
	bp := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	capacity := 1 // cache capacity
	repo := blogrepo.NewRepository(capacity)
	_, err := repo.AddBlogPost(t.Context(), bp)

	assert.Nil(t, err)

	testCases := []struct {
		name        string
		id          uint64
		output      uint64
		expectedErr error
	}{
		{
			name:        "Success",
			id:          1,
			output:      1,
			expectedErr: nil,
		},
		{
			name:        "Failure",
			id:          2,
			output:      0,
			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := repo.DeleteBlogPost(t.Context(), tc.id)

			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.output, output)
		})
	}
}

func TestUpdateBlogPost(t *testing.T) {
	bp := blogbus.AddBlogPost{
		Title:       "Title",
		Description: "Description",
		Body:        "Body",
	}

	capacity := 1 // cache capacity
	repo := blogrepo.NewRepository(capacity)
	_, err := repo.AddBlogPost(t.Context(), bp)

	assert.Nil(t, err)

	testCases := []struct {
		name        string
		id          uint64
		input       blogbus.UpdateBlogPost
		output      uint64
		expectedErr error
	}{
		{
			name: "Success",
			id:   1,
			input: blogbus.UpdateBlogPost{
				Title:       "Updated Title",
				Description: "Updated Description",
				Body:        "Updated Body",
			},
			output:      1,
			expectedErr: nil,
		},
		{
			name: "Failure",
			id:   2,
			input: blogbus.UpdateBlogPost{
				Title:       "Updated Title",
				Description: "Updated Description",
				Body:        "Updated Body",
			},
			output:      0,
			expectedErr: cache.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := repo.UpdateBlogPost(t.Context(), tc.id, tc.input)

			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.output, output)
		})
	}
}
