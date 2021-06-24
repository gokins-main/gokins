package giteeapi

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//
//func TestGiteeContents(t *testing.T) {
//
//	u := fmt.Sprintf(ApiGiteeGetRepositoryContent, "SuperHeroJim", "gokins-test", ".gokins", "1065cd3f8791b97224a823c954a0ec98")
//	resp, err := http.Get(u)
//	if err != nil {
//		fmt.Println(fmt.Errorf("Gitee Api :%v err : %v", u, err))
//		return
//	}
//	all, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(fmt.Errorf("Gitee ReadAll :%v err : %v", u, err))
//		return
//	}
//	fmt.Println(string(all))
//}

func TestGiteeCreateFile(t *testing.T) {
	newDefault := NewDefault()
	repos, err := newDefault.Repositories.GetRepos("abd17cf076f3a208cc359302dfaadc42", "all", Pushed, OrderDesc, "", 1, 20)
	if err != nil {
		fmt.Println("err", err)
	} else {
		all, err := ioutil.ReadAll(repos.Body)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println(string(all))
	}

}
