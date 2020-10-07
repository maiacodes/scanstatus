package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func dispatchAlert(alert string, s *discordgo.Session) {
	// Send alert to console
	log.Warn(alert)

	// Send message to the chosen alert channel
	alertChannel := os.Getenv("ALERTCHANNEL")
	if alertChannel == "" {
		return
	}
	_, err := s.ChannelMessageSend(alertChannel, alert)
	if err != nil {
		// Catch failed messages
		log.Error("Cannot send alert message: ", err.Error())
	}
}
