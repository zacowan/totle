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
	"github.com/spf13/viper"
	"github.com/zacowan/totle/internal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new note to today's note file",
	Long: `If no note file exists for today, a new note file is created with the
contents of the note you want to add. If a note file already exists for today,
the contents of the note you want to add are appended to that file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No note was provided to add to the note file")
			os.Exit(0)
		}
		providedNote := args[0]

		notesDirName := viper.GetString(notesDirNameKey)
		notesMeta := internal.GetNotesMeta(notesDirName)

		createYearMonthDir(notesMeta)

		if !internal.PathExists(notesMeta.TodayNotePath) {
			titleWithProvidedNoteAsMarkdown := "# " + notesMeta.TodayFormatted.Full + "\n\n- " + providedNote
			createNoteFile(notesMeta.TodayNotePath, titleWithProvidedNoteAsMarkdown)
			os.Exit(0)
		}

		lastLineOfNoteFile, err := getLastLineOfFile(notesMeta.TodayNotePath)
		if err != nil {
			fmt.Println("Failed to add note to note file at", notesMeta.TodayNotePath)
			cobra.CheckErr(err)
		}
		providedNoteAsMarkdown := "- " + providedNote + "\n"
		if lastLineOfNoteFile != "" {
			providedNoteAsMarkdown = "\n" + providedNoteAsMarkdown
		}
		appendToFile(notesMeta.TodayNotePath, providedNoteAsMarkdown)
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

func createYearMonthDir(notesMeta internal.NotesMeta) {
	created, err := internal.CreateDirectoryIfNotFound(notesMeta.YearMonthDir)
	if err != nil {
		fmt.Println("Failed to create year/month directory", notesMeta.YearMonthDir)
		cobra.CheckErr(err)
	}
	if created {
		fmt.Println("Created new directory at", notesMeta.YearMonthDir)
	}
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
