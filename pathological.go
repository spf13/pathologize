//    Copyright 2024 Steve Francia & the Pathologize Authors
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package pathologize

// This package strives to create a filename that is safe to use on all modern operating/file systems.
// It is intentionally restrictive, and will not allow filenames that are too long,
// contain invalid characters, or are reserved names on any modern operating system, even if not the host OS.

// Wikipedia article on reserved filenames:
// https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words

import (
	"path/filepath"
	"regexp"
	"strings"
)

const (
	characterFilter      = `[\x00-\x1F\\/:*?"<>|@!]` // some fat systems don't like @ and ! in filenames
	dosReservedNames     = "CON PRN AUX NUL CLOCK$ CONFIG$ SCREEN$ $IDLE$ COM0 COM1 COM2 COM3 COM4 COM5 COM6 COM7 COM8 COM9 LPT0 LPT1 LPT2 LPT3 LPT4 LPT5 LPT6 LPT7 LPT8 LPT9"
	windowsReservedNames = "$Mft $MftMirr $LogFile $Volume $AttrDef $Bitmap $Boot $BadClus $Secure $Upcase $Extend $Quota $ObjId $Reparse"
	defaultName          = "file"
)

var (
	CharacterFilterRegex = regexp.MustCompile(characterFilter)
	maxLength            = 255
)

func CleanPath(path string) string {
	pathParts := strings.Split(path, string(filepath.Separator))
	for i, part := range pathParts {
		pathParts[i] = Clean(part)
	}

	return strings.Join(pathParts, string(filepath.Separator))
}

func IsClean(filename string) bool {
	return Clean(filename) == filename
}

func Clean(filename string) string {
	filename = removeInvalidCharacters(filename)
	filename = removeTrailing(filename)
	filename = removeLeadingSpaces(filename)
	filename = removeReservedNames(filename)
	filename = removeReservedWithExtension(filename)
	return filenameNotBlank(filename)
}

func filenameNotBlank(filename string) string {
	if filename == "" {
		return defaultName
	}
	return filename
}

func removeReservedWithExtension(filename string) string {
	basefilename := filenameWithoutExtension(filename)
	newfilename := removeReservedNames(basefilename)
	if basefilename != newfilename {
		return newfilename + filepath.Ext(filename)
	}
	return filename
}

// CleanFilename cleans a filename by removing all characters that are not allowed in filenames
func removeInvalidCharacters(filename string) string {
	filename = CharacterFilterRegex.ReplaceAllString(filename, "")
	return filename
}

func filenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func removeReservedNames(filename string) string {
	reservedNames := strings.Split(dosReservedNames+" "+windowsReservedNames, " ")
	for _, reservedName := range reservedNames {
		if strings.EqualFold(filename, reservedName) {
			return reservedName + "_"
		}
	}
	return filename
}

func removeTrailing(filename string) string {
	// Define the regular expression to match trailing dots and spaces
	re := regexp.MustCompile(`[.\s]+$`)

	// Replace all trailing dots and spaces with an empty string
	return re.ReplaceAllString(filename, "")
}

func removeLeadingSpaces(filename string) string {
	return strings.TrimSpace(filename)
}

func truncateFilename(filename string) string {
	if len(filename) > maxLength {
		return filename[:maxLength]
	}
	return filename
}
