package pangu_test

import (
	"testing"

	"github.com/vinta/pangu"
)

type spaceTextCase struct {
	input, want string
}

func testSpaceText(t *testing.T, cases []spaceTextCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			if got := pangu.SpaceText(tc.input); got != tc.want {
				t.Errorf("SpaceText(%q)\n got: %q\nwant: %q", tc.input, got, tc.want)
			}
		})
	}
}

// testSpaceTextFails ports vitest's it.fails: it passes while any case still fails, and fails once they all pass, so a fixed FIXME gets noticed.
func testSpaceTextFails(t *testing.T, cases []spaceTextCase) {
	t.Helper()
	for _, tc := range cases {
		if pangu.SpaceText(tc.input) != tc.want {
			return
		}
	}
	t.Error("every case passes now: switch to testSpaceText and drop the FIXME")
}

func TestSpaceText(t *testing.T) {
	testSpaceText(t, []spaceTextCase{
		{"聽說Hadoop工程師睡不著的時候都會MapReduce羊", "聽說 Hadoop 工程師睡不著的時候都會 MapReduce 羊"},
		{"遇到了一個問題，決定用 thread 來解決，嗯，在現有我兩個問了題", "遇到了一個問題，決定用 thread 來解決，嗯，在現有我兩個問了題"},
	})
}

// Formatter contract: a second pass never changes the output, so format-then-check always passes
func TestSpaceTextIdempotent(t *testing.T) {
	for _, text := range []string{"\"字+\"", "\"字|\"", "你好\"字+\"世界", "多行\"字+\"\n下行\"字|\"", "聽說Hadoop工程師睡不著的時候都會MapReduce羊"} {
		t.Run(text, func(t *testing.T) {
			once := pangu.SpaceText(text)
			if twice := pangu.SpaceText(once); twice != once {
				t.Errorf("SpaceText(%q)\n got: %q\nwant: %q", once, twice, once)
			}
			if !pangu.HasProperSpacing(once) {
				t.Errorf("HasProperSpacing(%q) = false, want true", once)
			}
		})
	}
}

func TestHasProperSpacing(t *testing.T) {
	for text, want := range map[string]bool{
		"♫ 每條大街小巷，每個工程師的嘴裡，見面第一句話，就是不要在過年前 Deploy ♫": true,
		"♫每條大街小巷，每個工程師的嘴裡，見面第一句話，就是不要在過年前Deploy♫":    false,
	} {
		t.Run(text, func(t *testing.T) {
			if got := pangu.HasProperSpacing(text); got != want {
				t.Errorf("HasProperSpacing(%q) = %v, want %v", text, got, want)
			}
		})
	}
}
