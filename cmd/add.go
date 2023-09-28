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
	"os"
	"path"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new note to today's note file",
	Run: func(cmd *cobra.Command, args []string) {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Could not find home directory")
			os.Exit(1)
		}
		// Check if the notes directory exists
		notesDirectory := path.Join(homeDirectory, "Documents", "totle")
		_, err = safeCreateDirectory(notesDirectory)
		if err != nil {
			fmt.Println("Failed to create notes directory")
			os.Exit(2)
		}
		// Check if the directory for the year exists
		now := time.Now()
		year := strconv.Itoa(now.Year())
		yearDirectory := path.Join(notesDirectory, year)
		_, err = safeCreateDirectory(yearDirectory)
		if err != nil {
			fmt.Println("Failed to create year directory")
			os.Exit(3)
		}
		// Check if the directory for the month exists
		month := strconv.Itoa(int(now.Month()))
		monthDirectory := path.Join(yearDirectory, month)
		_, err = safeCreateDirectory(monthDirectory)
		if err != nil {
			fmt.Println("Failed to create month directory")
			os.Exit(4)
		}
		// Check if a note file for the day exists
		// If no note file exists, create one with a new note
		// If a note file exists, append a new note to the file
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func safeCreateDirectory(path string) (created bool, err error) {
	created = false
	if !pathExists(path) {
		created = true
		err = os.Mkdir(path, os.ModePerm)
	}
	return
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
