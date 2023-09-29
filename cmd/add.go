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
	"strings"

	"github.com/spf13/cobra"
	"github.com/zacowan/totle/internal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new note to today's note file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No note was provided to jot down")
			os.Exit(0)
		}
		providedNote := args[0]

		notesMeta := internal.GetNotesMeta("totle")

		// Verify the directory for the note the reside in exists
		created, err := createDirectoryIfNotFound(notesMeta.YearMonthDir)
		if err != nil {
			fmt.Println("Error creating year/month directory", notesMeta.YearMonthDir)
			panic(err)
		}
		if created {
			fmt.Println("Created new directory at", notesMeta.YearMonthDir)
		}

		// If no note file exists, create one with a new note
		if !pathExists(notesMeta.TodayNotePath) {
			initialFileContents := "# " + notesMeta.TodayFormatted.Full + "\n\n- " + providedNote
			err := appendToFile(notesMeta.TodayNotePath, initialFileContents)
			if (err != nil) {
				fmt.Println("Failed to create new note at", notesMeta.TodayNotePath)
				panic(err)
			}
			fmt.Println("Created new note at", notesMeta.TodayNotePath)
			os.Exit(0)
		}

		// If a note file exists, append a new note to the file
		lastLine, err := getLastLineOfFile(notesMeta.TodayNotePath)
		if err != nil {
			fmt.Println("Failed to add note to file", notesMeta.TodayNotePath)
			panic(err)
		}
		fileContentsToAppend := "- " + providedNote + "\n"
		if lastLine != "" {
			fileContentsToAppend = "\n" + fileContentsToAppend
		}
		appendToFile(notesMeta.TodayNotePath, fileContentsToAppend)
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

func createDirectoryIfNotFound(path string) (created bool, err error) {
	if !pathExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
		return err != nil, err
	}
	return false, nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func appendToFile(path string, contents string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(contents)); err != nil {
		return err
	}

	return nil
}

func getLastLineOfFile(path string) (lastLine string, err error) {
	contents, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
	contentsByNewline := strings.Split(string(contents), "\n")
	return contentsByNewline[len(contentsByNewline)-1], nil
}
