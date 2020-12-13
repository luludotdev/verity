package version

import (
	"runtime"
)

var (
	sha1ver = "unknown"
	gitTag  string
)

// IsDev returns true on dev builds
func IsDev() bool {
	return gitTag == ""
}

// AppVersion returns the current application version
// or "dev"
func AppVersion() string {
	if IsDev() {
		return "dev"
	}

	return gitTag
}

// CommitHash returns the current git commit hash
// or "unknown"
func CommitHash() string {
	return sha1ver
}

// GoVersion returns the current golang runtime version
func GoVersion() string {
	return runtime.Version()
}
