package chkitem

import (
	"testing"
)

func TestGetChkList(t *testing.T) {
	chkList, err := GetChkList("./check_list.json")
	if err != nil {
		t.Error(err)
		return
	}

	size := len(chkList)
	if size != 1 {
		t.Errorf("len(chkList) = %d, expect %d, %+v", size, 1, chkList)
		return
	}

	expectedStatus := "A"
	if chkList[0].Status != expectedStatus {
		t.Errorf("status = %s, expect %s", chkList[0].Status, expectedStatus)
		return
	}
}

func TestGetChkProject(t *testing.T) {
	project, err := GetChkProject("./check_list.json")
	if err != nil {
		t.Error(err)
		return
	}

	expectedProjectName := "T1559591"
	if project.ProjectName != expectedProjectName {
		t.Errorf("ProjectName = %s, expect %s", project.ProjectName,
			expectedProjectName)
	}
}
