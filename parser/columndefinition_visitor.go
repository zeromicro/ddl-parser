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

type ColumnDefinition struct {
	DataType         DataType
	ColumnConstraint *ColumnConstraint
}

type ColumnConstraint struct {
	NotNull         bool
	HasDefaultValue bool
	AutoIncrement   bool
	Primary         bool
	Key             bool
	Unique          bool
	Comment         string
}

type key bool
type primary bool

// VisitColumnDefinition visits a parse tree produced by MySqlParser#columnDefinition.
func (v *Visitor) VisitColumnDefinition(ctx *gen.ColumnDefinitionContext) interface{} {
	v.trace("VisitColumnDefinition")

	var (
		constraint ColumnConstraint
		out        ColumnDefinition
	)
	out.DataType = v.VisitDataType(ctx.DataType())
	for _, e := range ctx.AllColumnConstraint() {
		switch tx := e.(type) {
		case *gen.NullColumnConstraintContext:
			constraint.NotNull = v.VisitNullColumnConstraint(tx)
		case *gen.DefaultColumnConstraintContext:
			constraint.HasDefaultValue = v.VisitDefaultColumnConstraint(tx)
		case *gen.AutoIncrementColumnConstraintContext:
			constraint.AutoIncrement = v.VisitAutoIncrementColumnConstraint(tx)
		case *gen.PrimaryKeyColumnConstraintContext:
			ret := v.VisitPrimaryKeyColumnConstraint(tx)
			if c, ok := ret.(*primary); ok {
				constraint.Primary = bool(*c)
			} else {
				c := ret.(*key)
				constraint.Key = bool(*c)
			}
		case *gen.UniqueKeyColumnConstraintContext:
			constraint.Unique = v.VisitUniqueKeyColumnConstraint(tx)
		case *gen.CommentColumnConstraintContext:
			constraint.Comment = v.VisitCommentColumnConstraint(tx)
		case *gen.ReferenceColumnConstraintContext:
			v.panicWithExpr(tx.GetStart(), "Unsupported reference definition")
		}
	}
	out.ColumnConstraint = &constraint
	return &out
}

// VisitNullColumnConstraint visits a parse tree produced by MySqlParser#nullColumnConstraint.
func (v *Visitor) VisitNullColumnConstraint(ctx *gen.NullColumnConstraintContext) bool {
	v.trace("VisitNullColumnConstraint")
	if ret, ok := ctx.NullNotnull().(*gen.NullNotnullContext); ok {
		return v.VisitNullNotnull(ret)
	}

	return false
}

// VisitDefaultColumnConstraint visits a parse tree produced by MySqlParser#defaultColumnConstraint.
func (v *Visitor) VisitDefaultColumnConstraint(ctx *gen.DefaultColumnConstraintContext) bool {
	v.trace("VisitDefaultColumnConstraint")
	text := ctx.DefaultValue().GetText()
	text = strings.Trim(text, "`")
	text = strings.Trim(text, "'")
	text = strings.NewReplacer("\r", "", "\n", "").Replace(text)
	if strings.ToUpper(text) == "NULL" || len(text) == 0 {
		return false
	}

	return true
}

// VisitAutoIncrementColumnConstraint visits a parse tree produced by MySqlParser#autoIncrementColumnConstraint.
func (v *Visitor) VisitAutoIncrementColumnConstraint(_ *gen.AutoIncrementColumnConstraintContext) bool {
	v.trace("VisitAutoIncrementColumnConstraint")
	return true
}

// VisitPrimaryKeyColumnConstraint visits a parse tree produced by MySqlParser#primaryKeyColumnConstraint.
func (v *Visitor) VisitPrimaryKeyColumnConstraint(ctx *gen.PrimaryKeyColumnConstraintContext) interface{} {
	v.trace("VisitPrimaryKeyColumnConstraint")
	if ctx.PRIMARY() == nil {
		var ret key
		ret = true
		return &ret
	}

	var ret primary
	ret = true
	return &ret
}

// VisitUniqueKeyColumnConstraint visits a parse tree produced by MySqlParser#uniqueKeyColumnConstraint.
func (v *Visitor) VisitUniqueKeyColumnConstraint(_ *gen.UniqueKeyColumnConstraintContext) bool {
	v.trace("VisitUniqueKeyColumnConstraint")
	return true
}

// VisitCommentColumnConstraint visits a parse tree produced by MySqlParser#commentColumnConstraint.
func (v *Visitor) VisitCommentColumnConstraint(ctx *gen.CommentColumnConstraintContext) string {
	v.trace("VisitCommentColumnConstraint")
	value := parseTerminalNode(ctx.STRING_LITERAL(), withTrim("`"), withTrim(`"`), withTrim(`'`), withReplacer(`\r`, "", `\n`, ""))
	return value
}

// VisitNullNotnull visits a parse tree produced by MySqlParser#nullNotnull.
func (v *Visitor) VisitNullNotnull(ctx *gen.NullNotnullContext) bool {
	v.trace("VisitNullNotnull")
	if ctx.NOT() != nil {
		return true
	}

	return false
}
