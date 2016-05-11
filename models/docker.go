package models

import "time"

type DockerServer struct {
	ID       int64
	Name     string `xorm:"not null unique"`
	Endpoint string
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
	Deleted  time.Time `xorm:"deleted"`
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
