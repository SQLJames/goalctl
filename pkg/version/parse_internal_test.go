package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVersion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName      string
		VersionNumber string
		Major         int64
		Minor         int64
		Patch         int64
	}{
		{
			"MajorOnly",
			"1",
			1,
			0,
			0,
		},
		{
			"Major.MinorOnly",
			"1.1",
			1,
			1,
			0,
		},
		{
			"Major.Minor.Patch",
			"1.2.3",
			1,
			2,
			3,
		},
		{
			"Major.Minor.Patch",
			"1.2.0",
			1,
			2,
			0,
		},
	}
	for _, testCase := range testCases {
		version, _ := parseVersion(testCase.VersionNumber)
		major, minor, patch := version.Major(), version.Minor(), version.Patch()
		assert.Equalf(t, major, testCase.Major, "TestName: %s - Found Major versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Major, major)
		assert.Equalf(t, minor, testCase.Minor, "TestName: %s - Found Minor versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Minor, minor)
		assert.Equalf(t, patch, testCase.Patch, "TestName: %s - Found Patch versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Patch, patch)
	}
}

func TestUpgradeAvailable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName      string
		VersionNumber string
		remoteVersion string
		Expected      bool
	}{
		{
			"Version is the same.",
			"1.0.0",
			"1.0.0",
			false,
		},
		{
			"Version is less than Remote Version.",
			"0.0.0",
			"1.0.0",
			true,
		},
		{
			"Version is greater than Remote Version.",
			"1.0.0",
			"0.0.0",
			false,
		},
	}
	for _, testCase := range testCases {
		buildVersion = testCase.VersionNumber
		available, _ := UpgradeAvailable(testCase.remoteVersion)
		assert.Equalf(t, available, testCase.Expected, "TestName: %s Expected: %t Got: %t", testCase.TestName, testCase.Expected, available)
	}
}
