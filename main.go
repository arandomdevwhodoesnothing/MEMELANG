package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// MemeLang - A Turing-complete language where syntax is pure meme
//
// KEYWORDS:
//   yolo <var> <val>         → assign variable
//   bussin <var> <val>       → assign variable (alias)
//   noice <var>              → print variable
//   slay                     → print newline
//   vibe <var> <op> <a> <b> → arithmetic: op = rizzup(+) lowkey(-) gyatt(*) ratio(/)
//   npc <label>:             → label (jump target)
//   yeet <label>             → unconditional jump
//   sheesh <a> <op> <b> <lbl> → conditional jump: op = fr(==) sus(!=) based(>) mid(<)
//   rizz <var>               → read input into var
//   poggers <code>           → eval raw expression (unused, placeholder)
//   gg                       → exit program
//   W <msg>                  → print string literal
//   L <msg>                  → print to stderr
//   bruh <var>               → increment var by 1
//   oof <var>                → decrement var by 1
//   copium <label> <var> <n> → loop: jump to label n times using var as counter
//   touch_grass              → sleep 0 (no-op, just vibes)
//   ratio <a> <b> <var>     → var = a % b
//   cook <var> <val>         → push val to stack named var
//   ate <var> <dest>         → pop from stack var into dest
//   rent_free <var>          → print stack size of var

type Interpreter struct {
	vars   map[string]int64
	stacks map[string][]int64
	lines  []string
	labels map[string]int
	pc     int
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		vars:   make(map[string]int64),
		stacks: make(map[string][]int64),
		labels: make(map[string]int),
	}
}

func (interp *Interpreter) load(src string) {
	interp.lines = strings.Split(src, "\n")
	// First pass: collect labels
	for i, line := range interp.lines {
		line = strings.TrimSpace(line)
		if strings.HasSuffix(line, ":") && strings.HasPrefix(line, "npc ") {
			label := strings.TrimPrefix(line, "npc ")
			label = strings.TrimSuffix(label, ":")
			interp.labels[strings.TrimSpace(label)] = i
		}
	}
}

func (interp *Interpreter) getVal(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return v
	}
	return interp.vars[s]
}

