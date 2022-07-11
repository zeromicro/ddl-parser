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
	"github.com/zeromicro/ddl-parser/gen"
)

const (
	_ int = iota
	LongVarBinary
	LongVarChar
	GeometryCollection
	GeomCollection
	LineString
	MultiLineString
	MultiPoint
	MultiPolygon
	Point
	Polygon
	Json
	Geometry
	Enum
	Set
	Bit
	Time
	Timestamp
	DateTime
	Binary
	VarBinary
	Blob
	Year
	Decimal
	Dec
	Fixed
	Numeric
	Float
	Float4
	Float8
	Double
	Real
	TinyInt
	SmallInt
	MediumInt
	Int
	Integer
	BigInt
	MiddleInt
	Int1
	Int2
	Int3
	Int4
	Int8
	Date
	TinyBlob
	MediumBlob
	LongBlob
	Bool
	Boolean
	Serial
	NVarChar
	NChar
	Char
	Character
	VarChar
	TinyText
	Text
	MediumText
	LongText
)

// DataType describes the data type and value of the column in table
type DataType interface {
	Type() int
	Unsigned() bool
	// Value returns the values if the data type is Enum or Set
	Value() []string
}

var _ DataType = (*NormalDataType)(nil)
var _ DataType = (*EnumSetDataType)(nil)

// NormalDataType describes the data type which not contains Enum and Set of column
type NormalDataType struct {
	tp       int
	unsigned bool
}

// Unsigned returns true if the data type is unsigned.
func (n *NormalDataType) Unsigned() bool {
	return n.unsigned
}

// Type returns the data type of column
func (n *NormalDataType) Type() int {
	return n.tp
}

// Value returns nil default
func (n *NormalDataType) Value() []string {
	return nil
}

func with(tp int, unsigned bool, value ...string) DataType {
	if len(value) > 0 {
		return &EnumSetDataType{
			tp:    tp,
			value: value,
		}
	}
	return &NormalDataType{tp: tp, unsigned: unsigned}
}

// EnumSetDataType describes the data type  Enum and Set of column
type EnumSetDataType struct {
	tp    int
	value []string
}

// Type returns the data type of column
func (e *EnumSetDataType) Type() int {
	return e.tp
}

// Unsigned returns true if the data type is unsigned.
func (e *EnumSetDataType) Unsigned() bool {
	return false
}

// Value returns the value of data type Enum and Set
func (e *EnumSetDataType) Value() []string {
	return e.value
}

// visitDataType visits data type by switch-case
func (v *visitor) visitDataType(ctx gen.IDataTypeContext) DataType {
	v.trace("VisitDataType")
	switch t := ctx.(type) {
	case *gen.StringDataTypeContext:
		return v.visitStringDataType(t)
	case *gen.NationalStringDataTypeContext:
		return v.visitNationalStringDataType(t)
	case *gen.NationalVaryingStringDataTypeContext:
		return v.visitNationalVaryingStringDataType(t)
	case *gen.DimensionDataTypeContext:
		return v.visitDimensionDataType(t)
	case *gen.SimpleDataTypeContext:
		return v.visitSimpleDataType(t)
	case *gen.CollectionDataTypeContext:
		return v.visitCollectionDataType(t)
	case *gen.SpatialDataTypeContext:
		return v.visitSpatialDataType(t)
	case *gen.LongVarcharDataTypeContext:
		return v.visitLongVarcharDataType(t)
	case *gen.LongVarbinaryDataTypeContext:
		return v.visitLongVarbinaryDataType(t)
	}

	v.panicWithExpr(ctx.GetStart(), "invalid data type: "+ctx.GetText())
	return nil
}

