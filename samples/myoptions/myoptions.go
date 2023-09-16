package myoptions

import myflag "github.com/dtweaveio/goframe/app/flag"

type MyOptions struct {
	MysqlOptions *MysqlOptions `json:"mysql" mapstructure:"mysql"`
}

func NewMyOptions() *MyOptions {
	return &MyOptions{
		MysqlOptions: NewMysqlOptions(),
	}
}

func (o *MyOptions) Flags() (fss myflag.NamedFlagSets) {
	o.MysqlOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}

func (o *MyOptions) Validate() []error {
	return nil
}

func (o *MyOptions) Complete() error {
	return nil
}
