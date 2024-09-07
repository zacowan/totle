/*
Copyright © 2023 Zachary Cowan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zacowan/totle/pkg/fileops"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the note file for today",
	Run: func(cmd *cobra.Command, args []string) {
		notesMeta := GetNotesMeta()
		openNoteFile(notesMeta)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func openNoteFile(notesMeta NotesMeta) {
	openCmdFromConfig := viper.GetString(openCmdConfigKey)
	openWithCmd(openCmdFromConfig, notesMeta.TodayNotePath)
}

func openWithCmd(cmd string, path string, arg ...string) {
	if !fileops.PathExists(path) {
		fmt.Println("Failed to open - no note file exists at", path)
		return
	}
	args := []string{path}
	args = append(args, arg...)
	openCmdExec := exec.Command(cmd, args...)
	err := openCmdExec.Start()
	if err != nil {
		fmt.Printf("Failed while starting '%s' command\n", cmd)
		cobra.CheckErr(err)
	}
}
