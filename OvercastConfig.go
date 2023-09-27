package main

type OvercastConfig struct {
		LoginUrl string `yaml:"loginUrl" env:"OVERCAST_LOGINURL"  env-default:"https://overcast.fm/login"`
		DataUrl  string `yaml:"dataUrl"  env:"OVERCAST_DATAURL"   env-default:"https://overcast.fm/account/export_opml/extended"`
    Email    string `yaml:"email"    env:"OVERCAST_EMAIL"`
    Password string `yaml:"password" env:"OVERCAST_PASSWORD"`
}

