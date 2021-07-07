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

type CreateTable struct {
	// Name describes the literal of table
	Name        string
	Columns     []*ColumnDeclaration
	Constraints []*TableConstraint
}

type ColumnDeclaration struct {
	Name             string
	ColumnDefinition *ColumnDefinition
}

type createDefinition struct {
	ColumnDeclaration *ColumnDeclaration
	TableConstraint   *TableConstraint
}

// VisitCreateTable visits a parse tree produced by MySqlParser#createTable.
func (v *Visitor) VisitCreateTable(ctx gen.ICreateTableContext) *CreateTable {
	switch tx := ctx.(type) {
	case *gen.CopyCreateTableContext:
		v.panicWithExpr(tx.GetStart(), "Unsupported creating a table by copying from another table")
	case *gen.QueryCreateTableContext:
		v.panicWithExpr(tx.GetStart(), "Unsupported creating a table by querying from another table")
	case *gen.ColumnCreateTableContext:
		return v.VisitColumnCreateTable(tx)
	}

	v.panicWithExpr(ctx.GetStart(), "Unknown creating a table")
	return nil
}

// VisitColumnCreateTable visits a parse tree produced by MySqlParser#columnCreateTable.
func (v *Visitor) VisitColumnCreateTable(ctx *gen.ColumnCreateTableContext) *CreateTable {
	var ret CreateTable
	tableName := ctx.TableName().GetText()
	tableName = strings.Trim(tableName, "`")
	tableName = strings.Trim(tableName, "'")
	tableName = strings.NewReplacer("\r", "", "\n", "").Replace(tableName)
	ret.Name = tableName
	if ctx.CreateDefinitions() != nil {
		if createDefinitionsContext, ok := ctx.CreateDefinitions().(*gen.CreateDefinitionsContext); ok {
			definitions := v.VisitCreateDefinitions(createDefinitionsContext)
			v.convertCreateDefinition(definitions, &ret)
		}
	}

	return &ret
}

// VisitCreateDefinitions visits a parse tree produced by MySqlParser#createDefinitions.
func (v *Visitor) VisitCreateDefinitions(ctx *gen.CreateDefinitionsContext) []*createDefinition {
	var ret []*createDefinition
	for _, e := range ctx.AllCreateDefinition() {
		data := v.VisitCreateDefinition(e)
		if data == nil {
			continue
		}

		switch r := data.(type) {
		case *ColumnDeclaration:
			ret = append(ret, &createDefinition{
				ColumnDeclaration: r,
			})
		case *TableConstraint:
			ret = append(ret, &createDefinition{
				TableConstraint: r,
			})
		}
	}
	return ret
}

// VisitCreateDefinition visits a parse tree produced by MySqlParser#createDefinition.
func (v *Visitor) VisitCreateDefinition(ctx gen.ICreateDefinitionContext) interface{} {
	switch tx := ctx.(type) {
	case *gen.ColumnDeclarationContext:
		var ret ColumnDeclaration
		ret.Name = v.visitUid(tx.Uid())
		iDefinitionContext := tx.ColumnDefinition()
		definitionContext, ok := iDefinitionContext.(*gen.ColumnDefinitionContext)
		if ok {
			out := v.VisitColumnDefinition(definitionContext)
			if cd, ok := out.(*ColumnDefinition); ok {
				ret.ColumnDefinition = cd
			}
		}

		return &ret
	case *gen.ConstraintDeclarationContext:
		if tx.TableConstraint() != nil {
			if constraintCtx, ok := tx.TableConstraint().(*gen.TableConstraintContext); ok {
				return v.VisitTableConstraint(constraintCtx)
			}
		}
	}

	return nil
}

func (v *Visitor) convertCreateDefinition(list []*createDefinition, table *CreateTable) {
	for _, e := range list {
		if e.ColumnDeclaration != nil {
			table.Columns = append(table.Columns, e.ColumnDeclaration)
		}
		if e.TableConstraint != nil {
			table.Constraints = append(table.Constraints, e.TableConstraint)
		}
	}
}
