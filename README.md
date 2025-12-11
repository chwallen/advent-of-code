# Advent of Code solutions in Go

<!--toc:start-->
- [Advent of Code solutions in Go](#advent-of-code-solutions-in-go)
  - [Requirements](#requirements)
    - [Installing lp_solve](#installing-lp_solve)
      - [Linux](#linux)
      - [Mac](#mac)
      - [Windows](#windows)
<!--toc:end-->

These are my solutions for [Advent of Code](https://adventofcode.com) written in
Go. Each year is kept in a separate folder, where days are grouped into folders.
Things that are shared between days are placed inside the _internal_ directory.

The repository is based on this template: <https://github.com/chwallen/go-aoc-template>.

## Requirements

Go 1.25 or later is required.

lp_solve is required for solving 2025/day10. If you don't want to install it,
the solution directory must either be removed or the solution to part two has
to be replaced with or removing the solution folder for 2025/day10.

### Installing lp_solve

Downloads for lp_solve can be found at <https://sourceforge.net/projects/lpsolve/files/lpsolve>.
However, it is also often available via more convenient means.

#### Linux

For Linux, the lp-solve package is found in the core repositories for many
distributions. For example:

- Ubuntu/Debian: `apt install liblpsolve55-dev`
- Arch Linux: `pacman -S lpsolve`

Set this variable so the Go compiler can find the files:
`export CGO_CFLAGS="-I/usr/include/lpsolve"`. Adjust as appropriate if your
distribution installs the files somewhere else. For Ubuntu/Debian, you may
also have to set this: `CGO_LDFLAGS="-llpsolve55 -lm -ldl"`.

#### Mac

For Mac, it is available via Homebrew: `brew install lp_solve`

Set these variables so the Go compiler can find the files:

```sh
export CGO_CFLAGS="-I/opt/homebrew/opt/lp_solve/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/lp_solve/lib"
```

#### Windows

For Windows, you must download a _dev_ zip from
<https://sourceforge.net/projects/lpsolve/files/lpsolve>, i.e.,
_lp\_solve\_5.5.2.11\_dev\_win64.zip_. How to make it compile left as an
exercise.
