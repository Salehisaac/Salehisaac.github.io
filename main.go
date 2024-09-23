package main

import (
	"log"
	"os"
	"time"
	"fmt"

	tele "gopkg.in/telebot.v3"
	"github.com/subosito/gotenv"
)

func init(){
	gotenv.Load()
}

func main() {

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("error creating bot %v", err)
		return
	}

	webApp := &tele.WebApp{
		URL: "https://salehisaac.github.io/webapp/",
	}
	
	b.Handle("/start", func(c tele.Context) error {
		markup := b.NewMarkup()
		markup.Inline(markup.Row(markup.WebApp("Open", webApp)))
		return c.Send("Open this app!", markup)
	})
	
	b.Handle(tele.OnWebApp, func(c tele.Context) error {
		webapp := c.Message().WebAppData
		fmt.Println(webapp.Data)
		c.Send("Preferences received: " + webapp.Data)
		return nil
	})

	log.Println("listening...")
	b.Start()
	
}
