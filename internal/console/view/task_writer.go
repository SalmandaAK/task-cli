package view

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/SalmandaAK/task-cli/internal/task/domain"
)

// PrintTasks will print tasks with following format:
// \c| ID | TASK DESCRIPTION |\t\c STATUS \c\t| CREATED AT | UPDATED AT |\r\n
// \d| t.Id | t.Description |\t\c t.Status \d\t| t.CreatedAt | t.UpdatedAt |\r\n
// with \c = colored mode, \d = DEFAULT mode, \r = RESET, \t = tab which will be used by tabwritter.
// Text Modes used must have the same length and in the same position for each line, otherwise columns set by the tabwriter will break.
func PrintTasks(tasks []*domain.Task) {
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)

	fmt.Fprintf(tabwriter, "%s|\tID\t|\tTASK DESCRIPTION\t|\t%sSTATUS%s\t|\tCREATED AT\t|\tUPDATED AT\t|%s\n", cyanBold, cyanBold, cyanBold, resetMode)
	for _, t := range tasks {
		if t.UpdatedAt == "" {
			t.UpdatedAt = "-"
		}
		fmt.Fprintf(tabwriter, "%s|\t%d\t|\t%s\t|\t%s\t|\t%s\t|\t%s\t|%s\n", defaultMode, t.Id, t.Description, formattedStatusMap[t.Status], t.CreatedAt, t.UpdatedAt, resetMode)
	}
	tabwriter.Flush()
}

// Text Modes based on ANSI Escapes Code https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
const (
	cyanBold    = "\x1b[1;36m"
	greenDim    = "\x1b[2;32m"
	yellowDim   = "\x1b[2;33m"
	magentaDim  = "\x1b[2;35m"
	defaultMode = "\x1b[2;39m"
	resetMode   = "\x1b[0m"
)

var formattedStatusMap = map[string]string{
	"todo":        magentaDim + "todo" + defaultMode,
	"in-progress": yellowDim + "in progress..." + defaultMode,
	"done":        greenDim + "done!" + defaultMode,
}
