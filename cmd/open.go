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
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the note file for today",
	Run: func(cmd *cobra.Command, args []string) {
		notesMeta := GetNotesMeta()

		gotoPath := getGotoPathForTodayNote(notesMeta)
		err := openWithVsCode(notesMeta.NotesDir, gotoPath)

		if err != nil {
			fmt.Println("Failed while starting 'code' command")
			cobra.CheckErr(err)
		}
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

func openWithVsCode(path string, gotoPath string) error {
	if !PathExists(path) {
		fmt.Println("Failed note - no note exists at", path)
		os.Exit(0)
	}
	codeCmd := exec.Command("code", path, "--goto " + gotoPath)
	return codeCmd.Start()
}

func getGotoPathForTodayNote(notesMeta NotesMeta) string {
	return path.Join(notesMeta.TodayFormatted.Year, notesMeta.TodayFormatted.Month, notesMeta.TodayNoteFilename)
}
