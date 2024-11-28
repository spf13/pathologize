package pathologize

import (
	"testing"
)

func Test_truncateFilename(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{"empty", "", ""},
		{"short", "foo", "foo"},
		{"long",
			"foobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbaz",
			"foobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoobarbazfoo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateFilename(tt.file); got != tt.want {
				t.Errorf("truncateFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSurroundingSpaces(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{"empty", "", ""},
		{"short", "foo", "foo"},
		{"long", "foo bar baz", "foo bar baz"},
		{"leading", " foo", "foo"},
		{"trailing", "foo ", "foo"},
		{"both", " foo ", "foo"},
		{"multiple", "  foo  ", "foo"},
		{"multiple2", "  foo  bar  ", "foo  bar"},
		{"multiple3", "  foo  bar  baz  ", "foo  bar  baz"},
		{"multiple4", "  foo  bar  baz  qux  ", "foo  bar  baz  qux"},
		{"multiple5", "  foo  bar  baz  qux  quux  ", "foo  bar  baz  qux  quux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeadingSpaces(tt.file); got != tt.want {
				t.Errorf("removeSurroundingSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeTrailingDots(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", ""},
		{"short", "foo", "foo"},
		{"long", "foo.bar.baz", "foo.bar.baz"},
		{"trailing", "foo.", "foo"},
		{"multiple", "foo..", "foo"},
		{"multiple2", "foo...", "foo"},
		{"leading", ".foo", ".foo"},
		{"both", ".foo.", ".foo"},
		{"space dot space", ".foo . ", ".foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTrailing(tt.filename); got != tt.want {
				t.Errorf("removeTrailingDots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeReservedNames(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", ""},
		{"short", "foo", "foo"},
		{"long", "foo.bar.baz", "foo.bar.baz"},
		{"reserved", "CON", "CON_"},
		{"reserved2", "PRN", "PRN_"},
		{"reserved3", "AUX", "AUX_"},
		{"reserved4", "NUL", "NUL_"},
		{"reserved5", "COM1", "COM1_"},
		{"reserved6", "COM2", "COM2_"},
		{"reserved7", "COM3", "COM3_"},
		{"reserved8", "COM4", "COM4_"},
		{"reserved9", "COM5", "COM5_"},
		{"reserved10", "COM6", "COM6_"},
		{"reserved11", "COM7", "COM7_"},
		{"reserved12", "COM8", "COM8_"},
		{"reserved13", "COM9", "COM9_"},
		{"reserved14", "LPT1", "LPT1_"},
		{"reserved15", "LPT2", "LPT2_"},
		{"reserved16", "LPT3", "LPT3_"},
		{"reserved17", "LPT4", "LPT4_"},
		{"reserved18", "LPT5", "LPT5_"},
		{"reserved19", "LPT6", "LPT6_"},
		{"reserved20", "LPT7", "LPT7_"},
		{"reserved21", "LPT8", "LPT8_"},
		{"reserved22", "LPT9", "LPT9_"},
		{"windows reserved", "$Mft", "$Mft_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeReservedNames(tt.filename); got != tt.want {
				t.Errorf("removeReservedNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CleanFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", defaultName},
		{"trailing period", "CON.", "CON_"},
		{"trailing space", "PRN ", "PRN_"},
		{"leading space", " PRN", "PRN_"},
		{"space period", " AUX.", "AUX_"},
		{"period space", ".NUL ", ".NUL"},
		{"double period", "CON..", "CON_"},
		{"triple period", "CON...", "CON_"},
		{"triple period space", "CON... ", "CON_"},
		{"triple period space period", "CON... .", "CON_"},
		{"triple period space period space", "CON... . ", "CON_"},
		{"windows reserved", "$Mft", "$Mft_"},
		{"windows reserved period", "$Mft.", "$Mft_"},
		{"reserved with extension", "$Mft.txt", "$Mft_.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clean(tt.filename); got != tt.want {
				t.Errorf("removeReservedNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paths(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", "file"},
		{"short", "foo", "foo"},
		{"long", "foo.bar.baz", "foo.bar.baz"},
		{"several dirs", "foo/bar/baz", "foo/bar/baz"},
		{"several dirs2", "foo\\bar\\baz", "foobarbaz"},
		{"dirs with spaces", "foo bar/baz", "foo bar/baz"},
		{"dirs with bad chars", "foo*bar/baz", "foobar/baz"},
		{"example", "foo/bar:baz*qux", "foo/barbazqux"},
		{"readme", "C:/Users/dir:e*c?t<o>r|y/CON..", "C/Users/directory/CON_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanPath(tt.filename); got != tt.want {
				t.Errorf("paths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeInvalidCharacters(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", ""},
		{"short", "foo", "foo"},
		{"emoji", "👍", "👍"},
		{"long", "foo.bar.baz", "foo.bar.baz"},
		{"invalid", "foo/bar", "foobar"},
		{"invalidx01", "\x01foo/bar", "foobar"},
		{"invalid2", "foo\\bar", "foobar"},
		{"invalid3", "foo:bar", "foobar"},
		{"invalid4", "foo*bar", "foobar"},
		{"invalid5", "foo?bar", "foobar"},
		{"invalid6", "foo\"bar", "foobar"},
		{"invalid7", "foo<bar", "foobar"},
		{"invalid8", "foo>bar", "foobar"},
		{"invalid9", "foo|bar", "foobar"},
		{"invalid10", "foo/bar/baz", "foobarbaz"},
		{"invalid11", "foo\\bar\\baz", "foobarbaz"},
		{"invalid12", "foo:bar:baz", "foobarbaz"},
		{"invalid13", "foo*bar*baz", "foobarbaz"},
		{"invalid@", "foo@bar", "foobar"},
		{"invalid!", "foo!bar", "foobar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidCharacters(tt.filename); got != tt.want {
				t.Errorf("removeInvalidCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_characterFilter(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{"empty", "", ""},
		{"valid characters", "filename", "filename"},
		{"invalid characters", "file:name", "filename"},
		{"multiple invalid characters", "file:name*with?invalid|chars", "filenamewithinvalidchars"},
		{"leading invalid characters", ":*?<>|filename", "filename"},
		{"trailing invalid characters", "filename:*?<>|", "filename"},
		{"mixed valid and invalid characters", "file:name*with?valid|chars", "filenamewithvalidchars"},
		{"only invalid characters", ":*?<>|", ""},
		{"invalid characters with spaces", "file name*with?invalid|chars", "file namewithinvalidchars"},
		{"invalid characters with unicode", "file👍name*with?invalid|chars", "file👍namewithinvalidchars"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharacterFilterRegex.ReplaceAllString(tt.filename, ""); got != tt.want {
				t.Errorf("characterFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
