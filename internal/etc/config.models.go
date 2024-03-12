package etc

type Configuration struct {
	Db    db    `toml:"db"`
	Web   web   `toml:"web"`
	Mq    mq    `toml:"mq"`
	Redis redis `toml:"redis"`
}

type web struct {
	Listen string `toml:"listen"`
}

type redis struct {
	Enable   bool
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}
type mq struct {
	Conn string `toml:"conn"`
}
