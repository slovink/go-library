package terraform
import (
    "fmt"
    "os"
    "os/exec"
)

func InitCommand(path string) {

    os.Chdir(path)

    goExecutable, _ := exec.LookPath( "terraform" )
    cmdGoVer := &exec.Cmd {
        Path: goExecutable,
		Args: []string{ goExecutable, "init" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    if err := cmdGoVer.Run(); err != nil {
        fmt.Println( "Error:", err );
    }
}
