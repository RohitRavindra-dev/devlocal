package git

import "fmt"

func Generate(files []string) error {

	filesToDiff := sanitizeToTrackedFiles(files)

	if len(filesToDiff) == 0 {
		return fmt.Errorf("[Error] No files provided to generate patches that are being tracked by git")
	}

	return nil
}
