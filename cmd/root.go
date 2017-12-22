// Copyright © 2017 Lee Briggs <lee@leebriggs.co.uk>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// external packages
	"github.com/ahmdrz/goinsta"
	//"github.com/davecgh/go-spew/spew"
	log "github.com/Sirupsen/logrus"
)

var cfgFile string
var username string
var password string
var account string

var dryRun bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "like-my-friend",
	Short: "Like all of the pictures on an instgram account quickly",
	Long: `If your friend likes you to be quick to like their photos,
this app will poll the latest feed, determine if you've liked their feed recently
and like the photo if you haven't`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		username = viper.GetString("username")
		password = viper.GetString("password")
		account = viper.GetString("account")
		dryRun = viper.GetBool("dryRun")

		if username == "" {
			log.Fatal("Please specify a username to login as")
		}
		if password == "" {
			log.Fatal("Please specify a password to login with")
		}
		if account == "" {
			log.Fatal("Please specify an account to follow")
		}

		insta := goinsta.New(username, password)
		if err := insta.Login(); err != nil {
			log.Fatal(err)
		}
		defer insta.Logout()

		friend, err := insta.GetUserByUsername(account)
		if err != nil {
			log.Fatal(err)
			return
		}

		friendID := friend.User.ID
		//spew.Dump(friend.User)

		log.Printf("Checking User %s with name %s and Instagram ID is %v", account, friend.User.FullName, friendID)

		feed, err := insta.LatestUserFeed(friendID)

		for _, item := range feed.Items {

			//spew.Dump(item)

			if item.HasLiked == false {
				log.WithFields(log.Fields{"status": item.HasLiked, "ID": item.ID, "#_comments": item.CommentCount}).Warning("Photo not liked, liking it!")
				if dryRun == true {
					log.WithFields(log.Fields{"status": item.HasLiked, "ID": item.ID, "#_comments": item.CommentCount}).Info("Would have liked photo, but dry run enabled")
				} else {
					_, err := insta.Like(item.ID)
					if err != nil {
						log.Error("Error liking photo", err)
					}
				}
			} else {
				log.WithFields(log.Fields{"status": item.HasLiked, "ID": item.ID, "#_comments": item.CommentCount}).Info("Photo already liked!")
			}
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.like-my-friend.yaml)")
	RootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "The username for your account")
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "The password for your account")
	RootCmd.PersistentFlags().StringVarP(&account, "account", "a", "", "The instagram account to watch")
	RootCmd.PersistentFlags().BoolVarP(&dryRun, "dryrun", "d", false, "Show what photos you'd like")
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("username", RootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("account", RootCmd.PersistentFlags().Lookup("account"))
	viper.BindPFlag("dryrun", RootCmd.PersistentFlags().Lookup("dryrun"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".like-my-friend") // name of config file (without extension)
		viper.AddConfigPath("$HOME")           // adding home directory as first search path
		viper.SetEnvPrefix("insta")
		viper.AutomaticEnv() // read in environment variables that match
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
