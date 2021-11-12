package conf

import (
	"fmt"
	"frame/pkg/appmanage"
)

type ConfDB struct {
	Host     string
	Port     string
	User     string
	Password string
}

type GrpcConf struct {
	Host string
	Port string
}

func (c *GrpcConf) Addr() string {
	if len(c.Port) == 0 {
		c.Port = "8090"
	}
	if len(c.Host) == 0 {
		c.Host = "0.0.0.0"
	}
	return c.Host + ":" + c.Port
}

type HttpConf struct {
	Host string
	Port string
}

func (c *HttpConf) Addr() string {
	if len(c.Port) == 0 {
		c.Port = "8080"
	}
	if len(c.Host) == 0 {
		c.Host = "0.0.0.0"
	}
	return c.Host + ":" + c.Port
}

func GenConf() (*ConfDB, *HttpConf, *GrpcConf) {
	configs := appmanage.ReadConfig2Map("customer")
	fmt.Printf("%#v\n", configs)
	db := configs["db"].(map[string]interface{})
	http := configs["http"].(map[string]interface{})
	grpc := configs["grpc"].(map[string]interface{})
	return &ConfDB{
			Host:     db["host"].(string),
			Port:     db["port"].(string),
			User:     db["user"].(string),
			Password: db["password"].(string),
		}, &HttpConf{
			Host: http["host"].(string),
			Port: http["port"].(string),
		}, &GrpcConf{
			Host: grpc["host"].(string),
			Port: grpc["port"].(string),
		}
}
