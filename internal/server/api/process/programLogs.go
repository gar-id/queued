package process

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/tools"
	"github.com/hpcloud/tail"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
)

func TailLogs(c *fiber.Ctx) (result error) {
	// Setup header for SSE
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	// Parse GET queries
	params := c.Queries()

	// Check if processName params is not null and exist
	if params["processName"] != "" {
		// Parse processname
		name := strings.Split(params["processName"], ":")
		if len(name) == 3 {
			var newName []string
			newName = append(newName, name[1:]...)
			name = newName
		} else if len(name) < 2 {
			return nil
		}
		processIndex, _ := strconv.Atoi(name[1])
		programName := &name[0]

		_, ok := caches.Data.ProgramConfig[*programName]
		if ok {
			// Send logs
			c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
				// Check logs directory
				var wg sync.WaitGroup
				// Stderr logs output
				if caches.Data.ProgramConfig[*programName].Stdout != "" && !strings.Contains(caches.Data.ProgramConfig[*programName].Stdout, "/dev/stderr") && !strings.Contains(caches.Data.ProgramConfig[*programName].Stdout, "/dev/stdout") {
					stdoutLog := tools.TextTemplate(caches.Data.ProgramConfig[*programName].Stdout, caches.Data.ProgramConfig[*programName].Process[processIndex])
					wg.Add(1)
					go func() {
						defer wg.Done()

						tailLogs, err := tail.TailFile(stdoutLog, tail.Config{
							Follow: true,
						})
						if err != nil {
							message := fmt.Sprintf("Error when tailing %v stdout. Error: %v", params["processName"], err)
							tools.ZapLogger("both").Error(message)
							fmt.Fprint(w, message)
							w.Flush()
						}

						var stdOutCount int = 0
						lastLine, err := tools.FileLinesCount(tailLogs.Filename)
						if err != nil {
							tools.ZapLogger("both").Error(err.Error())
							return
						} else if lastLine == 0 {
							tools.ZapLogger("both").Error(fmt.Sprintf("file %v is empty or doesn't exist", tailLogs.Filename))
							return
						} else if lastLine > 5 {
							lastLine = lastLine - 5
						}
						for line := range tailLogs.Lines {
							stdOutCount++
							if stdOutCount >= lastLine {
								fmt.Fprintf(w, "data: %v stdout: %v\n\n", params["processName"], line.Text)
								w.Flush()
							}
						}
					}()
				} else {
					message := "For now, QueueD cannot tail logs to /dev/stderr or /dev/stdout"
					fmt.Fprint(w, message)
					w.Flush()
				}

				// Stdout logs output
				if caches.Data.ProgramConfig[*programName].Stderr != "" && !strings.Contains(caches.Data.ProgramConfig[*programName].Stderr, "/dev/stderr") && !strings.Contains(caches.Data.ProgramConfig[*programName].Stderr, "/dev/stdout") && caches.Data.ProgramConfig[*programName].Stderr != caches.Data.ProgramConfig[*programName].Stdout {
					stderrLog := tools.TextTemplate(caches.Data.ProgramConfig[*programName].Stderr, caches.Data.ProgramConfig[*programName].Process[processIndex])
					wg.Add(1)
					go func() {
						defer wg.Done()

						tailLogs, err := tail.TailFile(stderrLog, tail.Config{
							Follow: true,
						})
						if err != nil {
							message := fmt.Sprintf("Error when tailing %v stderr. Error: %v", params["processName"], err)
							tools.ZapLogger("both").Error(message)
							fmt.Fprint(w, message)
							w.Flush()
						}

						var stdErrCount int = 0
						lastLine, err := tools.FileLinesCount(tailLogs.Filename)
						if err != nil {
							tools.ZapLogger("both").Error(err.Error())
							return
						} else if lastLine == 0 {
							tools.ZapLogger("both").Error(fmt.Sprintf("file %v is empty or doesn't exist", tailLogs.Filename))
							return
						} else if lastLine > 5 {
							lastLine = lastLine - 5
						}
						for line := range tailLogs.Lines {
							stdErrCount++
							if stdErrCount >= lastLine {
								fmt.Fprintf(w, "data: %v stderr: %v\n\n", params["processName"], line.Text)
								w.Flush()
							}
						}
					}()
				} else if caches.Data.ProgramConfig[*programName].Stderr == caches.Data.ProgramConfig[*programName].Stdout {
				} else {
					message := "For now, QueueD cannot tail logs to /dev/stderr or /dev/stdout\n\n"
					fmt.Fprint(w, message)
					w.Flush()
				}

				wg.Wait()
			}))

			return nil
		} else {
			// Send error message
			c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
				message := fmt.Sprintf("processName %v not found\n\n", params["processName"])
				fmt.Fprint(w, message)
				w.Flush()
			}))
			return nil
		}
	}
	// Send error message if programName params is null
	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		message := "please insert processName value\n\n"
		fmt.Fprint(w, message)
		w.Flush()
	}))
	return nil
}
