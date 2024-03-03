package models

type ToRsMessage struct {
	Text        string `json:"text"`
	Tip         string `json:"tip"`
	Name        string `json:"name"`
	NameMention string `json:"nameMention"`
	Lvlkz       string `json:"lvlkz"`
	Timekz      string `json:"timekz"`
	Mesid       string `json:"mesid"`
	Nameid      string `json:"nameid"`
	Guildid     string `json:"guildid"`
	Avatar      string `json:"avatar"`
}
