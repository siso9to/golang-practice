package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"time"
	"github.com/jinzhu/gorm"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id	int	 `gorm:"primary_key"`
	Email string
	Login	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func gormConnect() *gorm.DB {
	var dbname string = os.Getenv("BOT_DB")
	db, err := gorm.Open("mysql", "root:@/" + dbname + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				user := User{Login: ev.Text}

				db := gormConnect()
				db.Find(&user)

				log.Printf("Message: %v\n", user.Id)
				rtm.SendMessage(rtm.NewOutgoingMessage(ev.Text + "のIdは" + strconv.Itoa(user.Id) + "です", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	os.Exit(run(api))
}
