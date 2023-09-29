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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zacowan/totle/internal"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a note file for today and opens it",
	Run: func(cmd *cobra.Command, args []string) {
		notesDirName := viper.GetString(notesDirNameKey)
		notesMeta := internal.GetNotesMeta(notesDirName)

		createYearMonthDir(notesMeta)

		if !internal.PathExists(notesMeta.TodayNotePath) {
			todayAsMarkdownTitle := "# " + notesMeta.TodayFormatted.Full
			createNoteFile(notesMeta.TodayNotePath, todayAsMarkdownTitle)
		}

		gotoPath := getGotoPathForTodayNote(notesMeta)
		openWithVsCode(notesMeta.NotesDir, gotoPath)
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

func createNoteFile(path string, contents string) {
	err := os.WriteFile(path, []byte(contents), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create note file at", path)
		cobra.CheckErr(err)
	}
	fmt.Println("Created new note file at", path)
}