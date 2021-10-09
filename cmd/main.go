package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/julienschmidt/httprouter"
)

const (
	BOT_TOKEN = "YOUR_TOKEN_HERE"
	PREFIX    = "go"
)

func handleMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Got a message, ", msg.Content)
	message := strings.Split(strings.ToLower(msg.Content), " ")

	if message[0] == PREFIX {
		commands := message[1:]
		fmt.Println("Got command, ", commands)

		// YOUR CODE
		switch commands[0] {
		case "help":
			session.ChannelMessageSend(msg.ChannelID, "I'll look for therapy places for you in my free time")
		}
		//
	}
}

func createHttpServer() {
	router := httprouter.New()
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	bot, err := discordgo.New("Bot " + BOT_TOKEN)
	fmt.Println("API version:", discordgo.APIVersion)
	if err != nil {
		fmt.Println("Error creating bot session!")
		panic(err)
	}
	bot.AddHandler(handleMessage)
	bot.Open()
	createHttpServer()
	bot.Close()
}
