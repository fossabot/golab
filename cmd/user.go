// Copyright © 2017 Michael Lihs
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
	"errors"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/spf13/viper"
	"strconv"
)

var user, email, password, skype, linkedin, twitter, websiteUrl, organization, username, externUid, provider, bio, location string
var projectsLimit int
var admin, canCreateGroup, confirm, external bool

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Gitlab users",
	Long: `Allows create, update and deletion of a user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of the subcommands, see `golab user -h`")
	},
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get user details",
	Long: `Get detailed information for given user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if (id == "" && user == "") || (id != "" && user != "") {
			return errors.New("you either have to provide an id or a username")
		}
		var userId int
		if user != "" {
			users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{Username: &user})
			if err != nil {
				return err
			}
			if len(users) != 1 {
				return errors.New("Number of users found for username: " + strconv.Itoa(len(users)))
			}
			userId = users[0].ID
		} else {
			var err error = nil
			userId, err = strconv.Atoi(id)
			if err != nil {
				return err
			}
		}
		user, _, err := gitlabClient.Users.GetUser(userId)
		if err != nil {
			return err
		}
		err = OutputJson(user)
		return err
	},
}

var lsCmd = &cobra.Command{
	Use: "ls",
	Short: "Get list of all users",
	Long: `Get a list of all users on the Gitlab server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
		if err != nil {
			return err
		}
		err = OutputJson(users)
		return err
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long: `Allows creation of a new user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		createUserOptions := &gitlab.CreateUserOptions{
			Admin: &admin,
			Bio: &bio,
			CanCreateGroup: &canCreateGroup,
			Confirm: &confirm,
			Email: &email,
			Linkedin: &linkedin,
			Name: &name,
			Password: &password,
			ProjectsLimit: &projectsLimit,
			Skype: &skype,
			Twitter: &twitter,
			Username: &username,
			WebsiteURL: &websiteUrl,
		}
		if provider != "" {
			createUserOptions.Provider = &provider
			createUserOptions.ExternUID = &externUid
		}
		user, _, err := gitlabClient.Users.CreateUser(createUserOptions)
		if err != nil {
			return err
		}
		err = OutputJson(user)
		return err
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user",
	Long: `Delete a user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("delete called")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user",
	Long: `Allows updating a user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update called")
	},
}


func init() {
	initUserGetCommand()
	userCmd.AddCommand(lsCmd)
	initUserCreateCommand()
	userCmd.AddCommand(updateCmd)
	userCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(userCmd)
}

func initUserGetCommand() {
	getCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "(mandatory if id is unset) name of the user to look up")
	getCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "(mandatory if user is unset) id of the user to look up")
	viper.BindPFlag("user", getCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("id", getCmd.PersistentFlags().Lookup("id"))
	userCmd.AddCommand(getCmd)
}

func initUserCreateCommand() {
	createCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "(mandatory) Email")
	createCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "(mandatory) Password")
	createCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "(mandatory) Username")
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(mandatory) Name")
	createCmd.PersistentFlags().StringVarP(&skype, "skype", "", "", "(optional) Skype ID")
	createCmd.PersistentFlags().StringVarP(&linkedin, "linkedin", "", "", "(optional) LinkedIn")
	createCmd.PersistentFlags().StringVarP(&twitter, "twitter", "", "", "(optional) Twitter account")
	createCmd.PersistentFlags().StringVarP(&websiteUrl, "website_url", "", "", "(optional) Website URL")
	createCmd.PersistentFlags().StringVarP(&organization, "organization", "", "", "(optional) Organization name")
	createCmd.PersistentFlags().IntVarP(&projectsLimit, "projects_limit", "", 10, "(optional) Number of projects user can create (10 is default)")
	createCmd.PersistentFlags().StringVarP(&externUid, "extern_uid", "", "", "(optional) External UID")
	createCmd.PersistentFlags().StringVarP(&provider, "provider", "", "", "(optional) External provider name")
	createCmd.PersistentFlags().StringVarP(&bio, "bio", "", "", "(optional) User's biography")
	createCmd.PersistentFlags().StringVarP(&location, "location", "", "", "(optional) User's location")
	createCmd.PersistentFlags().BoolVarP(&admin, "admin", "a", false, "(optional) User is admin - true or false (default)")
	createCmd.PersistentFlags().BoolVarP(&canCreateGroup, "can_create_group", "", false, "(optional) User can create groups - true or false (default)")
	createCmd.PersistentFlags().BoolVarP(&confirm, "confirm", "", true, "(optional) Require confirmation - true (default) or false")
	createCmd.PersistentFlags().BoolVarP(&external, "external", "", false, "(optional) Flags the user as external - true or false(default)")
	viper.BindPFlag("email", createCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("password", createCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("username", createCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("name", createCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("skype", createCmd.PersistentFlags().Lookup("skype"))
	viper.BindPFlag("linkedin", createCmd.PersistentFlags().Lookup("linkedin"))
	viper.BindPFlag("twitter", createCmd.PersistentFlags().Lookup("twitter"))
	viper.BindPFlag("website_url", createCmd.PersistentFlags().Lookup("website_url"))
	viper.BindPFlag("organization", createCmd.PersistentFlags().Lookup("organization"))
	viper.BindPFlag("projects_limit", createCmd.PersistentFlags().Lookup("projects_limit"))
	viper.BindPFlag("extern_uid", createCmd.PersistentFlags().Lookup("extern_uid"))
	viper.BindPFlag("provider", createCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("bio", createCmd.PersistentFlags().Lookup("bio"))
	viper.BindPFlag("location", createCmd.PersistentFlags().Lookup("location"))
	viper.BindPFlag("admin", createCmd.PersistentFlags().Lookup("admin"))
	viper.BindPFlag("can_create_group", createCmd.PersistentFlags().Lookup("can_create_group"))
	viper.BindPFlag("confirm", createCmd.PersistentFlags().Lookup("confirm"))
	viper.BindPFlag("external", createCmd.PersistentFlags().Lookup("external"))
	userCmd.AddCommand(createCmd)
}
