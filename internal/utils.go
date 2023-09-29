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
package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const NOTE_EXT = ".md"

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

func GetNotesMeta(notesDirName string) NotesMeta {
	notesMeta := NotesMeta{}
	notesMeta.NotesDir = getNotesDir(notesDirName)
	notesMeta.TodayFormatted = getTodayFormatted()
	notesMeta.YearMonthDir = path.Join(notesMeta.NotesDir, notesMeta.TodayFormatted.Year, notesMeta.TodayFormatted.Month)
	notesMeta.TodayNoteFilename = notesMeta.TodayFormatted.Full + NOTE_EXT
	notesMeta.TodayNotePath = path.Join(notesMeta.YearMonthDir, notesMeta.TodayNoteFilename)
	return notesMeta
}

func getNotesDir(name string) string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory", err)
		panic(err)
	}
	return path.Join(homeDirectory, "Documents", name)
}

// See https://gosamples.dev/date-format-yyyy-mm-dd/#:~:text=To%20format%20date%20in%20Go,%2F01%2F2006%22%20layout.
func getTodayFormatted() TodayFormatted {
	now := time.Now().UTC()
	full := now.Format("2006-01-02")
	return TodayFormatted{
		Full: full,
		Year: strings.Split(full, "-")[0],
		Month: strings.Split(full, "-")[1],
		Day: strings.Split(full, "-")[2],
	}
}