// visitStringDataType visits a parse tree produced by MySqlParser#stringDataType.
func (v *visitor) visitStringDataType(ctx *gen.StringDataTypeContext) DataType {
	v.trace(`VisitStringDataType`)
	text := parseToken(ctx.GetTypeName(), withUpperCase(), withTrim("`"))
	switch text {
	case `CHAR`:
		return with(Char, false)
	case `CHARACTER`:
		return with(Character, false)
	case `VARCHAR`:
		return with(VarChar, false)
	case `TINYTEXT`:
		return with(TinyText, false)
	case `TEXT`:
		return with(Text, false)
	case `MEDIUMTEXT`:
		return with(MediumText, false)
	case `LONGTEXT`:
		return with(LongText, false)
	case `NCHAR`:
		return with(NChar, false)
	case `NVARCHAR`:
		return with(NVarChar, false)
	case `LONG`:
		return with(LongVarChar, false)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitNationalStringDataType visits a parse tree produced by MySqlParser#nationalVaryingStringDataType.
func (v *visitor) visitNationalStringDataType(ctx *gen.NationalStringDataTypeContext) DataType {
	v.trace(`VisitNationalStringDataType`)
	text := parseToken(ctx.GetTypeName(), withUpperCase(), withTrim("`"))
	switch text {
	case `VARCHAR`:
		return with(NVarChar, false)
	case `CHARACTER`:
		return with(NChar, false)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitNationalVaryingStringDataType visits a parse tree produced by MySqlParser#nationalVaryingStringDataType.
func (v *visitor) visitNationalVaryingStringDataType(_ *gen.NationalVaryingStringDataTypeContext) DataType {
	v.trace("VisitNationalVaryingStringDataType")
	return with(NVarChar, false)
}

// visitDimensionDataType visits a parse tree produced by MySqlParser#dimensionDataType.
func (v *visitor) visitDimensionDataType(ctx *gen.DimensionDataTypeContext) DataType {
	v.trace("VisitDimensionDataType")
	text := parseToken(ctx.GetTypeName(), withUpperCase(), withTrim("`"))
	unsigned := ctx.UNSIGNED() != nil
	switch text {
	case `BIT`:
		return with(Bit, unsigned)
	case `TIME`:
		return with(Time, unsigned)
	case `TIMESTAMP`:
		return with(Timestamp, unsigned)
	case `DATETIME`:
		return with(DateTime, unsigned)
	case `BINARY`:
		return with(Binary, unsigned)
	case `VARBINARY`:
		return with(VarBinary, unsigned)
	case `BLOB`:
		return with(Blob, unsigned)
	case `YEAR`:
		return with(Year, unsigned)
	case `DECIMAL`:
		return with(Decimal, unsigned)
	case `DEC`:
		return with(Dec, unsigned)
	case `FIXED`:
		return with(Fixed, unsigned)
	case `NUMERIC`:
		return with(Numeric, unsigned)
	case `FLOAT`:
		return with(Float, unsigned)
	case `FLOAT4`:
		return with(Float4, unsigned)
	case `FLOAT8`:
		return with(Float8, unsigned)
	case `DOUBLE`:
		return with(Double, unsigned)
	case `REAL`:
		return with(Real, unsigned)
	case `TINYINT`:
		return with(TinyInt, unsigned)
	case `SMALLINT`:
		return with(SmallInt, unsigned)
	case `MEDIUMINT`:
		return with(MediumInt, unsigned)
	case `INT`:
		return with(Int, unsigned)
	case `INTEGER`:
		return with(Integer, unsigned)
	case `BIGINT`:
		return with(BigInt, unsigned)
	case `MIDDLEINT`:
		return with(MiddleInt, unsigned)
	case `INT1`:
		return with(Int1, unsigned)
	case `INT2`:
		return with(Int2, unsigned)
	case `INT3`:
		return with(Int3, unsigned)
	case `INT4`:
		return with(Int4, unsigned)
	case `INT8`:
		return with(Int8, unsigned)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitSimpleDataType visits a parse tree produced by MySqlParser#simpleDataType.
func (v *visitor) visitSimpleDataType(ctx *gen.SimpleDataTypeContext) DataType {
	v.trace("VisitSimpleDataType")
	text := parseToken(
		ctx.GetTypeName(),
		withUpperCase(),
		withTrim("`"),
	)

	switch text {
	case `DATE`:
		return with(Date, false)
	case `TINYBLOB`:
		return with(TinyBlob, false)
	case `MEDIUMBLOB`:
		return with(MediumBlob, false)
	case `LONGBLOB`:
		return with(LongBlob, false)
	case `BOOL`:
		return with(Bool, false)
	case `BOOLEAN`:
		return with(Boolean, false)
	case `SERIAL`:
		return with(Serial, false)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitCollectionDataType visits a parse tree produced by MySqlParser#collectionDataType.
func (v *visitor) visitCollectionDataType(ctx *gen.CollectionDataTypeContext) DataType {
	v.trace("VisitCollectionDataType")
	text := parseToken(
		ctx.GetTypeName(),
		withUpperCase(),
		withTrim("`"),
	)

	var values []string
	if ctx.CollectionOptions() != nil {
		optionsCtx, ok := ctx.CollectionOptions().(*gen.CollectionOptionsContext)
		if ok {
			for _, e := range optionsCtx.AllSTRING_LITERAL() {
				value := parseTerminalNode(
					e, withTrim("`"),
					withTrim(`"`),
					withTrim(`'`),
				)
				values = append(values, value)
			}
		}
	}

	switch text {
	case `ENUM`:
		return with(Enum, false, values...)
	case `SET`:
		return with(Set, false, values...)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitSpatialDataType visits a parse tree produced by MySqlParser#spatialDataType.
func (v *visitor) visitSpatialDataType(ctx *gen.SpatialDataTypeContext) DataType {
	v.trace("VisitSpatialDataType")
	text := parseToken(
		ctx.GetTypeName(),
		withUpperCase(),
		withTrim("`"),
	)

	switch text {
	case `GEOMETRYCOLLECTION`:
		return with(GeometryCollection, false)
	case `GEOMCOLLECTION`:
		return with(GeomCollection, false)
	case `LINESTRING`:
		return with(LineString, false)
	case `MULTILINESTRING`:
		return with(MultiLineString, false)
	case `MULTIPOINT`:
		return with(MultiPoint, false)
	case `MULTIPOLYGON`:
		return with(MultiPolygon, false)
	case `POINT`:
		return with(Point, false)
	case `POLYGON`:
		return with(Polygon, false)
	case `JSON`:
		return with(Json, false)
	case `GEOMETRY`:
		return with(Geometry, false)
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return nil
}

// visitLongVarcharDataType visits a parse tree produced by MySqlParser#longVarcharDataType.
func (v *visitor) visitLongVarcharDataType(_ *gen.LongVarcharDataTypeContext) DataType {
	v.trace("VisitLongVarcharDataType")
	return with(LongVarChar, false)
}

// visitLongVarbinaryDataType visits a parse tree produced by MySqlParser#longVarbinaryDataType.
func (v *visitor) visitLongVarbinaryDataType(_ *gen.LongVarbinaryDataTypeContext) DataType {
	v.trace("VisitLongVarbinaryDataType")
	return with(LongVarBinary, false)
}
