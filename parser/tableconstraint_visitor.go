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
	"strings"

	"github.com/zeromicro/ddl-parser/gen"
)

type TableConstraint struct {
	// ColumnPrimaryKey describes the name of columns
	ColumnPrimaryKey []string
	// ColumnUniqueKey describes the name of columns
	ColumnUniqueKey []string
}

// VisitTableConstraint visits a parse tree produced by MySqlParser#tableConstraint.
func (v *Visitor) VisitTableConstraint(ctx gen.ITableConstraintContext) *TableConstraint {
	var ret TableConstraint
	switch tx := ctx.(type) {
	case *gen.PrimaryKeyTableConstraintContext:
		if tx.IndexColumnNames() != nil {
			indexColumnNamesCtx, ok := tx.IndexColumnNames().(*gen.IndexColumnNamesContext)
			if ok {
				ret.ColumnPrimaryKey = v.VisitIndexColumnNames(indexColumnNamesCtx)
			}
		}
	case *gen.UniqueKeyTableConstraintContext:
		if tx.IndexColumnNames() != nil {
			indexColumnNamesCtx, ok := tx.IndexColumnNames().(*gen.IndexColumnNamesContext)
			if ok {
				ret.ColumnUniqueKey = v.VisitIndexColumnNames(indexColumnNamesCtx)
			}
		}
	case *gen.ForeignKeyTableConstraintContext:
		v.panicWithExpr(tx.GetStart(), "Unsupported foreign key constraint")
	}

	return &ret
}

// VisitIndexColumnNames visits a parse tree produced by MySqlParser#indexColumnNames.
func (v *Visitor) VisitIndexColumnNames(ctx *gen.IndexColumnNamesContext) []string {
	var columns []string
	for _, e := range ctx.AllIndexColumnName() {
		indexCtx, ok := e.(*gen.IndexColumnNameContext)
		if !ok {
			continue
		}

		columns = append(columns, v.VisitIndexColumnName(indexCtx))
	}

	return columns
}

// VisitIndexColumnName visits a parse tree produced by MySqlParser#indexColumnName.
func (v *Visitor) VisitIndexColumnName(ctx *gen.IndexColumnNameContext) string {
	var column string
	if ctx.Uid() != nil {
		column = v.visitUid(ctx.Uid())
	} else {
		column = parseTerminalNode(ctx.STRING_LITERAL(), withTrim("`"), withTrim("'"), withReplacer("\r", "", "\n", ""))
	}

	return column
}

func (v *Visitor) visitUid(ctx gen.IUidContext) string {
	str := ctx.GetText()
	str = strings.Trim(str, "`")
	str = strings.Trim(str, "'")
	str = strings.NewReplacer("\r", "", "\n", "").Replace(str)
	return str
}
