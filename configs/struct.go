package configs

type Config struct {
	Logs      Logs      `yaml:"logs"`
	API       API       `yaml:"api"`
	Stalcraft Stalcraft `yaml:"stalcraft"`
}
type Logs struct {
	LogLvl string `yaml:"loglevel"`
}
type API struct {
	AdminAPI AdminAPI `yaml:"admin"`
	BotAPI   BotAPI   `yaml:"bot"`
}
type AdminAPI struct {
	PortAdminAPI int `yaml:"port_adminapi"`
}
type BotAPI struct {
	PortTgBot int `yaml:"port_tgbot"`
}
type Stalcraft struct {
	StalcraftID      string `yaml:"stalcraft_id"`
	StalcraftTgToken string `yaml:"stalcraft_tg_token"`
	StalcraftToken   string `yaml:"stalcraft_token"`
}
