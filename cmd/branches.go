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
	"errors"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// see https://docs.gitlab.com/ce/api/branches.html#branches-api
var branchesCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:     "branches",
		Aliases: []string{"branch"},
		Short:   "Branches",
		Long:    `Manage repository branches`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot use this command without further sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#list-repository-branches
type branchesListFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
}

var branchesListCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesListFlags{},
	Opts:   &gitlab.ListBranchesOptions{},
	Cmd: &cobra.Command{
		Use:   "list",
		Short: "List repository branches",
		Long:  `Get a list of repository branches from a project, sorted by name alphabetically. This endpoint can be accessed without authentication if the repository is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesListFlags)
		opts := cmd.Opts.(*gitlab.ListBranchesOptions)
		branches, _, err := gitlabClient.Branches.ListBranches(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(branches)
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#get-single-repository-branch
type branchesGetSingleFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Branch *string `flag_name:"branch" short:"b" type:"string" required:"yes" description:"The name of the branch"`
}

var branchesGetSingleCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesGetSingleFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get single repository branch",
		Long:  `Get a single project repository branch. This endpoint can be accessed without authentication if the repository is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesGetSingleFlags)
		branch, _, err := gitlabClient.Branches.GetBranch(*flags.Id, *flags.Branch)
		if err != nil {
			return err
		}
		return OutputJson(branch)
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#protect-repository-branch
type branchesProtectFlags struct {
	Id                 *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project owned by the authenticated user"`
	Branch             *string `flag_name:"branch" short:"b" type:"string" required:"yes" description:"The name of the branch"`
	DevelopersCanPush  *bool   `flag_name:"developers_can_push" short:"p" type:"boolean" required:"no" description:"Flag if developers can push to the branch"`
	DevelopersCanMerge *bool   `flag_name:"developers_can_merge" short:"m" type:"boolean" required:"no" description:"Flag if developers can merge to the branch"`
}

var branchesProtectCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesProtectFlags{},
	Opts:   &gitlab.ProtectBranchOptions{},
	Cmd: &cobra.Command{
		Use:   "protect",
		Short: "Protect repository branch",
		Long:  `Protects a single project repository branch. This is an idempotent function, protecting an already protected repository branch still returns a 200 OK status code.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesProtectFlags)
		opts := cmd.Opts.(*gitlab.ProtectBranchOptions)
		branch, _, err := gitlabClient.Branches.ProtectBranch(*flags.Id, *flags.Branch, opts)
		if err != nil {
			return err
		}
		return OutputJson(branch)
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#unprotect-repository-branch
type branchesUnprotectFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Branch *string `flag_name:"branch" short:"b" type:"string" required:"yes" description:"The name of the branch"`
}

var branchesUnprotectCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesUnprotectFlags{},
	Cmd: &cobra.Command{
		Use:   "unprotect",
		Short: "Unprotect repository branch",
		Long:  `Unprotects a single project repository branch. This is an idempotent function, unprotecting an already unprotected repository branch still returns a 200 OK status code.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesUnprotectFlags)
		branch, _, err := gitlabClient.Branches.UnprotectBranch(*flags.Id, *flags.Branch)
		if err != nil {
			return err
		}
		return OutputJson(branch)
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#create-repository-branch
type branchesCreateFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Branch *string `flag_name:"branch" short:"b" type:"string" required:"yes" description:"The name of the branch"`
	Ref    *string `flag_name:"ref" short:"r" type:"string" required:"yes" description:"The branch name or commit SHA to create branch from"`
}

var branchesCreateCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesCreateFlags{},
	Opts:   &gitlab.CreateBranchOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create repository branch",
		Long:  `Create repository branch`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesCreateFlags)
		opts := cmd.Opts.(*gitlab.CreateBranchOptions)
		branch, _, err := gitlabClient.Branches.CreateBranch(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(branch)
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#delete-repository-branch
type branchesDeleteFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Branch *string `flag_name:"branch" short:"b" type:"string" required:"yes" description:"The name of the branch"`
}

var branchesDeleteCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesDeleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete repository branch",
		Long:  `Delete repository branch`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesDeleteFlags)
		_, err := gitlabClient.Branches.DeleteBranch(*flags.Id, *flags.Branch)
		return err
	},
}

// see https://docs.gitlab.com/ce/api/branches.html#delete-merged-branches
type branchesDeleteMergedFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var branchesDeleteMergedCmd = &golabCommand{
	Parent: branchesCmd.Cmd,
	Flags:  &branchesDeleteMergedFlags{},
	Cmd: &cobra.Command{
		Use:   "delete-merged",
		Short: "Delete merged branches",
		Long: `Will delete all branches that are merged into the project's default branch.

Protected branches will not be deleted as part of this operation.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*branchesDeleteMergedFlags)
		_, err := gitlabClient.Branches.DeleteMergedBranches(*flags.Id)
		return err
	},
}

func init() {
	branchesCmd.Init()
	branchesListCmd.Init()
	branchesGetSingleCmd.Init()
	branchesProtectCmd.Init()
	branchesUnprotectCmd.Init()
	branchesDeleteCmd.Init()
	branchesCreateCmd.Init()
	branchesDeleteMergedCmd.Init()
}
