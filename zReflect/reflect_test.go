package zReflect

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/stretchr/testify/assert"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/23.
type User struct {
	ID               bson.ObjectId `json:"_id,omitempty" bson:"_id" bw:"_id"`
	Name             string        `json:"name,omitempty" bson:"name" bw:"name"`
	Password         string        `json:"password,omitempty" bson:"password" bw:"password"`
	Projects         Project       `json:"projects,omitempty" bson:"projects"`
	Statis           UserStatis    `json:"statis,omitempty" bson:"statis" bw:"statis"`
	CurrentAuthority string        `json:"currentAuthority,omitempty" bson:"currentauthority"`
	Resource         struct {
		Cpu    float64 `json:"cpu" bson:"cpu" bw:"cpu"`
		Memory float64 `json:"memory" bson:"memory"`
	} `json:"resource omitempty" bson:"resource"`
}

type UserStatis struct {
	BuildSucc    int `json:"build_succ" bson:"buildsucc"`
	BuildFailed  int `json:"build_failed" bson:"buildfailed"`
	DeploySucc   int `json:"deploy_succ" bson:"deploysucc"`
	DeployFailed int `json:"deploy_failed" bson:"deployfailed" bw:"deployfailed"`
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
	assert.Equal(t, "iddd", structInfo["projects.id"].([]string)[0])
	assert.Equal(t, 4, structInfo["statis.deployfailed"].(int))
	assert.Equal(t, "dev", structInfo["currentauthority"])
}

func TestReflectStructInfoWithTag(t *testing.T) {
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

	structInfo := ReflectStructInfoWithTag(u, true, "bw")

	assert.Equal(t, "andy@gmail.com", structInfo["name"])
	assert.Equal(t, "pbkdf2_sha256$12000$sYPLrXcUlw0r$lNZsiNWBHS/9DUNsYvKYtL1UjxUPv+IKaYJ5JMJtz9U=", structInfo["password"])
	assert.Equal(t, nil, structInfo["projects.id"])
	assert.Equal(t, 4, structInfo["statis.deployfailed"].(int))
	assert.Equal(t, nil, structInfo["currentauthority"])
}

func TestExtractValuePtrFromStruct(t *testing.T) {
	type User struct {
		Name string `json:"name" bw:"name"`
		Age  int    `json:"age" bw:"age"`
		Addr string `json:"addr"`
	}

	var fileds = []string{
		"name", "age", "addr",
	}
	u := new(User)
	vals, err := ExtractValuePtrFromStruct(u, fileds)

	assert.Nil(t, err)
	assert.EqualValues(t, 3, len(vals))
}
