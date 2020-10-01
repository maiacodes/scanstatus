package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func updatedStatus(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	if m.Game == nil {
		return
	}
	if m.Game.Name == "Custom Status" {
		status := strings.ToLower(m.Game.State)
		for _, word := range badWords {
			if strings.Contains(status, word) {
				// Do something
				log.Warn("Bad word `", word, "` sent by ", m.User.ID)
			}
		}
	}
}
