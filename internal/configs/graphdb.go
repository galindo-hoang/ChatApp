package configs

type GraphDataBase struct {
	Type     string `yaml:"type"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
}
