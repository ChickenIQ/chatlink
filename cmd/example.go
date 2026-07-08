package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/chickeniq/chatlink/pkg/proto"
	"github.com/chickeniq/chatlink/pkg/socket"
)

type Ev struct{}

func (*Ev) HandleBotInfo(usernames []proto.BotInfo) {
	fmt.Printf("Bot Info:\n")
	for _, bot := range usernames {
		fmt.Printf("[%d], Username: %s\n", bot.ID, bot.Content)
	}
}

func (*Ev) HandleSignedMessage(msg proto.SignedMessage) {
	fmt.Printf("Signed Message: [%d]: %s\n", msg.ID, msg.Content)
}

func (*Ev) HandleMessage(msg proto.Message) {
	fmt.Printf("Message: [%d]: %s\n", msg.ID, msg.Content)
}

func fromStdin(s *socket.Socket) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		pack, err := proto.MessagePacket(0, text)
		if err != nil {
			fmt.Println("Error creating packet:", err)
			continue
		}

		if err := s.SendPacket(pack); err != nil {
			fmt.Println("Error sending packet:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read stdin:", err)
	}
}

func main() {
	sock, err := socket.NewSocket("/tmp/guildlink.sock", &Ev{})
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer sock.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		<-interrupt
		sock.Close()
		os.Exit(0)
	}()

	fromStdin(sock)
}
