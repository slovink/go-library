package terraform
import (
    "fmt"
    "os"
    "os/exec"
)

func ApplyCommand(path,inputfile string) {
	
    // os.Chdir("/mnt/e/slovink/terraformlibrary/terraform/vpc")
    os.Chdir(path)
    goExecutable, _ := exec.LookPath( "terraform" )

    cmdGoVer := &exec.Cmd {
        Path: goExecutable,
		Args: []string{ goExecutable, "apply" , "-input=false", "-auto-approve", "-var-file=./input.auto.tfvars.json" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    if err := cmdGoVer.Run(); err != nil {
        fmt.Println( "Error:", err );
    }
}
