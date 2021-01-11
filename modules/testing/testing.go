package testing
import (
    "fmt"
    "os"
    "os/exec"
)
func TestCommand(testFilePath,path,inputfile string) {
	// os.Chdir("/mnt/e/slovink/terraformlibrary/terratest")

    os.Chdir(testFilePath)

    goExecutable, _ := exec.LookPath( "go" )
    cmdGoVer := &exec.Cmd {
        Path: goExecutable,
		Args: []string{ goExecutable, "test","-path",path,"-inputfile",inputfile },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    if err := cmdGoVer.Run(); err != nil {
        fmt.Println( "Error:", err );
    }
}
