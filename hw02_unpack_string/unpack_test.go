package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "xa2bcd2e", expected: "xaabcdde"},
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "a0a0a0bbc0", expected: "bb"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "1"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func Test_returnMultiValue(t *testing.T) {
	tests := []struct {
		name        string
		r           rune
		c           rune
		expected    string
		needErr     bool
		expectedErr error
	}{
		{
			name:        "Положительный тест добавления символов",
			r:           't',
			c:           '4',
			expected:    "tttt",
			needErr:     false,
			expectedErr: nil,
		},
		{
			name:        "Положительный тест удаления символов",
			r:           't',
			c:           '0',
			expected:    "",
			needErr:     false,
			expectedErr: nil,
		},
		{
			name:        "Отрицательный тест",
			r:           't',
			c:           't',
			expected:    "",
			needErr:     true,
			expectedErr: errors.New("strconv.Atoi: parsing \"t\": invalid syntax"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := returnMultiValue(tc.r, tc.c)
			if tc.needErr {
				require.Error(t, err)
				require.EqualError(t, tc.expectedErr, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
