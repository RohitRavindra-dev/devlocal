package git

import "fmt"

func checkPatchValidity(patchFile string) error {
	checkApplyArgs := append(
		[]string{
			"apply",
			"--check",
		},
		patchFile,
	)

	return executeGitCommand(checkApplyArgs)
}

func ApplyPatches(filesToPatch []string) error {
	for _, filePath := range filesToPatch {
		if err := checkPatchValidity(filePath); err != nil {
			fmt.Printf("\t[Error] while trying to patch file %s : %s\n", filePath, err.Error())
		}
	}

	return nil
}
