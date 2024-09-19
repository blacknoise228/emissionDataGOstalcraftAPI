package configs

type Config struct {
	Logs      Logs      `mapstructure:"logs"`
	API       API       `mapstructure:"api"`
	Stalcraft Stalcraft `mapstructure:"stalcraft"`
	Database  Database  `mapstructure:"database"`
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
	PortTgBot int    `mapstructure:"port"`
	Token     string `mapstructure:"token"`
}
type Stalcraft struct {
	StalcraftID    string `mapstructure:"id"`
	StalcraftToken string `mapstructure:"token"`
}
type Database struct {
	DatabaseURL string `mapstructure:"databaseurl"`
}
