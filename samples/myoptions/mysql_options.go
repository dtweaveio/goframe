package myoptions

import (
	"fmt"

	"github.com/spf13/pflag"
)

type MysqlOptions struct {
	Host     string
	Username string
	Password string
	Database string
}

func NewMysqlOptions() *MysqlOptions {
	return &MysqlOptions{
		Host: "127.0.0.1",
	}
}

func (o *MysqlOptions) Validate() []error {
	fmt.Println("validate params")
	return nil
}

func (o *MysqlOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host,
		"mysql host")
	fs.StringVar(&o.Username, "mysql.username", o.Username,
		"mysql username")
	fs.StringVar(&o.Password, "mysql.password", o.Password,
		"mysql password")
}
