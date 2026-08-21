package config

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Title    string `yaml:"title" json:"title"`
	Desc     string `yaml:"desc" json:"desc"`
	Duration string `yaml:"duration" json:"duration"`
}
