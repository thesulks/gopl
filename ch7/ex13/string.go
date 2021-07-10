package eval

import (
	"fmt"
)

var depth = -1

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.6g", l)
}

func (u unary) String() string {
	depth++
	ast := fmt.Sprintf("%c\n%s%s", u.op,
		indent(), lastChildForm(u.x))
	depth--
	return ast
}

func (b binary) String() string {
	depth++
	ast := fmt.Sprintf("%c\n%s%s\n%s%s", b.op,
		indent(), childForm(b.x),
		indent(), lastChildForm(b.y))
	depth--
	return ast
}

func (c call) String() string {
	depth++
	ast := c.fn
	for i, e := range c.args {
		ast += "\n" + indent()
		if i == len(c.args)-1 {
			ast += lastChildForm(e)
		} else {
			ast += childForm(e)
		}
	}
	depth--
	return ast
}

func indent() string {
	return fmt.Sprintf("%*s", depth*4, "")
}

func childForm(e Expr) string {
	return fmt.Sprintf("├── %s", e)
}

func lastChildForm(e Expr) string {
	return fmt.Sprintf("└── %s", e)
}
