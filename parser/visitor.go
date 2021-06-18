/*
 * MIT License
 *
 * Copyright (c) 2021 zeromicro
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package parser

import (
	"fmt"

	"github.com/zeromicro/ddl-parser/console"
	"github.com/zeromicro/ddl-parser/gen"
)

type Visitor struct {
	gen.BaseMySqlParserVisitor
	prefix string
	debug  bool
	logger console.Console
}

func (v *Visitor) trace(msg ...interface{}) {
	if v.debug {
		v.logger.Debug("Visit Trace: " + fmt.Sprint(msg...))
	}
}

func (v *Visitor) panic(line, column int, msg string) {
	if len(v.prefix) == 0 {
		panic(fmt.Errorf("%v:%v %s", line, column, msg))
		return
	}

	panic(fmt.Errorf("%v line %v:%v %s", v.prefix, line, column, msg))
}

func (v *Visitor) panicWithExpr(expr Token, msg string) {
	if len(v.prefix) == 0 {
		panic(fmt.Errorf("%v:%v %s", expr.GetLine(), expr.GetColumn(), msg))
		return
	}

	panic(fmt.Errorf("%v line %v:%v %s", v.prefix, expr.GetLine(), expr.GetColumn(), msg))
}

func (v *Visitor) panicWithExpected(expr Token, expected, actual string) {
	if len(v.prefix) == 0 {
		panic(fmt.Errorf("%v:%v expected %s, but found is %s", expr.GetLine(), expr.GetColumn(), expected, actual))
		return
	}

	panic(fmt.Errorf("%v line %v:%v expected %s, but found is %s", v.prefix, expr.GetLine(), expr.GetColumn(), expected, actual))
}
