package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/discord/handlers"
	"github.com/joho/godotenv"
)

var Token string
var BotID string
var BotPrefix string = "!MixR"

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	flag.StringVar(&Token, "t", os.Getenv("DISCORD_BOT_TOKEN"), "Bot Token")
	flag.Parse()
}

func main() {
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}
	u, err := s.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details: ", err)
	}
	BotID = u.ID
	s.AddHandler(handlers.Ready)
	s.AddHandler(handlers.AssignRoleOnUserJoin)
	s.AddHandler(handlers.MessageCreate)
	bot := s.Open()
	if bot != nil {
		fmt.Println("Error opening connection to Discord: ", bot)
		return
	}
	defer s.Close()
	fmt.Println("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