func (interp *Interpreter) run() {
	reader := bufio.NewReader(os.Stdin)
	for interp.pc < len(interp.lines) {
		line := strings.TrimSpace(interp.lines[interp.pc])
		interp.pc++

		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		// Skip label lines
		if strings.HasSuffix(line, ":") && strings.HasPrefix(line, "npc ") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "yolo", "bussin":
			if len(parts) >= 3 {
				interp.vars[parts[1]] = interp.getVal(parts[2])
			}

		case "noice":
			if len(parts) >= 2 {
				fmt.Print(interp.vars[parts[1]])
			}

		case "slay":
			fmt.Println()

		case "W":
			// print string (rest of line after W)
			msg := strings.Join(parts[1:], " ")
			fmt.Print(msg)

		case "L":
			msg := strings.Join(parts[1:], " ")
			fmt.Fprintln(os.Stderr, msg)

		case "vibe":
			// vibe <dest> <op> <a> <b>
			if len(parts) >= 5 {
				a := interp.getVal(parts[3])
				b := interp.getVal(parts[4])
				var result int64
				switch parts[2] {
				case "rizzup":
					result = a + b
				case "lowkey":
					result = a - b
				case "gyatt":
					result = a * b
				case "ratio":
					if b != 0 {
						result = a / b
					}
				}
				interp.vars[parts[1]] = result
			}

		case "ratio":
			// ratio <a> <b> <dest>
			if len(parts) >= 4 {
				a := interp.getVal(parts[1])
				b := interp.getVal(parts[2])
				if b != 0 {
					interp.vars[parts[3]] = a % b
				}
			}

		case "yeet":
			if len(parts) >= 2 {
				label := parts[1]
				if idx, ok := interp.labels[label]; ok {
					interp.pc = idx + 1
				} else {
					fmt.Fprintln(os.Stderr, "yeet failed: unknown label", label)
				}
			}

		case "sheesh":
			// sheesh <a> <op> <b> <label>
			if len(parts) >= 5 {
				a := interp.getVal(parts[1])
				op := parts[2]
				b := interp.getVal(parts[3])
				label := parts[4]
				cond := false
				switch op {
				case "fr":
					cond = a == b
				case "sus":
					cond = a != b
				case "based":
					cond = a > b
				case "mid":
					cond = a < b
				}
				if cond {
					if idx, ok := interp.labels[label]; ok {
						interp.pc = idx + 1
					}
				}
			}

		case "bruh":
			if len(parts) >= 2 {
				interp.vars[parts[1]]++
			}

		case "oof":
			if len(parts) >= 2 {
				interp.vars[parts[1]]--
			}

		case "rizz":
			if len(parts) >= 2 {
				fmt.Print("> ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
				v, err := strconv.ParseInt(text, 10, 64)
				if err == nil {
					interp.vars[parts[1]] = v
				}
			}

		case "cook":
			// push val onto stack
			if len(parts) >= 3 {
				stackName := parts[1]
				val := interp.getVal(parts[2])
				interp.stacks[stackName] = append(interp.stacks[stackName], val)
			}

		case "ate":
			// pop from stack into var
			if len(parts) >= 3 {
				stackName := parts[1]
				dest := parts[2]
				s := interp.stacks[stackName]
				if len(s) > 0 {
					interp.vars[dest] = s[len(s)-1]
					interp.stacks[stackName] = s[:len(s)-1]
				}
			}

		case "rent_free":
			if len(parts) >= 2 {
				fmt.Print(len(interp.stacks[parts[1]]))
			}

		case "touch_grass":
			// no-op, you just touched grass. good job.

		case "sus":
			// random number into var: sus <var> <max>
			if len(parts) >= 3 {
				max := interp.getVal(parts[2])
				if max > 0 {
					interp.vars[parts[1]] = rand.Int63n(max)
				}
			}

		case "gg":
			os.Exit(0)

		default:
			// Unknown command, just skip (it's giving NPC energy)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
 __  __                   _                       
|  \/  | ___ _ __ ___   ___| |    __ _ _ __   __ _ 
| |\/| |/ _ \ '_ ' _ \ / _ \ |   / _' | '_ \ / _' |
| |  | |  __/ | | | | |  __/ |__| (_| | | | | (_| |
|_|  |_|\___|_| |_| |_|\___|_____\__,_|_| |_|\__, |
                                               |___/ 
  🐸 The Only Programming Language That Slaps 🐸

Usage: memelang <file.meme>

SYNTAX (no cap):
  yolo <var> <val>              - assign variable (you only live once)
  bussin <var> <val>            - also assign (it's bussin fr fr)
  noice <var>                   - print variable
  slay                          - print newline
  W <message>                   - print string (this is a W)
  vibe <dest> <op> <a> <b>      - math: rizzup(+) lowkey(-) gyatt(*) ratio(/)
  ratio <a> <b> <dest>          - dest = a % b
  bruh <var>                    - increment (bruh moment++)
  oof <var>                     - decrement (oof--)
  npc <label>:                  - define label
  yeet <label>                  - jump to label
  sheesh <a> <op> <b> <label>   - conditional jump: fr(==) sus(!=) based(>) mid(<)
  rizz <var>                    - read input
  cook <stack> <val>            - push to stack
  ate <stack> <dest>            - pop from stack
  rent_free <stack>             - print stack size
  sus <var> <max>               - random number
  touch_grass                   - no-op (good for you)
  gg                            - exit

Example programs in examples/ folder.
`)
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "L + ratio: can't read file:", err)
		os.Exit(1)
	}

	interp := NewInterpreter()
	interp.load(string(data))
	interp.run()
}
