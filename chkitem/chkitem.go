/*
 * File: chkitem.go
 * Author: wdeqin
 * Time: Sun 10 Apr 2016 10:56:05 AM CST
 * Purpose: Doc check item persistence
 */
package chkitem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ChkItem struct {
	Status           string   `json:"status"`
	Name             string   `json:"name"`
	CheckProjectName string   `json:"check_project_name"`
	CheckProjectCode string   `jsong:"check_project_code"`
	Elements         []string `json:"elements"`
	Alternatives     []string `json:"alternatives"`
	Suffixs          []string `json:"suffixs"`
}

type chkSlice struct {
	ChkList []ChkItem `json:"check_list"`
}

func GetChkList(file string) ([]ChkItem, error) {
	bytes, err := getFileByte(file)
	if err != nil {
		return nil, err
	}
	slice := chkSlice{}
	err = json.Unmarshal(bytes, &slice)
	if err != nil {
		return nil, err
	}

	return slice.ChkList, nil
}

func getFileByte(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

type ChkProject struct {
	ProjectName string `json:"project_name"`
	ProjectCode string `json:"project_code"`
}

type chkProject struct {
	Project ChkProject `json:"check_project"`
}

func GetChkProject(file string) (*ChkProject, error) {
	bytes, err := getFileByte(file)
	if err != nil {
		return nil, err
	}
	project := chkProject{}
	err = json.Unmarshal(bytes, &project)
	if err != nil {
		return nil, err
	}

	return &project.Project, nil
}

func CheckItemFile(chkProject ChkProject, chkItem ChkItem, fileName string) error {
	if chkItem.Status == "Y" {
		if !strings.HasPrefix(fileName, chkProject.ProjectCode) {
			return fmt.Errorf("Check item [%s] failed: \n\t%s do not start with %s\n\n",
				chkItem.Name, fileName, chkProject.ProjectCode)
		}
	}

	for _, e := range chkItem.Elements {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}
		if !strings.Contains(fileName, e) {
			return fmt.Errorf("Check item [%s] failed: \n\t%s do not contain %s\n\n",
				chkItem.Name, fileName, e)
		}
	}

	suffixOk := false
	for _, s := range chkItem.Suffixs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if strings.HasSuffix(fileName, s) {
			suffixOk = true
			break
		}
	}

	if !suffixOk {
		return fmt.Errorf("Check item [%s] failed: \n\t%s do not end with any one of %s\n\n",
			chkItem.Name, fileName, chkItem.Suffixs)
	}

	return nil
}
