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
type CorporationConfig struct {
	CorpName    string `json:"corp_name"`
	DsChannel   string `json:"ds_channel"`
	TgChannel   string `json:"tg_channel"`
	Language    string `json:"language"`
	MesIdDsHelp string `json:"mes_id_ds_help"`
	MesIdTgHelp string `json:"mes_id_tg_help"`
	Forward     bool   `json:"forward"`
	DelMessage  bool   `json:"del_message"`
	GuildId     string `json:"guild_id"`
}
