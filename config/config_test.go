package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	str, err := os.Getwd()
	if err != nil {
		t.Errorf("get current dir fail, err: %s", err)
	}

	path := str + "/conf_test.yml"
	conf, e := LoadConfig(path)
	if e != nil {
		t.Errorf("load config fail, err: %s", err)
	}

	t.Logf("config: %v", conf)
	if conf.Server.Host != "127.0.0.1" {
		t.Fatal("test fail")
	}

	if len(conf.Mysql.Master) != 2 {
		t.Fatalf("test fail, expected 2, real: %d", len(conf.Mysql.Master))
	}

	if len(conf.Mysql.Slave) != 2 {
		t.Fatalf("test fail, expected 2, real: %d", len(conf.Mysql.Slave))
	}

	if len(conf.Redis.Master) != 2 {
		t.Fatalf("test fail, expected 2, real: %d", len(conf.Redis.Master))
	}

	if len(conf.Redis.Slave) != 2 {
		t.Fatalf("test fail, expected 2, real: %d", len(conf.Redis.Slave))
	}

	if conf.Mysql.Master[0].Host != "127.0.0.1" {
		t.Fatal("test fail")
	}

	if !conf.Server.Package.OpenCheck {
		t.Fatal("test fail")
	}
}
