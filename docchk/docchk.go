/*
 * File: docchk.go
 * Author: wdeqin
 * Time: Sun 10 Apr 2016 04:13:11 PM CST
 * Purpose: Doc check main entry
 */
package main

import (
	"fmt"
	"github.com/wdeqin/docchk/chkitem"
	"github.com/wdeqin/docchk/ls"
	"os"
	"strings"
)

func main() {
	args := os.Args
	argCnt := len(args)
	if argCnt < 3 || argCnt > 4 {
		printHelp()
		return
	}

	projectCode := args[1]
	if !checkProjectCode(projectCode) {
		fmt.Printf("Invalid project code: %s\n\n", projectCode)
		return
	}

	path := args[2]
	fileList, err := ls.GetFilenames(path)
	if err != nil {
		fmt.Printf("Read project directory failed:\n\t%s\n\n", err)
		return
	}

	if len(fileList) <= 0 {
		fmt.Printf("No files found in directory: %s\n\n", path)
		return
	}

	chkList, err := chkitem.GetChkList("./check_list.json")
	if err != nil {
		if argCnt == 4 {
			checkListFile := args[3]
			chkList, err = chkitem.GetChkList(checkListFile)
			if err != nil {
				fmt.Printf("Read check list file failed: %s\n\t%s\n\n",
					checkListFile, err)
				return
			}
		} else {
			fmt.Printf("Read default check list file ./check_list.json failed!\n\n")
			return
		}
	}

	if len(chkList) <= 0 {
		fmt.Printf("Find no check item, config check list file first!\n\n")
		return
	}

	chkProject := chkitem.ChkProject{}
	chkProject.ProjectCode = projectCode

	fmt.Printf("%-20s%-10s%50s\n", "Check Item", "OK", "File")
	fmt.Printf("================================================================================\n")
	for _, chkItem := range chkList {
		if chkItem.Status != "A" {
			continue
		}
		chkOk := false
		var chkOkFile string
		for _, fileName := range fileList {
			if chkitem.CheckItemFile(chkProject, chkItem, fileName) == nil {
				chkOk = true
				chkOkFile = fileName
				break
			}
		}
		if chkOk {
			fmt.Printf("%-20s%-10s%50s\n", chkItem.Name, "PASSED", chkOkFile)
		} else {
			fmt.Printf("%-20s%-10s%50s\n", chkItem.Name, "FAILED", "NONE")
		}
	}

	fmt.Printf("\n")
}

func printHelp() {
	help :=
		`Help:
    docchk <project_code> <dir> [check_list]
`
	fmt.Println(help)
}

func checkProjectCode(projectCode string) bool {
	projectCode = strings.TrimSpace(projectCode)
	bytes := []byte(projectCode)
	if len(bytes) != 8 {
		println("len")
		return false
	}

	c0 := bytes[0]
	if c0 != 'T' && c0 != 'P' && c0 != 'E' {
		return false
	}

	for i := 1; i < len(bytes); i += 1 {
		ci := bytes[i]
		if !(ci >= '0' && ci <= '9') {
			println("num_char " + string(ci))
			return false
		}
	}

	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
