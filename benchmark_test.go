package pathologize

import "testing"

func BenchmarkClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clean("CON..")
		Clean("foo/bar:baz*qux")
		Clean("normal_file_name.txt")
		Clean("$Mft.txt")
	}
}

func BenchmarkCleanPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CleanPath("C:/Users/dir:e*c?t<o>r|y/CON..")
	}
}
