package commands

import (
	"fmt"
	"guest/body"
	"guest/knock"
	"guest/settings"
	"guest/workspace"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"golang.org/x/term"
)

func Knock(c *cli.Context) error {
	vars, err := resolveVariables(c)
	if err != nil {
		return err
	}

	w, err := resolveWorkspace(c, []string{})
	if err != nil {
		return err
	}

	path := c.Args().First()

	knockPath := w.PathInfo.Resolve(path)

	err = executeKnock(w, knockPath, vars)
	if err != nil {
		return err
	}

	return nil
}

func executeKnock(w *workspace.Workspace, path string, vars map[string]string) error {
	if !strings.HasSuffix(path, ".knock.json") {
		path += ".knock.json"
	}

	res, err := w.Knock(path, vars)
	if err != nil {
		return err
	}

	width, _, err := term.GetSize(0)
	if err != nil {
		return nil
	}

	for _, status := range res.Statuses {
		valueColorized := colorizeByStatusFlag(status.Flags, color.Bold.Sprint(status.Value))

		fmt.Printf("%s: %s ", status.Title, valueColorized)
	}

	fmt.Println()

	for _, entity := range res.Entities {
		titleLen := len(entity.Title)
		padLen := width/2 - titleLen - 1

		color.Red.Printf(
			color.Bold.Sprint("-%s%s\n"),
			strings.ToUpper(entity.Title),
			strings.Repeat("-", padLen),
		)

		body, ok := entity.Entity.(body.Body)
		if ok {
			quick.Highlight(os.Stdout, body.String(), "json", "terminal256", "swapoff")
			fmt.Println()
		}

		items, ok := entity.Entity.(settings.KeyValueMap)
		if ok {
			maxKeyLen := 0
			maxValueLen := 0

			gapLen := 4

			sortedItems := items.SortedSlice()

			for _, item := range sortedItems {
				keyLen := len(item.Key)
				valueLen := len(item.Value)

				if keyLen > maxKeyLen {
					maxKeyLen = keyLen
				}

				if valueLen > maxValueLen {
					maxValueLen = valueLen
				}
			}

			for _, item := range sortedItems {
				keyLen := len(item.Key)

				fmt.Printf(
					"%s%s%s\n",
					item.Key,
					strings.Repeat(" ", maxKeyLen+gapLen-keyLen),
					color.Bold.Sprint(item.Value),
				)
			}
		}
	}

	return nil
}

func colorizeByStatusFlag(flags []knock.StatusFlag, format string, a ...any) string {
	if slices.Contains(flags, knock.StatusFlagGood) {
		return color.Green.Sprintf(format, a...)
	}

	if slices.Contains(flags, knock.StatusFlagWarning) {
		return color.Yellow.Sprintf(format, a...)
	}

	if slices.Contains(flags, knock.StatusFlagBad) {
		return color.Red.Sprintf(format, a...)
	}

	return fmt.Sprintf(format, a...)
}
