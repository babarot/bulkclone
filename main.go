package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/goware/urlx"
	git "gopkg.in/src-d/go-git.v4"
	yaml "gopkg.in/yaml.v2"
)

// Version is version info
const Version = "0.0.2"

// Config is
type Config struct {
	Repos []Repo `yaml:"repos"`
}

// Repo is
type Repo string

// String returns repo as type string
func (r Repo) String() string {
	return string(r)
}

// GetUsername returns...
func (r Repo) GetUsername() (string, error) {
	url, err := urlx.Parse(r.String())
	if err != nil {
		return "", err
	}
	s := strings.Split(url.Path, "/")
	return s[1], nil
}

// GetReponame returns...
func (r Repo) GetReponame() (string, error) {
	url, err := urlx.Parse(r.String())
	if err != nil {
		return "", err
	}
	s := strings.Split(url.Path, "/")
	return s[2], nil
}

func main() {
	// var contents []byte
	var input io.Reader
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		input = f
	} else {
		input = os.Stdin
	}
	var cfg Config
	err := yaml.NewDecoder(input).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	eg := errgroup.Group{}
	for _, repo := range cfg.Repos {
		base := filepath.Join(os.Getenv("HOME"), "src", "github.com")
		username, err := repo.GetUsername()
		if err != nil {
			continue
		}
		reponame, err := repo.GetReponame()
		if err != nil {
			continue
		}
		eg.Go(func() error {
			path := filepath.Join(base, username, reponame)
			_, err = git.PlainClone(path, false, &git.CloneOptions{
				URL:      string(repo),
				Progress: ioutil.Discard,
			})
			return err
		})
		fmt.Printf("clone: %s/%s\n", username, reponame)
	}

}
