package comm

type Config struct {
	Server struct {
		Host      string   `yaml:"host"` //外网访问地址
		LoginKey  string   `yaml:"loginKey"`
		RunLimit  int      `yaml:"runLimit"`
		HbtpHost  string   `yaml:"hbtpHost"`
		Secret    string   `yaml:"secret"`
		Shells    []string `yaml:"shells"`
		DownToken string   `yaml:"DownToken"`
	} `yaml:"server"`
	Datasource struct {
		Driver string `yaml:"driver"`
		Url    string `yaml:"url"`
	} `yaml:"datasource"`
  Email struct {
    Host string `yaml:"host,omitempty"`
    Port int    `yaml:"port,omitempty"`
    User string `yaml:"user,omitempty"`
    Pass string `yaml:"pass,omitempty"`
    Sender string `yaml:"sender,omitempty"`
  } `yaml:"email,omitempty"`
}

const DATASOURCE_DRIVER_MYSQL = "mysql"
const DATASOURCE_DRIVER_POSTGRES = "postgres"
const DATASOURCE_DRIVER_SQLITE = "sqlite"
