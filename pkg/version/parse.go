package version

import (
	"fmt"

	"github.com/Masterminds/semver"
)

// parseVersion takes in a version string and then parses it to a Version object.
func parseVersion(ver string) (version *semver.Version, err error) {
	version, err = semver.NewVersion(ver)
	if err != nil {
		return nil, fmt.Errorf("parsing version: %w", err)
	}

	return version, nil
}

// UpgradeAvailable checks the version of the binary against a semantic version string.
// returns true if binary version is older than the semantic version.
func UpgradeAvailable(remoteVersion string) (upgradeAvailable bool, err error) {
	constraint, err := semver.NewConstraint(fmt.Sprintf("<%s", remoteVersion))
	if err != nil {
		return false, fmt.Errorf("semver NewConstraint: %w", err)
	}

	semVerCurrent, err := parseVersion(buildVersion)
	if err != nil {
		return false, fmt.Errorf("parseVersion: %w", err)
	}

	upgradeAvailable, _ = constraint.Validate(semVerCurrent)

	return upgradeAvailable, nil
}
