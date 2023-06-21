package test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	//cmd := exec.Command("egrep", "'2023-03-23,(pdf),(个人|企业)'", "/tmp/track_file_histories.csv")
	cmd := "egrep '2023-03-23,(pdf),(个人|企业)' /tmp/track_file_histories.csv|awk -F, '{sum +=$5} END {print sum}'"
	output, err := exec.Command("bash", "-c", cmd).Output()
	//"/tmp/track_file_histories.csv", "|", "awk", "-F,", "'{sum +=$5} END {print sum}'")
	if err != nil {
		t.Logf("output: %s err: %s", string(output), err.Error())
	}
	t.Logf("output: %s", string(output))
}

func DownloadFileFromKs3(ks3Dir, localDir string) {
	startTime := time.Now()
	cmd := exec.Command("sh", "-c", fmt.Sprintf("/Users/shanglikang/Downloads/ks3util-mac-amd64 cp %s %s -r", ks3Dir, localDir))
	cmd.Stdout = os.Stdout
	fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DownloadFileFromKs3 cost: %dms", time.Since(startTime).Milliseconds())
}

func TestDownloadFileFromKs3(t *testing.T) {
	ks3Dir := "ks3://2cbigdata-datasupport/external/cloud/tmp_dwd_new_cloudfile_detail_i_d_test230327/"
	localDir := "/tmp/tmp_dwd_new_cloudfile_detail_i_d_test230327"

	DownloadFileFromKs3(ks3Dir, localDir)
}

func TestExecStdOut(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo hi;sleep 10;echo hello")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
