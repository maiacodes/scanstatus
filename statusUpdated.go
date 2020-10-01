package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	userStatus map[string]string
)

func updatedStatus(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	// If map doesn't exist yet, make it.
	if userStatus == nil {
		userStatus = make(map[string]string)
		log.Info("Created map")
	}

	// If there isn't a game, return
	if m.Game == nil {
		return
	}

	// Check for custom status
	if m.Game.Name == "Custom Status" {
		status := strings.ToLower(m.Game.State)

		// Ignore existing statuses
		if userStatus[m.User.ID] == status {
			return
		}

		// Check badwords for matches
		for _, word := range badWords {
			if strings.Contains(status, word) {
				// Do something
				dispatchAlert(fmt.Sprintf("Bad word `%v` in status by <@%v>", word, m.User.ID), s)
			}
		}

		// Store current status
		userStatus[m.User.ID] = status
	}
}
