package main

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

type candidate struct {
	start token.Pos
	end   token.Pos
	value string
}

type collector struct {
	threshold     int
	cands         []candidate
	skippedConst  int
	skippedSmall  int
	skippedDecode int
}

// collect appends compressible string literals of file to cands.
func (c *collector) collect(file *ast.File) {
	ast.Inspect(file, c.inspect(false))
}

func (c *collector) inspect(constCtx bool) func(ast.Node) bool {
	return func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if constCtx {
			c.countConst(n)
			return false
		}

		switch n := n.(type) {
		case *ast.GenDecl:
			if n.Tok == token.CONST {
				c.countConst(n)
				return false
			}
			return true

		case *ast.ImportSpec:
			return false

		case *ast.CaseClause:
			for _, e := range n.List {
				ast.Inspect(e, c.inspect(true))
			}
			for _, s := range n.Body {
				ast.Inspect(s, c.inspect(false))
			}
			return false

		case *ast.Field:
			if n.Tag != nil {
				c.countConst(n.Tag)
			}
			ast.Inspect(n.Type, c.inspect(false))
			return false

		case *ast.ArrayType:
			if n.Len != nil {
				ast.Inspect(n.Len, c.inspect(true))
			}
			ast.Inspect(n.Elt, c.inspect(false))
			return false

		case *ast.CallExpr:
			if isUnsafeBuiltin(n.Fun) {
				for _, a := range n.Args {
					ast.Inspect(a, c.inspect(true))
				}
			} else {
				for _, a := range n.Args {
					ast.Inspect(a, c.inspect(false))
				}
			}
			ast.Inspect(n.Fun, c.inspect(false))
			return false

		case *ast.BasicLit:
			if n.Kind == token.STRING {
				c.addLeaf(n)
			}
			return false

		case *ast.BinaryExpr:
			if n.Op == token.ADD {
				if parts, ok := foldStringConcat(n); ok {
					c.addChain(n, parts)
					return false
				}
			}
			return true
		}
		return true
	}
}

func (c *collector) countConst(n ast.Node) {
	ast.Inspect(n, func(m ast.Node) bool {
		if m == nil {
			return false
		}
		if lit, ok := m.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			c.skippedConst++
		}
		return true
	})
}

func (c *collector) addLeaf(lit *ast.BasicLit) {
	s, err := strconv.Unquote(lit.Value)
	if err != nil {
		c.skippedDecode++
		return
	}
	if len(s) < c.threshold {
		c.skippedSmall++
		return
	}
	c.cands = append(c.cands, candidate{start: lit.Pos(), end: lit.End(), value: s})
}

func (c *collector) addChain(n ast.Node, parts []string) {
	value := strings.Join(parts, "")
	if len(value) < c.threshold {
		c.skippedSmall++
		return
	}
	c.cands = append(c.cands, candidate{start: n.Pos(), end: n.End(), value: value})
}

// foldStringConcat returns the decoded parts of a pure string concat
// chain ("a" + "b" + "c", transparent through parentheses).
func foldStringConcat(n ast.Node) ([]string, bool) {
	switch n := n.(type) {
	case *ast.BasicLit:
		if n.Kind != token.STRING {
			return nil, false
		}
		s, err := strconv.Unquote(n.Value)
		if err != nil {
			return nil, false
		}
		return []string{s}, true

	case *ast.ParenExpr:
		return foldStringConcat(n.X)

	case *ast.BinaryExpr:
		if n.Op != token.ADD {
			return nil, false
		}
		x, okX := foldStringConcat(n.X)
		if !okX {
			return nil, false
		}
		y, okY := foldStringConcat(n.Y)
		if !okY {
			return nil, false
		}
		parts := make([]string, 0, len(x)+len(y))
		parts = append(parts, x...)
		parts = append(parts, y...)
		return parts, true
	}
	return nil, false
}

func isUnsafeBuiltin(n ast.Node) bool {
	sel, ok := n.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	ident, ok := sel.X.(*ast.Ident)
	return ok && ident.Name == "unsafe" &&
		(sel.Sel.Name == "Sizeof" || sel.Sel.Name == "Alignof" || sel.Sel.Name == "Offsetof")
}
