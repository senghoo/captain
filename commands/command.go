package command

import (
	"encoding/json"
	"fmt"

	"github.com/senghoo/captain/models"
)

type Command interface {
	Run(build *models.Build) string
	Clone() Command
}

var commandMap = make(map[string]Command)

func RegisterCommand(name string, c Command) {
	commandMap[name] = c
}

func NewCommand(name string, args CommandArgs, build *models.Build) (Command, error) {
	c, ok := commandMap[name]
	if !ok {
		return nil, fmt.Errorf("Command %s not found", name)
	}
	c = c.Clone()
	err := UpdateArgs(c, args, build.Context)
	return c, err
}

type Node struct {
	CommandName string
	CommandArgs CommandArgs
	SubNode     map[string]*Node
}

func (n *Node) Command(build *models.Build) (Command, error) {
	return NewCommand(n.CommandName, n.CommandArgs, build)
}

func ParseNode(raw []byte) (n *Node, err error) {
	n = new(Node)
	err = json.Unmarshal(raw, n)
	return
}

func RunNode(n *Node, build *models.Build) {
	logger := build.Logger()
	command, err := n.Command(build)
	if err != nil {
		logger.Printf("create command error %s", err)
		return
	}

	status := command.Run(build)

	subnode, ok := n.SubNode[status]
	if !ok {
		return
	}

	RunNode(subnode, build)
}
