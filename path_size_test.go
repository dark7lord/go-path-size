package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testdataDir  = "./testdata"
	testdataFile = "./testdata/one-piece.csv"
)

func TestGetPathSize_File(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		human    bool
		expected string
	}{
		{
			name:     "size without flags",
			path:     testdataFile,
			expected: "1681B",
		},
		{
			name:     "human readable size",
			path:     testdataFile,
			human:    true,
			expected: "1.6KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetPathSize(tt.path, false, tt.human, false)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{
			name:      "size without flags",
			path:      testdataDir,
			recursive: false,
			expected:  "2351B",
		},
		{
			name:      "size with recursive flag",
			path:      testdataDir,
			recursive: true,
			expected:  "4713B",
		},
		{
			name:      "size with recursive and hidden flags",
			path:      testdataDir,
			recursive: true,
			all:       true,
			expected:  "9252B",
		},
		{
			name:      "size with recursive hidden human flags",
			path:      testdataDir,
			recursive: true,
			human:     true,
			all:       true,
			expected:  "9.0KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetPathSize_Errors(t *testing.T) {
	t.Run("non-existent file", func(t *testing.T) {
		_, err := GetPathSize("./testdata/non-existent", false, false, false)
		require.Error(t, err)
	})
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name            string
		bytes           int64
		isHumanReadable bool
		expected        string
	}{
		{
			name:            "zero bytes no flags",
			bytes:           0,
			isHumanReadable: false,
			expected:        "0B",
		},
		{
			name:            "512 bytes with human flag",
			bytes:           512,
			isHumanReadable: true,
			expected:        "512B",
		},
		{
			name:            "2048 bytes with human flag",
			bytes:           2048,
			isHumanReadable: true,
			expected:        "2.0KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FormatSize(tt.bytes, tt.isHumanReadable)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestFormatPathSize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "file path",
			path:     testdataFile,
			expected: "1681B\t./testdata/one-piece.csv",
		},
		{
			name:     "directory path",
			path:     testdataDir,
			expected: "2351B\t./testdata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := GetPathSize(tt.path, false, false, false)
			actual := FormatPathSize(tt.path, size)

			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
