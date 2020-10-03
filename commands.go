package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "tu!statuses" {
		msg := "User statuses:"
		for i, status := range userStatus {
			msg = msg + fmt.Sprintf("\n<@%v>: `%v`", i, status)
		}
		s.ChannelMessageSend(m.ChannelID, msg)
	}
}
