package commands

import (
	"context"
	"file-editor/api"
	"log"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	content  [][]rune
	cursorX  int
	cursorY  int
	fileName string
)

func initEditor(t api.TextEditorClient) string {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	draw()

	return handleEvents(t)
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y, line := range content {
		for x, ch := range line {
			termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.SetCursor(cursorX, cursorY)

	width, height := termbox.Size()
	helpMessageLeft := "CTRL+S para guardar"
	helpMessageRight := "ESC para salir"

	for i, ch := range helpMessageLeft {
		termbox.SetCell(i, height-1, ch, termbox.ColorDefault, termbox.ColorDefault)
	}

	for i, ch := range helpMessageRight {
		termbox.SetCell(width-len(helpMessageRight)+i, height-1, ch, termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.Flush()
}

func handleEvents(t api.TextEditorClient) string {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return "Archivo cerrado satisfactoriamente..."
			}
			if ev.Key == termbox.KeyArrowUp {
				moveCursor(0, -1)
			}
			if ev.Key == termbox.KeyArrowDown {
				moveCursor(0, 1)
			}
			if ev.Key == termbox.KeyArrowLeft {
				moveCursor(-1, 0)
			}
			if ev.Key == termbox.KeyArrowRight {
				moveCursor(1, 0)
			}
			if ev.Key == termbox.KeyCtrlS {
				saveFile(t)
				return "Archivo editado satisfactoriamente..."
			}
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				deleteChar()
			}
			if ev.Key == termbox.KeySpace {
				insertChar(' ')
			}
			if ev.Key == termbox.KeyEnter {
				insertNewline()
			}
			if ev.Ch != 0 {
				insertChar(ev.Ch)
			}
			draw()
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
	}
}

func moveCursor(dx, dy int) {
	cursorX += dx
	cursorY += dy

	if cursorY < 0 {
		cursorY = 0
	}
	if cursorY >= len(content) {
		cursorY = len(content) - 1
	}
	if cursorX < 0 {
		cursorX = 0
	}
	if cursorX >= len(content[cursorY]) {
		cursorX = len(content[cursorY])
	}
}

func insertChar(ch rune) {
	line := content[cursorY]
	content[cursorY] = append(line[:cursorX], append([]rune{ch}, line[cursorX:]...)...)
	cursorX++
}

func deleteChar() {
	if cursorX > 0 {
		line := content[cursorY]
		content[cursorY] = append(line[:cursorX-1], line[cursorX:]...)
		cursorX--
	} else if cursorY > 0 {
		prevLine := content[cursorY-1]
		currentLine := content[cursorY]
		content = append(content[:cursorY], content[cursorY+1:]...)
		content[cursorY-1] = append(prevLine, currentLine...)
		cursorY--
		cursorX = len(prevLine)
	}
}

func insertNewline() {
	line := content[cursorY]
	newLine := line[cursorX:]
	content[cursorY] = line[:cursorX]
	content = append(content[:cursorY+1], append([][]rune{newLine}, content[cursorY+1:]...)...)
	cursorY++
	cursorX = 0
}

func saveFile(t api.TextEditorClient) {
	var flatContent []byte
	for _, line := range content {
		flatContent = append(flatContent, []byte(string(line))...)
		flatContent = append(flatContent, '\n')
	}

	request := api.SaveFileRequest{
		Filename: fileName,
		Content:  flatContent,
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute*1)

	_, err := t.SaveFile(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
}

func splitLines(data []byte) [][]rune {
	var lines [][]rune
	var line []rune
	for _, ch := range string(data) {
		if ch == '\n' {
			lines = append(lines, line)
			line = nil
		} else {
			line = append(line, ch)
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}

type EditCommand struct {
	Name string
}

func (c *EditCommand) Run(t api.TextEditorClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	request := api.ReadFileRequest{
		Filename: c.Name,
	}

	r, err := t.ReadFile(ctx, &request)
	if err != nil {
		return "", err
	}

	content = splitLines([]byte(r.Content))
	fileName = c.Name

	var response strings.Builder
	response.WriteString("\n===== Resultado ======\n")
	response.WriteString(initEditor(t))
	response.WriteString("\n======================\n")

	return response.String(), nil
}

func (c *EditCommand) Print() {
	print("read command", c.Name)
}
