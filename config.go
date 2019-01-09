package gobot

type Config struct {
	Tuling Tuling `yaml:"tuling"`
}

type Rebot struct {
	Name string `yaml: "Name"`
	Key  string `yaml:"key"`
}
type Tuling struct {
	URL       string `yaml:"url"`
	RebotConf *Rebot
}

func Load() Config {
	cfg := Config{
		Tuling: Tuling{
			URL:       "http://www.tuling123.com/openapi/api",
			RebotConf: &Rebot{},
		},
	}
	cfg.Tuling.RebotConf.Name = "baoge"
	cfg.Tuling.RebotConf.Key = "0c11301074b544cb99c36bc6f9e23079"

	return cfg
}
