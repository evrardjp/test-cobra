package notifications

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Action uint

const (
	Drain = Action(iota)
	Reboot
)

type Notifier interface {
	Notify(action Action) error
}

type ShoutrrrNotifier struct {
	NotifyURL string
	DrainMsg  string
	RebootMsg string
}

func (shtrn ShoutrrrNotifier) Notify(action Action) error {
	log.Debug("Triggering Shoutrrr notification")
	switch action {
	case Drain:
		fmt.Println(shtrn.DrainMsg)
	case Reboot:
		fmt.Println(shtrn.RebootMsg)
	default:
		log.Error("INVALID")
	}
	return nil
}

type SlackNotifier struct {
	URL       string
	Username  string
	Channel   string
	DrainMsg  string
	RebootMsg string
}

func (slackn SlackNotifier) Notify(action Action) error {
	log.Debug("Triggering Slack notification")
	switch action {
	case Drain:
		fmt.Println(slackn.DrainMsg)
		fmt.Println(slackn.URL)
		fmt.Println(slackn.Channel)
		fmt.Println(slackn.Username)
	case Reboot:
		fmt.Println(slackn.RebootMsg)
		fmt.Println(slackn.URL)
		fmt.Println(slackn.Channel)
		fmt.Println(slackn.Username)
	default:
		log.Error("INVALID")
	}
	return nil
}

type SimpleNotify struct {
	DrainMsg  string
	RebootMsg string
}

func (smpl SimpleNotify) Notify(action Action) error {
	switch action {
	case Drain:
		log.Info(smpl.DrainMsg)
	case Reboot:
		log.Info(smpl.RebootMsg)
	default:
		log.Error("INVALID")
	}
	return nil
}
