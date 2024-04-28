package helpers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ListAllUsers(session *discordgo.Session, guildID string) []*discordgo.Member {
	members, err := session.GuildMembers(guildID, "", 1000)
	if err != nil {
		fmt.Println("Error getting members: ", err)
		return nil
	}
	var filteredMembers []*discordgo.Member
	for _, member := range members {
		if member.User.Bot {
			continue
		}
		filteredMembers = append(filteredMembers, member)
	}
	return filteredMembers
}
