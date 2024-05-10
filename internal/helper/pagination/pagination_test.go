package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/model"
)

func TestSQLOffset(t *testing.T) {
	tests := []struct {
		name                    string
		page, perPage, expected uint64
	}{
		{
			name:     "offset first page",
			page:     1,
			perPage:  10,
			expected: 0,
		},
		{
			name:     "offset nth page",
			page:     5,
			perPage:  8,
			expected: 32,
		},
		{
			name:     "offset zero page",
			page:     0,
			perPage:  8,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SQLOffset(tt.page, tt.perPage)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTotalPage(t *testing.T) {
	tests := []struct {
		name                        string
		totalRow, perPage, expected uint64
	}{
		{
			name:     "0 row",
			totalRow: 0,
			perPage:  10,
			expected: 0,
		},
		{
			name:     "1 row",
			totalRow: 1,
			perPage:  10,
			expected: 1,
		},
		{
			name:     "1 page",
			totalRow: 10,
			perPage:  10,
			expected: 1,
		},
		{
			name:     "multiple page",
			totalRow: 34,
			perPage:  8,
			expected: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalPage(tt.totalRow, tt.perPage)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDefaultPagination(t *testing.T) {
	tests := []struct {
		name     string
		input    model.PaginationRequest
		expected model.PaginationRequest
	}{
		{
			name:     "Page and PerPage valid",
			input:    model.PaginationRequest{Page: 2, PerPage: 10},
			expected: model.PaginationRequest{Page: 2, PerPage: 10},
		},
		{
			name:     "Page not valid, PerPage valid",
			input:    model.PaginationRequest{Page: 0, PerPage: 10},
			expected: model.PaginationRequest{Page: 1, PerPage: 10},
		},
		{
			name:     "Page valid, PerPage not valid",
			input:    model.PaginationRequest{Page: 3, PerPage: 0},
			expected: model.PaginationRequest{Page: 3, PerPage: _const.DefaultPerPage},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DefaultPagination(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
