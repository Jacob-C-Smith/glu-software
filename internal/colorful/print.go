/** !
 *
 *
 *
 */

// Package clause
package colorful

// Imports
import (
	"fmt"
)

// Type definitions
type Color int
type LogLevel int

// Enumerations
var Colors = struct {
	Black        Color
	Red          Color
	Green        Color
	Yellow       Color
	Blue         Color
	Magenta      Color
	Cyan         Color
	DarkGray     Color
	Gray         Color
	LightRed     Color
	LightGreen   Color
	LightYellow  Color
	LightBlue    Color
	LightMagenta Color
	LightCyan    Color
	White        Color
}{
	Black:        30,
	Red:          31,
	Green:        32,
	Yellow:       33,
	Blue:         34,
	Magenta:      35,
	Cyan:         36,
	DarkGray:     90,
	Gray:         37,
	LightRed:     91,
	LightGreen:   92,
	LightYellow:  93,
	LightBlue:    94,
	LightMagenta: 95,
	LightCyan:    96,
	White:        97,
}

var LogLevels = struct {
	Error   LogLevel
	Warning LogLevel
	Info    LogLevel
	Verbose LogLevel
}{
	Error:   LogLevel(Colors.Red),
	Warning: LogLevel(Colors.Yellow),
	Info:    LogLevel(Colors.LightCyan),
	Verbose: LogLevel(Colors.White),
}

// Function declarations
/** !
 * Log a message to standard out
 *
 * @param level  the severity of the log; < Error | Warning | Info | Verbose >
 * @param format the format string
 * @param a      variadic parameters
 *
 * @return void
 */
func Log(level LogLevel, format string, a ...any) {

	// Set the color with ANSI codes
	fmt.Printf("\033[%dm", level)

	// Printf wrapper
	fmt.Printf(format, a...)

	// Clear the color with ANSI codes
	fmt.Printf("\033[0m")
}

func Printf(color Color, format string, a ...any) (n int, err error) {

	// Set the color with ANSI codes
	fmt.Printf("\033[%dm", color)

	// Printf wrapper
	nw, errr := fmt.Printf(format, a...)

	// Clear the color with ANSI codes
	fmt.Printf("\033[0m")

	return nw, errr
}
