package command

import "github.com/senghoo/captain/models"

type RepoUpdateCommand struct {
	RepoIdentify string
	Status int
}

const (
	Success = iota
	WorkspaceNotFound
	RepoNotFound
)


func NewRepoUpdateCommand(identify string) *RepoUpdateCommand {
	return &RepoUpdateCommand{
		RepoIdentify: identify,
	}
}

func (r *RepoPullCommand) Run(build *models.Build) {
	ws := build.Workspace()
	if ws == nil{
		r.Status = WorkspaceNotFound
		return
	}
	repos, err := ws.Repositories()
	if err != nil{
		
	}

	for repo := range k
}
