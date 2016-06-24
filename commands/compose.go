package command

func AutoBuildCommand(repoID, dockerID int64, branch, name string) Command {
	updateCommand := NewRepoUpdateCommand(repoID)
	archiveCommand := NewRepoArchiveCommand(repoID, "tar", branch, "repo.tar")
	buildComand := NewDockerBuildArchiveCommand(name, "repo.tar", dockerID)

	updateCommand.SetNext(archiveCommand)
	archiveCommand.SetNext(buildComand)
	return updateCommand
}
