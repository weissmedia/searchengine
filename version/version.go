package version

import (
	"fmt"
	"runtime"
)

// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version returns the main version number that is being run at the moment.
var Version string

// BuildDate returns the date the binary was built. This will be filled in by the compiler.
var BuildDate string

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

// FullVersion returns a string with all the versioning details.
func FullVersion() string {
	return fmt.Sprintf(
		"Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s",
		Version, GitCommit, BuildDate, GoVersion, OsArch,
	)
}
