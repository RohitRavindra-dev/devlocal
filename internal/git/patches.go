package git

import "fmt"

func checkPatchValidity(filePath string) error {
	checkApplyArgs := append(
		[]string{
			"apply",
			"--check",
		},
		filePath,
	)

	return executeGitCommand(checkApplyArgs)
}

func patchFile(filePath string) error {
	patchFileArgs := append(
		[]string{
			"apply",
		},
		filePath,
	)

	return executeGitCommand(patchFileArgs)
}

func ApplyPatches(filesToPatch []string) error {
	for _, filePath := range filesToPatch {
		if err := checkPatchValidity(filePath); err != nil {
			fmt.Printf("\t[Error] while trying to patch file %s : %s\n", filePath, err.Error())
			continue
		}
		if err := patchFile(filePath); err != nil {
			fmt.Printf("\t[Error] while trying to patch file %s : %s\n", filePath, err.Error())
		}
	}

	return nil
}
