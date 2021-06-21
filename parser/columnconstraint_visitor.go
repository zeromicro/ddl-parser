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

import "github.com/zeromicro/ddl-parser/gen"

type ColumnConstraint struct {
	NotNull         bool
	HasDefaultValue bool
	AutoIncrement   bool
	Primary         bool
	Key             bool
	Unique          bool
	Comment         string
}

// VisitColumnDefinition visits a parse tree produced by MySqlParser#columnDefinition.
func (v *Visitor) VisitColumnDefinition(ctx *gen.ColumnDefinitionContext) *ColumnConstraint {
	return nil
}

// VisitNullColumnConstraint visits a parse tree produced by MySqlParser#nullColumnConstraint.
func (v *Visitor) VisitNullColumnConstraint(ctx *gen.NullColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitDefaultColumnConstraint visits a parse tree produced by MySqlParser#defaultColumnConstraint.
func (v *Visitor) VisitDefaultColumnConstraint(ctx *gen.DefaultColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitAutoIncrementColumnConstraint visits a parse tree produced by MySqlParser#autoIncrementColumnConstraint.
func (v *Visitor) VisitAutoIncrementColumnConstraint(ctx *gen.AutoIncrementColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitPrimaryKeyColumnConstraint visits a parse tree produced by MySqlParser#primaryKeyColumnConstraint.
func (v *Visitor) VisitPrimaryKeyColumnConstraint(ctx *gen.PrimaryKeyColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitUniqueKeyColumnConstraint visits a parse tree produced by MySqlParser#uniqueKeyColumnConstraint.
func (v *Visitor) VisitUniqueKeyColumnConstraint(ctx *gen.UniqueKeyColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitCommentColumnConstraint visits a parse tree produced by MySqlParser#commentColumnConstraint.
func (v *Visitor) VisitCommentColumnConstraint(ctx *gen.CommentColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitFormatColumnConstraint visits a parse tree produced by MySqlParser#formatColumnConstraint.
func (v *Visitor) VisitFormatColumnConstraint(ctx *gen.FormatColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitStorageColumnConstraint visits a parse tree produced by MySqlParser#storageColumnConstraint.
func (v *Visitor) VisitStorageColumnConstraint(ctx *gen.StorageColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitReferenceColumnConstraint visits a parse tree produced by MySqlParser#referenceColumnConstraint.
func (v *Visitor) VisitReferenceColumnConstraint(ctx *gen.ReferenceColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitCollateColumnConstraint visits a parse tree produced by MySqlParser#collateColumnConstraint.
func (v *Visitor) VisitCollateColumnConstraint(ctx *gen.CollateColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitGeneratedColumnConstraint visits a parse tree produced by MySqlParser#generatedColumnConstraint.
func (v *Visitor) VisitGeneratedColumnConstraint(ctx *gen.GeneratedColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitSerialDefaultColumnConstraint visits a parse tree produced by MySqlParser#serialDefaultColumnConstraint.
func (v *Visitor) VisitSerialDefaultColumnConstraint(ctx *gen.SerialDefaultColumnConstraintContext) *ColumnConstraint {
	return nil
}

// VisitCheckColumnConstraint visits a parse tree produced by MySqlParser#checkColumnConstraint.
func (v *Visitor) VisitCheckColumnConstraint(ctx *gen.CheckColumnConstraintContext) *ColumnConstraint {
	return nil
}
