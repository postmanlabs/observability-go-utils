package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	testCases := map[string]struct {
		slice    []string
		init     int
		f        func(int, string) int
		expected int
	}{
		"nil slice": {
			slice:    nil,
			init:     10,
			f:        func(int, string) int { panic("shouldn't happen") },
			expected: 10,
		},
		"empty slice": {
			slice:    []string{},
			init:     10,
			f:        func(int, string) int { panic("shouldn't happen") },
			expected: 10,
		},
		"sum string lengths": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, s string) int {
				return accum + len(s)
			},
			expected: 10 + 1 + 6 + 3 + 5 + 5,
		},
	}

	for name, tc := range testCases {
		actual := Fold(tc.slice, tc.init, tc.f)
		assert.Equal(t, tc.expected, actual, name)
	}
}

func TestFoldWithErr(t *testing.T) {
	testCases := map[string]struct {
		slice       []string
		init        int
		f           func(int, string) (int, error)
		expected    int
		expectedErr bool
	}{
		"nil slice": {
			slice:    nil,
			init:     10,
			f:        func(int, string) (int, error) { panic("shouldn't happen") },
			expected: 10,
		},
		"empty slice": {
			slice:    []string{},
			init:     10,
			f:        func(int, string) (int, error) { panic("shouldn't happen") },
			expected: 10,
		},
		"sum string lengths": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, s string) (int, error) {
				return accum + len(s), nil
			},
			expected: 10 + 1 + 6 + 3 + 5 + 5,
		},
		"sum string lengths, early exit": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, s string) (int, error) {
				var err error
				if len(s) == 5 {
					err = fmt.Errorf("exit")
				}
				return accum + len(s), err
			},
			expected:    10 + 1 + 6 + 3 + 5,
			expectedErr: true,
		},
	}

	for name, tc := range testCases {
		actual, err := FoldWithErr(tc.slice, tc.init, tc.f)
		if tc.expectedErr {
			assert.Error(t, err, name)
		} else {
			assert.NoError(t, err, name)
		}
		assert.Equal(t, tc.expected, actual, name)
	}
}

func TestFoldIndex(t *testing.T) {
	testCases := map[string]struct {
		slice    []string
		init     int
		f        func(int, int, string) int
		expected int
	}{
		"nil slice": {
			slice:    nil,
			init:     10,
			f:        func(int, int, string) int { panic("shouldn't happen") },
			expected: 10,
		},
		"empty slice": {
			slice:    []string{},
			init:     10,
			f:        func(int, int, string) int { panic("shouldn't happen") },
			expected: 10,
		},
		"sum string lengths": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, idx int, s string) int {
				return accum + (idx+1)*len(s)
			},
			expected: 10 + 1*1 + 2*6 + 3*3 + 4*5 + 5*5,
		},
	}

	for name, tc := range testCases {
		actual := FoldIndex(tc.slice, tc.init, tc.f)
		assert.Equal(t, tc.expected, actual, name)
	}
}

func TestFoldIndexWithErr(t *testing.T) {
	testCases := map[string]struct {
		slice       []string
		init        int
		f           func(int, int, string) (int, error)
		expected    int
		expectedErr bool
	}{
		"nil slice": {
			slice: nil,
			init:  10,
			f: func(int, int, string) (int, error) {
				panic("shouldn't happen")
			},
			expected: 10,
		},
		"empty slice": {
			slice: []string{},
			init:  10,
			f: func(int, int, string) (int, error) {
				panic("shouldn't happen")
			},
			expected: 10,
		},
		"sum string lengths": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, idx int, s string) (int, error) {
				return accum + (idx+1)*len(s), nil
			},
			expected: 10 + 1*1 + 2*6 + 3*3 + 4*5 + 5*5,
		},
		"sum string lengths, early exit": {
			slice: []string{"1", "123456", "123", "12345", "12345"},
			init:  10,
			f: func(accum int, idx int, s string) (int, error) {
				var err error
				if len(s) == 5 {
					err = fmt.Errorf("exit")
				}
				return accum + (idx+1)*len(s), err
			},
			expected:    10 + 1*1 + 2*6 + 3*3 + 4*5,
			expectedErr: true,
		},
	}

	for name, tc := range testCases {
		actual, err := FoldIndexWithErr(tc.slice, tc.init, tc.f)
		if tc.expectedErr {
			assert.Error(t, err, name)
		} else {
			assert.NoError(t, err, name)
		}
		assert.Equal(t, tc.expected, actual, name)
	}
}
