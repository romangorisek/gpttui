package text

import "testing"

func Test_ToMaxWidth(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		width    int
		expected string
	}{
		{
			name:     "Text shorter than width",
			text:     "hello",
			width:    10,
			expected: "hello",
		},
		{
			name:     "Text longer than width",
			text:     "hello world",
			width:    5,
			expected: "hello\n worl\nd",
		},
		{
			name:     "Text with newlines",
			text:     "hello\nworld",
			width:    10,
			expected: "hello\nworld",
		},
		{
			name:     "Text with newlines and longer than width",
			text:     "hello\nworld",
			width:    3,
			expected: "hel\nlo\nwor\nld",
		},
		{
			name:     "Width is zero",
			text:     "hello",
			width:    0,
			expected: "hello",
		},
		{
			name:     "Width is negative",
			text:     "hello",
			width:    -1,
			expected: "hello",
		},
		{
			name:     "Break after last space",
			text:     "This is a long sentence that should be wrapped.",
			width:    10,
			expected: "This is a\nlong\nsentence\nthat\nshould be\nwrapped.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToMaxWidth(tc.text, tc.width)
			if result != tc.expected {
				t.Errorf("expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
