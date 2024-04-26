package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateListeningStatus("Namastey")
	fmt.Println(session.State.User.Username + " is online!")
}

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	// convert the message content to lowercase
	messageContent := strings.ToLower(message.Content)
	if messageContent == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}

	if messageContent == "hello" {
		session.ChannelMessageSend(message.ChannelID, "Namastey!")
	}

	if messageContent == "pong" {
		session.ChannelMessageSend(message.ChannelID, "That's my line!, Anyway, Pong!")
	}

	if messageContent == "namastey" {
		session.ChannelMessageSend(message.ChannelID, "Namastey!")
	}

}

func AssignRoleOnUserJoin(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	guildID := event.GuildID
	memberID := event.Member.User.ID
	roleID := os.Getenv("DISCORD_ROLE_ID")
	fmt.Println("Assigning role to user: ", memberID)
	err := session.GuildMemberRoleAdd(guildID, memberID, roleID)
	if err != nil {
		fmt.Println("Error assigning role to user: ", err)
	}
}

func AssignRoleByUserId(session *discordgo.Session, guildID string, memberID string, roleID string) {
	err := session.GuildMemberRoleAdd(guildID, memberID, roleID)
	if err != nil {
		fmt.Println("Error assigning role to user: ", err)
	}
}
