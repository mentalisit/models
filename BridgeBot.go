package models

type ToBridgeMessage struct {
	Text          string              `json:"text"`
	Sender        string              `json:"sender"`
	Tip           string              `json:"tip"`
	ChatId        string              `json:"chatId"`
	MesId         string              `json:"mesId"`
	GuildId       string              `json:"guildId"`
	GuildName     string              `json:"guildName"`
	TimestampUnix int64               `json:"timestampUnix"`
	FileUrl       string              `json:"fileUrl"`
	Avatar        string              `json:"avatar"`
	Reply         *BridgeMessageReply `json:"reply"`
	Config        *BridgeConfig       `json:"config"`
}
type BridgeConfig struct {
	Id                int              `json:"id"`
	NameRelay         string           `json:"nameRelay"`
	HostRelay         string           `json:"hostRelay"`
	Role              []string         `json:"role"`
	ChannelDs         []BridgeConfigDs `json:"channelDs"`
	ChannelTg         []BridgeConfigTg `json:"channelTg"`
	ForbiddenPrefixes []string         `json:"forbiddenPrefixes"`
	Prefix            string           `json:"prefix"`
}
type BridgeConfigDs struct {
	ChannelId       string            `json:"channel_id"`
	GuildId         string            `json:"guild_id"`
	CorpChannelName string            `json:"corp_channel_name"`
	AliasName       string            `json:"alias_name"`
	MappingRoles    map[string]string `json:"mapping_roles"`
}
type BridgeConfigTg struct {
	ChannelId       string            `json:"channel_id"`
	CorpChannelName string            `json:"corp_channel_name"`
	AliasName       string            `json:"alias_name"`
	MappingRoles    map[string]string `json:"mapping_roles"`
}
type BridgeMessageReply struct {
	TimeMessage int64  `json:"time_message"`
	Text        string `json:"text"`
	Avatar      string `json:"avatar"`
	UserName    string `json:"userName"`
	FileUrl     string `json:"fileUrl"`
}
type BridgeSendToMessenger struct {
	Text      string
	Sender    string
	ChannelId []string
	Avatar    string
	FileUrl   string
	Reply     *BridgeMessageReply
}
