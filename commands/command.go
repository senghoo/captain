package command

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"

	"github.com/senghoo/captain/models"
)

type Command interface {
	Run(build *models.Build) string
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
	newCmd := copyCommand(c)
	err := UpdateArgs(newCmd, args, build.Context)
	return newCmd, err
}

func copyCommand(c Command) Command {
	val := reflect.ValueOf(c)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	return reflect.New(val.Type()).Interface().(Command)
}

type Node struct {
	CommandName string
	CommandArgs CommandArgs
	SubNode     map[string]*Node
}

func (n *Node) Command(build *models.Build) (Command, error) {
	return NewCommand(n.CommandName, n.CommandArgs, build)
}

func ParseNode(raw []byte, kind string) (n *Node, err error) {
	switch kind {
	case "json":
		return JSONParseNode(raw)
	case "yaml", "yml":
		return YMLParseNode(raw)
	default:
		return nil, fmt.Errorf("kind '%s' not defined", kind)
	}
}

func JSONParseNode(raw []byte) (n *Node, err error) {
	n = new(Node)
	err = json.Unmarshal(raw, n)
	return
}

// YMLParseNode ...
func YMLParseNode(raw []byte) (n *Node, err error) {
	n = new(Node)
	err = yaml.Unmarshal(raw, n)
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

func RunWorkflow(w *models.Workflow) error {
	node, err := ParseNode([]byte(w.Config), w.ConfigType)
	if err != nil {
		return err
	}

	ws, err := w.Workspace()
	if err != nil {
		return err
	}

	build, err := ws.NewBuild(w.Name)
	if err != nil {
		return err
	}

	RunNode(node, build)
	return nil
}
