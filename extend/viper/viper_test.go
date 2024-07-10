package exViper

import (
	"testing"
)

func TestNewViper(t *testing.T) {
	type args struct {
		filePath string
	}
	//tests := []struct {
	//	name string
	//	args args
	//	want *SuperViper
	//}{
	//	// TODO: Add test cases.
	//	{
	//		args: args{
	//			filePath: "redis.toml",
	//		},
	//		want: NewViper("redis.toml"),
	//	},
	//}

	type Redis struct {
		Username string   `json:"username"`
		Password string   `json:"password"`
		Addrs    []string `json:"addrs"`
	}

	type GenerateObj struct {
		Redis Redis `json:"redis"`
	}

	if got := NewViper("redis.toml"); got != nil {
		err := got.ReadInConfig()
		t.Logf("redis ReadInConfig err:%v", err)
		objConf := GenerateObj{}
		err = got.Unmarshal(&objConf)
		t.Logf("redis Unmarshal:%+v err:%v", objConf, err)
		got.ToJson()
		got.ToXml()
		got.ToYaml()
		got.ToIni()
	}
}
