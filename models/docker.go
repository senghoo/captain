package models

import (
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"
)

type DockerServer struct {
	ID       int64
	Name     string `xorm:"not null unique"`
	Endpoint string
	Created  time.Time          `xorm:"CREATED"`
	Updated  time.Time          `xorm:"UPDATED"`
	Deleted  time.Time          `xorm:"deleted"`
	_client  *docker.Client     `xorm:"-"`
	_info    *docker.DockerInfo `xorm:"-"`
}

func GetDockerServerByID(id int64) (*DockerServer, error) {
	s := new(DockerServer)
	has, err := x.Id(id).Get(s)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("Docker server id: %d not exist", id)
	}
	return s, nil
}

func NewDockerServer(name, endpoint string) *DockerServer {
	return &DockerServer{
		Name:     name,
		Endpoint: endpoint,
	}
}

func DockerServers() ([]*DockerServer, error) {
	var servers []*DockerServer
	return servers, x.Asc("id").Find(&servers)
}

func (d *DockerServer) Save() {
	if d.ID == 0 {
		// find same user
		cond := &DockerServer{
			Name: d.Name,
		}
		has, _ := x.Get(cond)
		if has {
			// got it
			x.Id(cond.ID).Update(d)
			return
		}
		x.Insert(d)
	} else {
		x.Id(d.ID).Update(d)
	}
}

func (d *DockerServer) Info() (*docker.DockerInfo, error) {
	if d._info == nil {
		client, err := d.client()
		if err != nil {
			return nil, err
		}

		info, err := client.Info()
		if err != nil {
			return nil, err
		}
		d._info = info
	}

	return d._info, nil
}

func (d *DockerServer) Build(opt docker.BuildImageOptions) (err error) {
	c, err := d.client()
	if err != nil {
		return
	}
	err = c.BuildImage(opt)
	return
}

func (d *DockerServer) client() (*docker.Client, error) {
	if d._client == nil {
		client, err := docker.NewClient(d.Endpoint)
		if err != nil {
			return nil, err
		}
		d._client = client
	}
	return d._client, nil
}
