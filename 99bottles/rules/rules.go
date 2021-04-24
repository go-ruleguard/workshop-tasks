package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func valueSwap(m dsl.Matcher) {
	m.Match(`$tmp := $x; $x = $y; $y = $tmp`).
		Suggest(`$x, $y = $y, $x`).
		Report(`could rewrite as parallel assignment: $x, $y = $y, $x`)
}

func assignOp(m dsl.Matcher) {
	m.Match(`$x = $x + $y`).
		Where(m["x"].Pure).
		Report(`could rewrite as '$x += $y'`)
}

func yodaExpr(m dsl.Matcher) {
	m.Match(`$lhs < $rhs`).
		Where(m["lhs"].Const && !m["rhs"].Const).
		Suggest(`$rhs > $lhs`).
		Report(`don't use yoda-style, rewrite as '$rhs > $lhs'`)

	m.Match(`$lhs > $rhs`).
		Where(m["lhs"].Const && !m["rhs"].Const).
		Suggest(`$rhs < $lhs`).
		Report(`don't use yoda-style, rewrite as '$rhs < $lhs'`)
}

func returnElse(m dsl.Matcher) {
	m.Match(`if $*_; $_ { $*_; return $*_; } else { $*_ }`).
		Report(`if block ends with return, so drop this else`)
}
