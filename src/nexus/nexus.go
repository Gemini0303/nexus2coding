package nexus

import (
	nexusrm "github.com/sonatype-nexus-community/gonexus/rm"
	"nexus2coding/src/utils"
)

type Nexus struct {
	// nexus admin username
	username string
	// nexus admin password
	password string
	// nexus access url
	endpoint string
}

type Repo struct {
	repoName string
	repoType string
	format   string
	url      string
}

func New(username string, password string, endpoint string) *Nexus {
	return &Nexus{
		username: username,
		password: password,
		endpoint: endpoint,
	}
}

func (n *Nexus) GetRepositories() (map[string]Repo, error) {
	rm, err := nexusrm.New(n.username, n.password, n.endpoint)
	if err != nil {
		return nil, err
	}

	repos, err := nexusrm.GetRepositories(rm)
	if err != nil {
		return nil, err
	}

	var result = make(map[string]Repo)
	for _, repo := range repos {
		result[repo.Name] = Repo{repoName: repo.Name, repoType: repo.Type, format: repo.Format, url: repo.URL}
	}

	return result, nil
}

func (n *Nexus) ChooseRepository() (string, error) {
	repositories, err := n.GetRepositories()
	if err != nil {
		return "", err
	}

	var items []utils.Needle
	for _, repository := range repositories {
		items = append(items, utils.Needle{Name: repository.repoName, Type: repository.repoType,
			Format: repository.format, URL: repository.url})
	}

	i := utils.SelectUI(utils.ExitOption(items), "Select Nexus Repository:")

	return items[i].Name, nil
}
