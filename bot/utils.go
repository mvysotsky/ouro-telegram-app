package bot

type ChannelID string

func (i ChannelID) Recipient() string {
	return string(i)
}
