package util

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/sirupsen/logrus"
)

func CloneRepo(path string, option *git.CloneOptions, ctx context.Context) (*git.Repository, error) {
	return git.PlainCloneContext(ctx,
		path,
		false,
		option,
	)
}

func CheckOutHash(repository *git.Repository, hash string) error {
	options := &git.CheckoutOptions{
		Force: true,
	}
	if len(hash)!=40{
		return CheckOutBranch(repository, hash)
	}else{
		options.Hash = plumbing.NewHash(hash)
		logrus.Debugf("CheckOut by hash, hash = %s", hash)
	}
	return CheckOut(repository, options)
}

func CheckOutBranch(repository *git.Repository, branch string) error {
	// refs/heads/<localBranchName>
	localBranchReferenceName := plumbing.NewBranchReferenceName(branch)
	// refs/remotes/origin/<remoteBranchName>
	remoteReferenceName := plumbing.NewRemoteReferenceName("origin", branch)

	err := repository.CreateBranch(&config.Branch{Name: branch,
		Remote: "origin", Merge: localBranchReferenceName, })
	if err!=nil{
		return err
	}
	newReference := plumbing.NewSymbolicReference(localBranchReferenceName , remoteReferenceName)
	err = repository.Storer.SetReference(newReference)
	if err != nil {
		return err
	}
	worktree, err := repository.Worktree()
	if err != nil {
		return err
	}
	err = worktree.Checkout(&git.CheckoutOptions{
		Create: false,
		Force: true,
		Keep: true,
		Branch: localBranchReferenceName,
	})
	return err
}

func CheckOut(repository *git.Repository, option *git.CheckoutOptions) error {
	worktree, err := repository.Worktree()
	if err != nil {
		return err
	}
	err = worktree.Checkout(option)
	if err != nil {
		return err
	}
	return nil
}

func GetLogsHash(repository *git.Repository, hash string) (object.CommitIter, error) {
	h := plumbing.NewHash(hash)
	options := &git.LogOptions{From: h}
	return GetLogs(repository, options)
}

func GetLogs(repository *git.Repository, option *git.LogOptions) (object.CommitIter, error) {
	return repository.Log(option)
}
