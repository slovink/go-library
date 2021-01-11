package terraform
import (
    "fmt"
    "os"
    "os/exec"
)

func PlanCommand(path string) {

    os.Chdir(path)
    goExecutable, _ := exec.LookPath( "terraform" )

    cmdGoVer := &exec.Cmd {
        Path: goExecutable,
		Args: []string{ goExecutable, "plan" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    if err := cmdGoVer.Run(); err != nil {
        fmt.Println( "Error:", err );
    }
}
