package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3475810687969-3448593142231-RExcMbDNeUf7MSq5vhIsELD3")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03EATEM7T2-3486933602976-43786ee8120d90c605186885828aa6ae014b1840bc2845197beff0dbcd9df3a9")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
