package project

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pior/dad/pkg/config"
	"github.com/pior/dad/pkg/executor"
	"github.com/pior/dad/pkg/manifest"
)

type Project struct {
	HostingPlatform  string
	OrganisationName string
	RepositoryName   string
	Path             string
	Manifest         *manifest.Manifest
}

func NewFromId(id string, conf *config.Config) (p *Project, err error) {
	reGithubFull := regexp.MustCompile(`([^/]+)/([^/]+)`)

	if match := reGithubFull.FindStringSubmatch(id); match != nil {
		p = &Project{
			HostingPlatform:  "github.com",
			OrganisationName: match[1],
			RepositoryName:   match[2],
		}
	} else {
		err = fmt.Errorf("Unrecognized remote project: %s", id)
		return
	}

	p.Path = filepath.Join(conf.SourceDir, p.HostingPlatform, p.OrganisationName, p.RepositoryName)
	return
}

func (p *Project) FullName() string {
	return fmt.Sprintf("%s:%s/%s", p.HostingPlatform, p.OrganisationName, p.RepositoryName)
}

func (p *Project) GetRemoteUrl() (url string, err error) {
	if p.HostingPlatform == "github.com" {
		url = fmt.Sprintf("git@github.com:%s/%s.git", p.OrganisationName, p.RepositoryName)
		return
	}
	err = fmt.Errorf("Unknown project hosting platform: %s", p.HostingPlatform)
	return
}

func (p *Project) Exists() bool {
	if p.Path == "" {
		panic("Project path can't be null")
	}
	if _, err := os.Stat(p.Path); err == nil {
		return true
	}
	return false
}

func (p *Project) Clone() (err error) {
	err = os.MkdirAll(filepath.Dir(p.Path), 0755)
	if err != nil {
		return
	}

	url, err := p.GetRemoteUrl()
	if err != nil {
		return
	}
	err = executor.Run("git", "clone", url, p.Path)
	return
}