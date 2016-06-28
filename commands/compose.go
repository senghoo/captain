package command

import "fmt"

func AutoBuildTree(repoID, dockerID int64, branch, name string) *Node {
	tree := `
{"CommandArgs": {"RepoID": %d}, "CommandName": "RepoUpdateCommand", "SubNode": {"success": {"CommandArgs": {"OutFile": "master.tar", "RepoID": %d, "Branch": "%s", "Format": "tar", "OutFile": "master.tar"}, "CommandName": "RepoArchiveCommand", "SubNode": {"success": {"CommandArgs": {"DockerID": %d, "Name": "%s", "File": "master.tar"}, "CommandName": "DockerBuildArchiveCommand", "SubNode": {}}}}}}
`
	tree = fmt.Sprintf(tree, repoID, repoID, branch, dockerID, name)

	node, err := ParseNode([]byte(tree))
	fmt.Printf("err: %s", err)
	return node

}
