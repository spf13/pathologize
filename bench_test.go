package pathologize

import "testing"

func BenchmarkClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clean("C:/Users/dir:e*c?t<o>r|y/CON..")
	}
}

func BenchmarkCleanPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CleanPath("C:/Users/dir:e*c?t<o>r|y/CON..")
	}
}

func BenchmarkRemoveInvalidCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeInvalidCharacters("C:/Users/dir:e*c?t<o>r|y/CON..")
	}
}

func BenchmarkRemoveReservedNames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeReservedNames("CON")
		removeReservedNames("file")
	}
}

func BenchmarkRemoveTrailing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeTrailing("foo.bar.. ")
	}
}
