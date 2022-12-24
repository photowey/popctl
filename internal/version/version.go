package version

const (
	versionNow = "0.1.0"
)

func Now() string {
	return "v" + versionNow
}
