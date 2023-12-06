package hitdir

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Golang_pj string

var create_go = &cobra.Command{
	Use:   "go",
	Short: "Create Go project",
	Run: func(cmd *cobra.Command, args []string) {
		create_project_go(Golang_pj)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Golang_pj, "golang", "g", "", "Create project in Go")
	rootCmd.AddCommand(create_go) // Регистрируем команду create_go
}

func create_project_go(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	crate_cmd_dir(name)
	create_logic_dir(name)
	create_dir_test(name)
	create_readmefile(name)
	fmt.Printf("Project %s created successfully!\n", name)
}

func crate_cmd_dir(name string) {
	path_cmd := name + "/cmd"
	name_path := path_cmd + "/" + name

	file_main := name_path + "/main.go"

	os.Mkdir(path_cmd, os.ModePerm)
	os.Mkdir(name_path, os.ModePerm)
	create_makefile(name, file_main)
	file, err := os.Create(file_main)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	fmt.Printf("Файл %s создан в катологе %s\n", file.Name(), path_cmd)

}

func create_logic_dir(name string) {
	logic := name + "/internal"
	logic_app := logic + "/" + name
	file_logik := logic_app + "/" + name + ".go"
	os.Mkdir(logic, os.ModePerm)
	os.Mkdir(logic_app, os.ModePerm)
	file, err := os.Create(file_logik)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	fmt.Printf("Файл %s создан в катологе %s\n", file.Name(), logic_app)
}

func create_dir_test(name string) {
	dir_test := name + "/test"
	dir_test_app := dir_test + "/" + name
	file_test_app := dir_test_app + "/test_" + name + ".go"
	os.Mkdir(dir_test, os.ModePerm)
	os.Mkdir(dir_test_app, os.ModePerm)
	file, err := os.Create(file_test_app)
	if err != nil {
		log.Println(file)
	}
	defer file.Close()
	fmt.Printf("Файл %s создан в катологе %s\n", file.Name(), dir_test_app)
}

func create_readmefile(name string) {
	path := name + "/Readme.md"
	data := "# " + name
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	file.Write([]byte(data))
	fmt.Printf("Файл %s создан в катологе %s\n", file.Name(), name)
}

func create_makefile(name string, path string) {
	text := fmt.Sprintf(".PHONY: run, build, fmt, install\nrun:\n\tgo run %s", path)
	path_make := name + "/Makefile"
	fmt.Println(text)
	file, err := os.Create(path_make)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	file.Write([]byte(text))
	fmt.Printf("Файл %s создан в катологе %s\n", file.Name(), path)
}
