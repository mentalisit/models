package models

type Timer struct {
	DsMesId  string `json:"ds_mes_id"`
	DsChatId string `json:"ds_chat_id"`
	TgMesId  string `json:"tg_mes_id"`
	TgChatId string `json:"tg_chat_id"`
	Timed    int    `json:"timed"`
}
