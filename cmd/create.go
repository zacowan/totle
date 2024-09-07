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

	"github.com/spf13/cobra"
	"github.com/zacowan/totle/pkg/fileops"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a note file for today and opens it",
	Run: func(cmd *cobra.Command, args []string) {
		notesMeta := GetNotesMeta()

		createYearMonthDir(notesMeta)

		if !fileops.PathExists(notesMeta.TodayNotePath) {
			todayAsMarkdownTitle := "# " + notesMeta.TodayFormatted.Full
			createNoteFile(notesMeta.TodayNotePath, todayAsMarkdownTitle)
		}

		openNoteFile(notesMeta)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createYearMonthDir(notesMeta NotesMeta) {
	created, err := fileops.CreateDirectoryIfNotFound(notesMeta.YearMonthDir)
	if err != nil {
		fmt.Println("Failed to create year/month directory", notesMeta.YearMonthDir)
		cobra.CheckErr(err)
	}
	if created {
		fmt.Println("Created new directory at", notesMeta.YearMonthDir)
	}
}

func createNoteFile(path string, contents string) {
	err := fileops.CreateFile(path, contents)
	if err != nil {
		fmt.Println("Failed to create note file at", path)
		cobra.CheckErr(err)
	}
	fmt.Println("Created new note file at", path)
}
