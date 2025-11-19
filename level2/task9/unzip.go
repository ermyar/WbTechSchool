package task9

import (
	"errors"
	"strings"
)

var (
	ErrNoChar         = errors.New("no character before digits")
	ErrUnexpectedChar = errors.New("expected digit after \\")
)

func unzip(s string) (string, error) {
	var (
		sb        = &strings.Builder{}
		ch   rune = -1
		cnt  int  = -1
		mode bool
	)
	for _, r := range s {
		switch {
		case isDigit(r):
			if mode {
				writeNRunes(sb, cnt, ch)
				ch = r
				cnt = 0
				mode = false
				continue
			}
			if ch == -1 {
				return "", ErrNoChar
			}
			cnt *= 10
			cnt += int(r - '0')

		case r == '\\':
			if mode {
				return "", ErrUnexpectedChar
			}
			mode = true

		default:
			if mode {
				return "", ErrUnexpectedChar
			}
			if ch == -1 {
				ch = r
				cnt = 0
				continue
			}
			writeNRunes(sb, cnt, ch)
			ch = r
			cnt = 0
		}
	}
	if mode {
		return "", ErrUnexpectedChar
	}
	writeNRunes(sb, cnt, ch)

	return sb.String(), nil
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func writeNRunes(sb *strings.Builder, cnt int, r rune) {
	if cnt == -1 {
		return
	}
	if cnt == 0 {
		sb.WriteRune(r)
	}
	for range cnt {
		sb.WriteRune(r)
	}
}
