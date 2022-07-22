package fileopener

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Opens a documentation file in read mode and displays it
func OpenFile(name string) {
	_, err := os.Stat(".dox")
	if os.IsNotExist(err) {
		// error. No dox ball found
		fmt.Println(".dox file does not exist!\n First create some documentation with \"dox c <file_name>\"")
		os.Exit(0)
	}

	// create temp dir to untar .dox
	directory_name, failed_to_create_temp_dir := ioutil.TempDir("", "tmp_untar_dox")
	CheckError(failed_to_create_temp_dir)

	// untar (tar -xf .dox -C <temp_dir>)
	CheckError(exec.Command("tar", "-xf", ".dox", "-C", directory_name+"/").Run())

	// search for file in temp
	_, file_not_found := os.Stat(directory_name + "/" + name + ".md")
	if os.IsNotExist(file_not_found) {
		// error. No documentation entry for file ...
		fmt.Printf("Error. No documentation for file \"%s\"\n", name)
		os.Exit(0)
	}

	// output the file if found
	cmd := exec.Command("vi", directory_name+"/"+name+".md")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// checks error value
func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
}
