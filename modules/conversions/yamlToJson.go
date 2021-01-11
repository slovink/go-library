package conversions
import (
	"fmt"
    "os"
    "io/ioutil"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

func YamlToJson(path,inputfile string) {
	fmt.Println("creating file")
	filename, _ := filepath.Abs(inputfile)
	yamlFile, err := ioutil.ReadFile(filename)
    jsonFile, err := yaml.YAMLToJSON([]byte(yamlFile))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

    os.Chdir(path)
	ioutil.WriteFile("./input.auto.tfvars.json", jsonFile, 0777)
}