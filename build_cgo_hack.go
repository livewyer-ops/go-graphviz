// +build required

package graphviz

// This file exists purely to prevent the golang toolchain from stripping
// away the c source directories and files when `go mod vendor` is used
// to populate a `vendor/` directory of a project depending on `go-gl/glfw`.
//
// How it works:
//  - every directory which only includes c source files receives a dummy.go file.
//  - every directory we want to preserve is included here as a _ import.
//  - this file is given a build to exclude it from the regular build.
import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/livewyer-ops/go-graphviz/internal"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/ast"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/cdt"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/cgraph"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/circogen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/common"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/dotgen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/edgepaint"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/expr"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/fdpgen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/glcomp"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/gvc"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/gvpr"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/ingraphs"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/label"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/mingle"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/neatogen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/ortho"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/osage"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/pack"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/patchwork"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/pathplan"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/rbtree"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/sfdpgen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/sfio"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/sfio/Sfio_f"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/sparse"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/spine"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/topfish"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/twopigen"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/vmalloc"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/vpsc"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/vpsc/pairingheap"
	_ "github.com/livewyer-ops/go-graphviz/internal/ccall/xdot"
	_ "github.com/livewyer-ops/go-graphviz/internal/expat"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/core"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/devil"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/dot_layout"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/gd"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/gdiplus"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/gdk"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/glitz"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/gs"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/gtk"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/lasi"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/ming"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/neato_layout"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/pango"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/poppler"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/quartz"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/rsvg"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/visio"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/webp"
	_ "github.com/livewyer-ops/go-graphviz/internal/plugin/xlib"
)
