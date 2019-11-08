package version

import "fmt"

const Version = "0.9.1"

var (
	Name      string
	GitCommit string

	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)
