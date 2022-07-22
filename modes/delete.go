package modes

import (
	opener "dox/file_opener"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func DeleteDocumentation(file string) {
	// create temp dir to untar .dox
	directory_name, failed_to_create_temp_dir := ioutil.TempDir("", "tmp_untar_dox")
	opener.CheckError(failed_to_create_temp_dir)

	_, err := os.Stat(".dox")
	if !os.IsNotExist(err) {
		// untar (tar -xf .dox -C <temp_dir>)
		opener.CheckError(exec.Command("tar", "-xf", ".dox", "-C", directory_name+"/").Run())
	}

	// search for file in temp
	_, file_not_found := os.Stat(directory_name + "/" + file + ".md")
	if os.IsNotExist(file_not_found) {
		// error. File does not exist
		fmt.Printf("Error. There is no documentation for %s", file)
		os.Exit(0)
	} else {
		// remove the file
		os.Remove(directory_name + "/" + file + ".md")

		// get all files in temp dir
		var file_name_array []string
		files, read_files := ioutil.ReadDir(directory_name + "/")
		opener.CheckError(read_files)

		for _, file := range files {
			if !file.IsDir() {
				file_name_array = append(file_name_array, file.Name())
			}
		}

		// re-ball all files
		wd, wd_err := os.Getwd()
		opener.CheckError(wd_err)
		new_cmd := exec.Command("tar", "-cf", wd+"/.dox")
		new_cmd.Args = append(new_cmd.Args, file_name_array...)
		new_cmd.Dir = directory_name
		opener.CheckError(new_cmd.Run())
	}
}
