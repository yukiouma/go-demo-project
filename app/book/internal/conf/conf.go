package conf

import (
	"frame/pkg/appmanage"
)

type ConfDB struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Customer struct {
	Addr string
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

func GenConf() (*ConfDB, *HttpConf, *GrpcConf, *Customer) {
	configs := appmanage.ReadConfig2Map("book")
	db := configs["db"].(map[string]interface{})
	http := configs["http"].(map[string]interface{})
	grpc := configs["grpc"].(map[string]interface{})
	customer := configs["customer"].(map[string]interface{})
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
		}, &Customer{
			Addr: customer["addr"].(string),
		}
}
