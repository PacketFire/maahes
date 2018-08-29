package lib

import (
	"fmt"

	"github.com/bramz/maahes/lib/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func StartSession() {
	discord, err := discordgo.New("Bot NDc2MzMzMTYyMzkwMzU1OTY4.DmVuFA.Qfc3-cu0lfzaym11Ee0V5_aOUQU")
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(Parser)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer discord.Close()

	<-make(chan struct{})
}

func Parser(session *discordgo.Session, message *discordgo.MessageCreate) {
	maahes, err := session.User("@me")
	if err != nil {
		fmt.Println(err)
		return
	}

	author := message.Author
	if author.ID == maahes.ID || author.Bot {
		return
	}

	content := message.Content

	cmds := map[string]func(string) string{
		"theo":   commands.TheoCmd(content),
		"define": commands.DefineCmd(content[1:]),
		//		"quit": commands.QuitCmd,
	}

	if string(content[0]) == "!" {
		ct := content[1:]
		out := cmds[ct]
		session.ChannelMessageSend(message.ChannelID, out(content))
	}

	// colorize terminal output
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("#%s <%s>: %s\n", green(message.ChannelID), red(message.Author), blue(message.Content))
}
