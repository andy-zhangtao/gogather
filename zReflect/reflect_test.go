package zReflect

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/stretchr/testify/assert"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/23.
type User struct {
	ID               bson.ObjectId `json:"_id,omitempty" bson:"_id"`
	Name             string        `json:"name,omitempty" bson:"name"`
	Password         string        `json:"password,omitempty" bson:"password"`
	Projects         Project       `json:"projects,omitempty" bson:"projects"`
	Statis           UserStatis    `json:"statis,omitempty" bson:"statis"`
	CurrentAuthority string        `json:"currentAuthority,omitempty" bson:"currentauthority"`
	Resource struct {
		Cpu    float64 `json:"cpu" bson:"cpu"`
		Memory float64 `json:"memory" bson:"memory"`
	} `json:"resource omitempty" bson:"resource"`
}

type UserStatis struct {
	BuildSucc    int `json:"build_succ" bson:"buildsucc"`
	BuildFailed  int `json:"build_failed" bson:"buildfailed"`
	DeploySucc   int `json:"deploy_succ" bson:"deploysucc"`
	DeployFailed int `json:"deploy_failed" bson:"deployfailed"`
}

type Project struct {
	ID []string `json:"id" bson:"id"`
}

func TestReflectStructInfo(t *testing.T) {
	u := User{
		Name:     "andy@gmail.com",
		Password: "pbkdf2_sha256$12000$sYPLrXcUlw0r$lNZsiNWBHS/9DUNsYvKYtL1UjxUPv+IKaYJ5JMJtz9U=",
		Projects: Project{
			ID: []string{
				"iddd",
			},
		},
		Statis: UserStatis{
			DeployFailed: 4,
		},
		CurrentAuthority: "dev",
	}

	structInfo := ReflectStructInfo(u)

	assert.Equal(t, "andy@gmail.com", structInfo["name"])
	assert.Equal(t, "pbkdf2_sha256$12000$sYPLrXcUlw0r$lNZsiNWBHS/9DUNsYvKYtL1UjxUPv+IKaYJ5JMJtz9U=", structInfo["password"])
	assert.Equal(t, "iddd", structInfo["projects"].(Project).ID[0])
	assert.Equal(t, 4, structInfo["statis"].(UserStatis).DeployFailed)
	assert.Equal(t, "dev", structInfo["currentauthority"])
}
