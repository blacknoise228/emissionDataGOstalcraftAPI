package configs

type Config struct {
	Logs      Logs      `mapstructure:"logs"`
	API       API       `mapstructure:"api"`
	Stalcraft Stalcraft `mapstructure:"stalcraft"`
}
type Logs struct {
	LogLvl string `mapstructure:"loglevel"`
}
type API struct {
	AdminAPI AdminAPI `mapstructure:"admin"`
	BotAPI   BotAPI   `mapstructure:"tgbot"`
}
type AdminAPI struct {
	PortAdminAPI int `mapstructure:"port"`
}
type BotAPI struct {
	PortTgBot        int    `mapstructure:"port"`
	StalcraftTgToken string `mapstructure:"token"`
}
type Stalcraft struct {
	StalcraftID    string `mapstructure:"id"`
	StalcraftToken string `mapstructure:"token"`
}
