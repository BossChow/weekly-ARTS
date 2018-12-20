package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("abcabc", func(t *testing.T) {
		str := "abcabc"
		got := lengthOfLongestSubstring(str)
		want := 3

		if got != want {
			t.Errorf("got: %d want: %d", got, want)
		}
	})

	t.Run("abcabcbb", func(t *testing.T) {
		str := "abcabcbb"
		got := lengthOfLongestSubstring(str)
		want := 3

		if got != want {
			t.Errorf("got %d want: %d", got, want)
		}
	})
}
