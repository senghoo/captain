package command

func AutoBuildComand(repoID, dockerID int64, branch, name string) Command {
	updateCommand := NewRepoUpdateCommand(repoID)
	archiveCommand := NewRepoArchiveCommand(repoID, "zip", branch, "repo.zip")
	buildComand := NewDockerBuildArchiveCommand(name, "repo.zip", dockerID)

	updateCommand.SetNext(archiveCommand)
	archiveCommand.SetNext(buildComand)
	return updateCommand
}
