# 🐸 MemeLang

> The only Turing-complete programming language where the syntax is pure meme. No cap.

## Running

## Install

```bash
git clone https://github.com/arandomdevwhodoesnothing/MEMELANG.git
```

### Go (compile it yourself, king):
```bash
go build -o memelang main.go
./memelang yourfile.meme
```



---

## Full Syntax Reference

| Command | Meaning | What it does |
|---------|---------|--------------|
| `yolo <var> <val>` | You Only Live Once | Assign a variable |
| `bussin <var> <val>` | It's Bussin fr fr | Also assign a variable |
| `noice <var>` | Noice | Print a variable |
| `slay` | Slay queen | Print a newline |
| `W <message>` | This is a W | Print a string literal |
| `L <message>` | Took an L | Print to stderr |
| `vibe <dest> <op> <a> <b>` | Vibe check | Math: `rizzup`(+) `lowkey`(-) `gyatt`(*) `ratio`(/) |
| `ratio <a> <b> <dest>` | You got ratioed | Modulo: dest = a % b |
| `bruh <var>` | Bruh moment | Increment variable |
| `oof <var>` | Oof | Decrement variable |
| `npc <label>:` | NPC behavior | Define a jump label |
| `yeet <label>` | Yeet | Unconditional jump |
| `sheesh <a> <op> <b> <label>` | Sheeeesh | Conditional jump: `fr`(==) `sus`(!=) `based`(>) `mid`(<) |
| `rizz <var>` | Got Rizz | Read integer from stdin |
| `cook <stack> <val>` | We're cooking | Push value onto a named stack |
| `ate <stack> <dest>` | Ate and left no crumbs | Pop from stack into variable |
| `rent_free <stack>` | Living rent free | Print size of stack |
| `sus <var> <max>` | Kinda sus | Generate random number 0..max-1 |
| `touch_grass` | Touch grass | No-op. Good for you. |
| `gg` | GG | Exit the program |

Comments start with `#` or `//`.

---

## Examples

### Hello World
```
W no cap this language is bussin
slay
gg
```

### FizzBuzz
```
yolo i 1
yolo limit 20
yolo three 3
yolo five 5
yolo zero 0

npc loop_start:
  sheesh i based limit done
  ratio i three r3
  ratio i five r5
  sheesh r3 fr zero fizzbuzz_check_five
  sheesh r5 fr zero buzz_only
  noice i
  yeet next

npc fizzbuzz_check_five:
  sheesh r5 fr zero fizzbuzz
  W Fizz
  yeet next

npc fizzbuzz:
  W FizzBuzz
  yeet next

npc buzz_only:
  W Buzz

npc next:
  slay
  bruh i
  yeet loop_start

npc done:
  gg
```

### Fibonacci
```
yolo a 0
yolo b 1
yolo count 0
yolo limit 15

npc fib_loop:
  sheesh count based limit fib_done
  noice a
  slay
  vibe tmp rizzup a b
  yolo a b
  yolo b tmp
  bruh count
  yeet fib_loop

npc fib_done:
  W bussin fr fr
  slay
  gg
```

---

## Turing Completeness

MemeLang is Turing-complete because it has:
- **Unbounded storage** (variables + named stacks)
- **Conditional branching** (`sheesh`)
- **Unconditional jump** (`yeet`)
- **Arithmetic** (`vibe`, `ratio`, `bruh`, `oof`)
- **I/O** (`W`, `noice`, `rizz`)

This is sufficient to simulate any Turing machine. The `cook`/`ate` stack instructions make it trivial to implement recursive algorithms iteratively.

---

*Built with zero chill and maximum rizz*
