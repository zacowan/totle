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
	"github.com/spf13/viper"
)

const noteExtension = ".md"
const notesDirConfigKey = "notes_dir"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "totle",
	Short: "A simple tool to jot down thoughts",
	Long: `Aristotle was an Ancient Greek philosopher and
polymath who made many important contributions to various
subjects of thinking. Importantly, Aristotle wrote down
his thoughts!

totle is a simple tool to allow developers to write jot
their thoughts for safe-keeping in a transferrable format.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.totle.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	defaultNotesDir := path.Join(home, "Documents", "totle")
	viper.SetDefault(notesDirConfigKey, defaultNotesDir)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".totle" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".totle")
	}

	viper.SetEnvPrefix("totle")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

type TodayFormatted struct {
	Full string
	Year string
	Month string
	Day string
}

type NotesMeta struct {
	NotesDir string
	YearMonthDir string
	TodayNotePath string
	TodayNoteFilename string
	TodayFormatted TodayFormatted
}

func GetNotesMeta() NotesMeta {
	notesMeta := NotesMeta{}
	notesMeta.NotesDir = viper.GetString(notesDirConfigKey)
	notesMeta.TodayFormatted = getTodayFormatted()
	notesMeta.YearMonthDir = path.Join(notesMeta.NotesDir, notesMeta.TodayFormatted.Year, notesMeta.TodayFormatted.Month)
	notesMeta.TodayNoteFilename = notesMeta.TodayFormatted.Full + noteExtension
	notesMeta.TodayNotePath = path.Join(notesMeta.YearMonthDir, notesMeta.TodayNoteFilename)
	return notesMeta
}

func CreateDirectoryIfNotFound(path string) (created bool, err error) {
	if !PathExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
		return err != nil, err
	}
	return false, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

// See https://gosamples.dev/date-format-yyyy-mm-dd/#:~:text=To%20format%20date%20in%20Go,%2F01%2F2006%22%20layout.
func getTodayFormatted() TodayFormatted {
	now := time.Now().Local()
	full := now.Format("2006-01-02")
	return TodayFormatted{
		Full: full,
		Year: strings.Split(full, "-")[0],
		Month: strings.Split(full, "-")[1],
		Day: strings.Split(full, "-")[2],
	}
}
