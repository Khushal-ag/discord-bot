package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateListeningStatus("You")
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
	guildID := os.Getenv("DISCORD_GUILD_ID")
	roleID := os.Getenv("DISCORD_ROLE_ID")
	memberID := event.User.ID
	fmt.Println("Assigning role to user: ", session.State.User.Username)
	err := session.GuildMemberRoleAdd(guildID, memberID, roleID)
	if err != nil {
		fmt.Println("Error assigning role to user: ", err)
	}
	session.ChannelMessageSend(os.Getenv("DISCORD_WELCOME_CHANNEL_ID"), fmt.Sprintf("Welcome %v to the server! You have been assigned the role : %v", event.User.Username, os.Getenv("DISCORD_ROLE_NAME")))
}

func AssignRoleByUserId(session *discordgo.Session, guildID string, memberID string, roleID string) {
	err := session.GuildMemberRoleAdd(guildID, memberID, roleID)
	if err != nil {
		fmt.Println("Error assigning role to user: ", err)
	}
}
