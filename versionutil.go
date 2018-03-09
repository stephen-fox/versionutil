package versionutil

import (
	"errors"
	"path"
	"strings"
	"strconv"
	"unicode"
)

const (
	ErrorCouldNotFindVersion  = "Failed to find a version number in the filename"
	ErrorVersionStringInvalid = "The specified version string is invalid"
)

/*
Version represents an artifact's version.
*/
type Version struct {
	HasBuild    bool
	Major       int
	Minor       int
	Patch       int
	Build       int
}

// Long gets the full version as a string. This includes Major, Minor, Patch,
// and the build number.
// Example: 3.2.1.999
func (v Version) Long() string {
	if v.HasBuild {
		return v.Short() + "." + strconv.Itoa(v.Build)
	}

	return v.Short()
}

// Short gets the short version as a string. This includes Major, Minor, and
// Patch. This does not include the build number.
// Example: 3.2.1
func (v Version) Short() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
}

// BuildString gets the build number as a string (if it has one).
func (v Version) BuildToString() string {
	if v.HasBuild {
		return strconv.Itoa(v.Build)
	}

	return ""
}

// StringToVersion converts a version string (or a string containing a
// version) to a Version.
func StringToVersion(str string) (Version, error) {
	str = path.Base(str)

	if !strings.Contains(str, ".") {
		return Version{}, errors.New(ErrorVersionStringInvalid)
	}

	replacement := "-"

	str = strings.Replace(str, "_", replacement, -1)
	str = strings.Replace(str, " ", replacement, -1)

	return getVersion(strings.Split(str, replacement))
}

func getVersion(strParts []string) (Version, error) {
	version := Version{}
	var versionsStrs []string

	for _, part := range strParts {
		if len(part) == 0 || !strings.Contains(part, "."){
			continue
		}

		var buff []string

		for i, c := range part {
			if unicode.IsNumber(rune(c)) {
				buff = append(buff, string(c))
			} else if string(c) == "." && len(buff) > 0 {
				versionsStrs = append(versionsStrs, strings.Join(buff, ""))
				buff = []string{}
				continue
			}

			if i == len(part) - 1 && len(buff) > 0 {
				versionsStrs = append(versionsStrs, strings.Join(buff, ""))
			}
		}

		if len(versionsStrs) > 0 {
			break
		}
	}

	if len(versionsStrs) == 0 {
		return version, errors.New(ErrorCouldNotFindVersion)
	}

	for i, numberStr := range versionsStrs {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return version, err
		}

		switch i {
		case 3:
			version.HasBuild = true
			version.Build = number
		case 2:
			version.Patch = number
		case 1:
			version.Minor = number
		case 0:
			version.Major = number
		}
	}

	return version, nil
}
