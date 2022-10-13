package glex

import "testing"

func TestSplitSingle(t *testing.T) {
	words, err := SplitCommand("foo")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "foo" {
		t.Error("expected word 'foo', got", words[0])
	}
}

func TestSplitDouble(t *testing.T) {
	words, err := SplitCommand("foo bar")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 2 {
		t.Error("expected 2 words, got", len(words))
		return
	}

	if words[0] != "foo" {
		t.Error("expected word 'foo', got", words[0])
	}

	if words[1] != "bar" {
		t.Error("expected word 'bar', got", words[1])
	}
}

func TestSingleQuoted(t *testing.T) {
	words, err := SplitCommand("'foo bar'")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "foo bar" {
		t.Error("expected word 'foo bar', got", words[0])
	}
}

func TestDoubleQuoted(t *testing.T) {
	words, err := SplitCommand("\"foo bar\"")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "foo bar" {
		t.Error("expected word 'foo bar', got", words[0])
	}
}

func TestSingleQuotedEscaped(t *testing.T) {
	words, err := SplitCommand("'foo\\'bar'")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "foo'bar" {
		t.Error("expected word 'foo'bar', got", words[0])
	}
}

func TestDoubleQuotedEscaped(t *testing.T) {
	words, err := SplitCommand("\"foo\\\"bar\"")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "foo\"bar" {
		t.Error("expected word 'foo\"bar', got", words[0])
	}
}

func TestEscapes(t *testing.T) {
	words, err := SplitCommand(`\n\r\t\b\f\v\\\"\'`)
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 1 {
		t.Error("expected 1 word, got", len(words))
		return
	}

	if words[0] != "\n\r\t\b\f\v\\\"'" {
		t.Error("expected word '\n\r\t\b\f\v\\\"'\\'', got", words[0])
	}
}

func TestFlagEquals(t *testing.T) {
	words, err := SplitCommand("--foo=bar")
	if err != nil {
		t.Error(err)
		return
	}

	if len(words) != 2 {
		t.Error("expected 2 words, got", len(words))
		return
	}

	if words[0] != "--foo" {
		t.Error("expected word '--foo', got", words[0])
	}

	if words[1] != "bar" {
		t.Error("expected word 'bar', got", words[1])
	}
}

func TestUnclosedSingleQuote(t *testing.T) {
	_, err := SplitCommand("'foo")
	if err == nil {
		t.Error("expected error, got nil")
		return
	}
}

func TestUnclosedDoubleQuote(t *testing.T) {
	_, err := SplitCommand("\"foo")
	if err == nil {
		t.Error("expected error, got nil")
		return
	}
}

func TestUnclosedEscape(t *testing.T) {
	_, err := SplitCommand("\\")
	if err == nil {
		t.Error("expected error, got nil")
		return
	}
}

func TestComplexReal(t *testing.T) {
	commands := map[string][]string{
		"foo bar":                   {"foo", "bar"},
		"foo 'bar baz'":             {"foo", "bar baz"},
		"foo \"bar baz\"":           {"foo", "bar baz"},
		"foo 'bar\\'baz'":           {"foo", "bar'baz"},
		"foo \"bar\\\"baz\"":        {"foo", "bar\"baz"},
		"foo --bar=baz":             {"foo", "--bar", "baz"},
		"foo --bar=baz \"123\" \\v": {"foo", "--bar", "baz", "123", "\v"},
	}

	for cmd, expect := range commands {
		words, err := SplitCommand(cmd)
		if err != nil {
			t.Error(err)
			return
		}

		if len(words) != len(expect) {
			t.Error("expected", len(expect), "words, got", len(words))
			return
		}

		for i, word := range words {
			if word != expect[i] {
				t.Error("expected word", expect[i], "got", word)
				return
			}
		}
	}
}
