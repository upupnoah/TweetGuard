/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
)

const (
	consumerKey    = "YOUR_CONSUMER_KEY"
	consumerSecret = "YOUR_CONSUMER_SEC"
	accessToken    = "YOUR_ACCESS_TOKEN"
	accessSecret   = "YOUR_ACCESS_SEC"
)

func createTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// create Twitter API client
	client := twitter.NewClient(httpClient)

	return client
}

func getYellowContentUsers(client *twitter.Client) ([]string, error) {
	// get the timeline of current user
	timeline, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 100,
	})
	if err != nil {
		return nil, err
	}

	// store the screen names of Twitter users with pronographic content
	var yellowContentUsers []string

	for _, tweet := range timeline {
		// judge whether the tweet contains a pronographic content
		if containsYellowContent(tweet.Text) {
			screenName := tweet.User.ScreenName
			yellowContentUsers = append(yellowContentUsers, screenName)
		}

		// get the content of reply, determine if there is a reply
		if tweet.InReplyToStatusID != 0 {
			reply, _, err := client.Statuses.Show(tweet.InReplyToStatusID, nil)
			if err != nil {
				return nil, err
			}

			if containsYellowContent(reply.Text) {
				screenName := reply.User.ScreenName
				yellowContentUsers = append(yellowContentUsers, screenName)
			}
		}
	}
	return yellowContentUsers, nil
}

func containsYellowContent(s string) bool {
	// custom
	pattern := `.*[色情|黄色|裸体|成人|淫秽|床戏|做爱|三级片|AV|艳照|情色|裸聊|黄网|激情|性爱|同性恋|合成图|SM|调教|捆绑|性虐|乳房|性交|阴茎|阴道|胸部|乳头|私处|阴毛|肛门|口交|手淫|性欲|性感|诱惑|私房照|性玩具|兽交|群交|潮吹|自慰|阳具|阴蒂|口活|性奴|虐待|猥亵|强奸|暴露狂|换妻|嫖娼].*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(s)
}

func blockUsers(client *twitter.Client, users []string) error {
	// block users
	for _, user := range users {
		_, _, err := client.Blocks.Create(&twitter.BlockCreateParams{
			ScreenName: user,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block users on Twitter",

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("block called")
		cmd.Usage()
	},
}

var pronCmd = &cobra.Command{
	Use:   "pron",
	Short: "Block users who post explicit content on Twitter",
	Run: func(cmd *cobra.Command, args []string) {
		// 在这里实现获取时间线上发布黄色内容的推特用户并执行 block 操作的逻辑
		client := createTwitterClient()
		yellowContentUsers, err := getYellowContentUsers(client)
		if err != nil {
			panic(err)
		}
		err = blockUsers(client, yellowContentUsers)
		if err != nil {
			panic(err)
		}
		fmt.Println("Blocking users who post explicit content on Twitter...")
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
	blockCmd.AddCommand(pronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
