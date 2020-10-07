package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func newChannel(s *discordgo.Session, m *discordgo.ChannelCreate) {
	if strings.HasPrefix(m.Name, "ticket-") {
		time.Sleep(time.Second * 5)
		_, err := s.ChannelMessageSend(m.ID, os.Getenv("TICKET_OPEN_MESSAGE"))
		if err != nil {
			logrus.Error("Cannot send ticket message: ", err.Error())
		}
	}
}
