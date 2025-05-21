package types

type MainConfig struct {
	QueueD struct {
		API          ConfigAPI                     `json:"api" yaml:"api"`
		Notification map[string]ConfigNotification `json:"notification" yaml:"notification"`
		Log          ConfigLog                     `json:"log" yaml:"log"`
		Include      string                        `json:"include" yaml:"include"`
	} `json:"queued" yaml:"queued"`
}

type ConfigAPI struct {
	HTTPListen   string `json:"httpListen" yaml:"httpListen"`
	Cors         string `json:"cors" yaml:"cors"`
	AuthEnabled  bool   `json:"authEnabled" yaml:"authEnabled"`
	AuthAdmin    string `json:"authAdmin" yaml:"authAdmin"`
	AuthReadOnly string `json:"authReadOnly" yaml:"authReadOnly"`
	AccessToken  string `json:"accessToken" yaml:"accessToken"`
}

type ConfigNotification struct {
	Level    string `json:"level" yaml:"level"` // [debug,info,warning,error,panic,fatal]
	Telegram struct {
		BotToken string `json:"botToken" yaml:"botToken"`
		ChatID   string `json:"chatId" yaml:"chatId"`
	} `json:"telegram" yaml:"telegram"`
	Lark struct {
		WebhookURL string `json:"webhookUrl" yaml:"webhookUrl"`
		Secret     string `json:"secret" yaml:"secret"`
	} `json:"lark" yaml:"lark"`
	Slack struct {
		WebhookURL string `json:"webhookUrl" yaml:"webhookUrl"`
		Channel    string `json:"channel" yaml:"channel"`
		User       string `json:"user" yaml:"user"`
	} `json:"slack" yaml:"slack"`
}

type ConfigLog struct {
	Level    string `json:"level" yaml:"level"` // [debug,info,warning,error,panic,fatal]
	Location string `json:"location" yaml:"location"`
}
