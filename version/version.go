package version

// Version components
const (
	Maj = "1"
	Min = "2"
	Fix = "20"

	AppVer = 1
)

var (
	// Must be a string because scripts like dist.sh read this file.
	Version = "1.2.20-release"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
