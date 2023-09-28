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
	"strings"
	"time"

	"github.com/spf13/cobra"
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
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory", err)
			panic(err)
		}
		// Check if the notes directory exists
		notesDirectory := path.Join(homeDirectory, "Documents", "totle")
		nowFormatted := getNowFormatted()
		year := strings.Split(nowFormatted, "-")[0]
		month := strings.Split(nowFormatted, "-")[1]
		currentNoteDirectory := path.Join(notesDirectory, year, month)
		created, err := createDirectoryIfNotFound(currentNoteDirectory)
		if err != nil {
			fmt.Println("Error creating note directory", currentNoteDirectory)
			panic(err)
		}
		if created {
			fmt.Println("Created new directory at", currentNoteDirectory)
		}
		currentNoteFilename := path.Join(currentNoteDirectory, nowFormatted + ".md")
		// If no note file exists, create one with a new note
		if !pathExists(currentNoteFilename) {
			initialFileContents := "# " + nowFormatted + "\n\n- " + providedNote
			err := appendToFile(currentNoteFilename, initialFileContents)
			if (err != nil) {
				fmt.Println("Error creating new note at", currentNoteFilename)
				panic(err)
			}
			fmt.Println("Created new note at", currentNoteFilename)
			os.Exit(0)
		}
		// If a note file exists, append a new note to the file
		lastLine, err := getLastLineOfFile(currentNoteFilename)
		if err != nil {
			fmt.Println("Failed to add note to file", currentNoteFilename)
			panic(err)
		}
		fileContentsToAppend := "- " + providedNote + "\n"
		if lastLine != "" {
			fileContentsToAppend = "\n" + fileContentsToAppend
		}
		appendToFile(currentNoteFilename, fileContentsToAppend)
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

// See https://gosamples.dev/date-format-yyyy-mm-dd/#:~:text=To%20format%20date%20in%20Go,%2F01%2F2006%22%20layout.
func getNowFormatted() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02")
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
