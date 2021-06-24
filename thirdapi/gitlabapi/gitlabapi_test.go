package gitlabapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGiteeContents(t *testing.T) {
	u := fmt.Sprintf(ApiGitlabGetRepos, "SuperHeroJim", "gokins-test", ".gokins", "1065cd3f8791b97224a823c954a0ec98")
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(fmt.Errorf("Gitee Api :%v err : %v", u, err))
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Gitee ReadAll :%v err : %v", u, err))
		return
	}
	fmt.Println(string(all))
}

func TestGiteeCode(t *testing.T) {
	u := fmt.Sprintf("https://gitlab.com/login/oauth/authorize&client_id=%s", "102c5b2608655a5b7683")
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(fmt.Errorf("Gitee Api :%v err : %v", u, err))
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Gitee ReadAll :%v err : %v", u, err))
		return
	}
	fmt.Println(string(all))
}
