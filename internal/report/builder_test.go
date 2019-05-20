package report

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/derailed/popeye/internal/issues"
	"github.com/stretchr/testify/assert"
)

func TestBuilderYAML(t *testing.T) {
	b, ta := NewBuilder(), NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(issues.Root, issues.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection("fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToYAML()

	assert.Nil(t, err)
	assert.Equal(t, reportYAML, s)
}

func TestBuilderJSON(t *testing.T) {
	b, ta := NewBuilder(), NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(issues.Root, issues.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection("fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToJSON()

	assert.Nil(t, err)
	assert.Equal(t, reportJSON, s)
}

func TestPrintSummary(t *testing.T) {
	b, ta := NewBuilder(), NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(issues.Root, issues.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection("fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := NewSanitizer(buff, 0, false)
	b.PrintSummary(san)

	assert.Equal(t, summary, buff.String())
}

func TestPrintHeader(t *testing.T) {
	b, ta := NewBuilder(), NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(issues.Root, issues.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection("fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := NewSanitizer(buff, 0, false)
	b.PrintHeader(san)

	assert.Equal(t, header, buff.String())
}

func TestPrintReport(t *testing.T) {
	b, ta := NewBuilder(), NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(issues.Root, issues.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection("fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := NewSanitizer(buff, 0, false)
	b.PrintReport(issues.OkLevel, san)

	assert.Equal(t, report, buff.String())
}

// ----------------------------------------------------------------------------
// Helpers...

var reportJSON = "{\"popeye\":{\"score\":100,\"grade\":\"A\",\"sanitizers\":[{\"sanitizer\":\"fred\",\"tally\":{\"ok\":1,\"info\":0,\"warning\":0,\"error\":0,\"score\":100},\"issues\":{\"blee\":[{\"group\":\"__root__\",\"level\":0,\"message\":\"Blah\"}]}}],\"errors\":[{}]}}"

var reportYAML = `popeye:
  score: 100
  grade: A
  sanitizers:
  - sanitizer: fred
    tally:
      ok: 1
      info: 0
      warning: 0
      error: 0
      score: 100
    issues:
      blee:
      - group: __root__
        level: 0
        message: Blah
  errors:
  - {}
`

var summary = "\n\x1b[38;5;75;mSUMMARY\x1b[0m\n\x1b[38;5;75;m" + strings.Repeat("┅", 100) + "\x1b[0m\nYour cluster score: 100 -- A\n                                                                                \x1b[38;5;82;mo          .-'-.     \x1b[0m\n                                                                                \x1b[38;5;82;m o     __| A    `\\  \x1b[0m\n                                                                                \x1b[38;5;82;m  o   `-,-`--._   `\\\x1b[0m\n                                                                                \x1b[38;5;82;m []  .->'  a     `|-'\x1b[0m\n                                                                                \x1b[38;5;82;m  `=/ (__/_       /  \x1b[0m\n                                                                                \x1b[38;5;82;m    \\_,    `    _)  \x1b[0m\n                                                                                \x1b[38;5;82;m       `----;  |     \x1b[0m\n\n"
var header = "\n\x1b[38;5;122;m ___     ___ _____   _____ \x1b[0m                                                     \x1b[38;5;75;mK          .-'-.     \x1b[0m\n\x1b[38;5;122;m| _ \\___| _ \\ __\\ \\ / / __|\x1b[0m                                                     \x1b[38;5;75;m 8     __|      `\\  \x1b[0m\n\x1b[38;5;122;m|  _/ _ \\  _/ _| \\ V /| _| \x1b[0m                                                     \x1b[38;5;75;m  s   `-,-`--._   `\\\x1b[0m\n\x1b[38;5;122;m|_| \\___/_| |___| |_| |___|\x1b[0m                                                     \x1b[38;5;75;m []  .->'  a     `|-'\x1b[0m\n\x1b[38;5;75;m  Biffs`em and Buffs`em!\x1b[0m                                                        \x1b[38;5;75;m  `=/ (__/_       /  \x1b[0m\n                                                                                \x1b[38;5;75;m    \\_,    `    _)  \x1b[0m\n                                                                                \x1b[38;5;75;m       `----;  |     \x1b[0m\n\n"
var report = "\n\x1b[38;5;75;mFREDS (1 SCANNED)\x1b[0m" + strings.Repeat(" ", 60) + "💥 0 😱 0 🔊 0 ✅ 1 \x1b[38;5;122;m100\x1b[0m٪\n\x1b[38;5;75;m" + strings.Repeat("┅", 100) + "\x1b[0m\n  · \x1b[38;5;155;mblee\x1b[0m\x1b[38;5;250;m" + strings.Repeat(".", 91) + "\x1b[0m✅\n    ✅ \x1b[38;5;155;mBlah.\x1b[0m\n\n"
