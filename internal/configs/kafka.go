package configs

type Kafka struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	GroupId string `yaml:"groupId"`
	Topic   string `yaml:"topic"`
}
