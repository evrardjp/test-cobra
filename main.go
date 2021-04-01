package main

import (
	"fmt"

	notifications "github.com/evrardjp/test-cobra/notifications"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version               = "testing"
	notifyURL             string
	slackHookURL          string
	slackUsername         string
	slackChannel          string
	messageTemplateDrain  string
	messageTemplateReboot string
)

func main() {
	rootCmd := &cobra.Command{
		Use:    "kured",
		Short:  "Kubernetes Reboot Daemon",
		PreRun: flagCheck,
		Run:    root,
	}

	rootCmd.PersistentFlags().StringVar(&slackHookURL, "slack-hook-url", "",
		"slack hook URL for notifications")
	rootCmd.PersistentFlags().StringVar(&slackUsername, "slack-username", "kured",
		"slack username for notifications")
	rootCmd.PersistentFlags().StringVar(&slackChannel, "slack-channel", "",
		"slack channel for reboot notfications")
	rootCmd.PersistentFlags().StringVar(&notifyURL, "notify-url", "",
		"notify URL for reboot notfications")
	rootCmd.PersistentFlags().StringVar(&messageTemplateDrain, "message-template-drain", "Draining node %s",
		"message template used to notify about a node being drained")
	rootCmd.PersistentFlags().StringVar(&messageTemplateReboot, "message-template-reboot", "Rebooting node %s",
		"message template used to notify about a node being rebooted")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func flagCheck(cmd *cobra.Command, args []string) {
	if slackHookURL != "" && notifyURL != "" {
		log.Warnf("Cannot use both --notify-url and --slack-hook-url flags. Kured will use --notify-url flag only...")
		slackHookURL = ""
	}
	if slackChannel != "" || slackHookURL != "" {
		log.Warnf("Deprecated flag(s). Please use --notify-url flag instead.")
	}

}

func root(cmd *cobra.Command, args []string) {
	nodeid := "node1"
	log.Infof("Kubernetes Reboot Daemon: %s", version)
	log.Infof("notifyURL hook url %s", notifyURL)
	log.Infof("Slack hook url: %s", slackHookURL)
	log.Infof("slackChannel: %s", slackChannel)
	log.Infof("slackUsername: %s", slackUsername)

	drainMsg := fmt.Sprintf(messageTemplateDrain, nodeid)
	rebootMsg := fmt.Sprintf(messageTemplateDrain, nodeid)

	var notifier notifications.Notifier
	if notifyURL != "" {
		notifier = notifications.ShoutrrrNotifier{NotifyURL: notifyURL, DrainMsg: drainMsg, RebootMsg: rebootMsg}
	} else if slackHookURL != "" {
		notifier = notifications.SlackNotifier{URL: slackHookURL, Username: slackUsername, Channel: slackChannel, DrainMsg: drainMsg, RebootMsg: rebootMsg}
	} else {
		notifier = notifications.SimpleNotify{DrainMsg: drainMsg, RebootMsg: rebootMsg}
	}
	notifier.Notify(notifications.Drain)
	// if slackHookURL != "" {
	// 	if err := slack.NotifyDrain(slackHookURL, slackUsername, slackChannel, messageTemplateDrain, nodename); err != nil {
	// 		log.Warnf("Error notifying slack: %v", err)
	// 	}
	// }
	// if notifyURL != "" {
	// 	if err := shoutrrr.Send(notifyURL, fmt.Sprintf(messageTemplateDrain, nodename)); err != nil {
	// 		log.Warnf("Error notifying: %v", err)
	// 	}
	// }
}
