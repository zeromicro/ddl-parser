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

const (
	invalidDataType DataType = iota
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

// DataType describes the data type of the column in table
type DataType int

// VisitDataType visits data type by switch-case
func (v *Visitor) VisitDataType(ctx gen.IDataTypeContext) DataType {
	switch t := ctx.(type) {
	case *gen.StringDataTypeContext:
		return v.VisitStringDataType(t)
	case *gen.NationalStringDataTypeContext:
		return v.VisitNationalStringDataType(t)
	case *gen.NationalVaryingStringDataTypeContext:
		return v.VisitNationalVaryingStringDataType(t)
	case *gen.DimensionDataTypeContext:
		return v.VisitDimensionDataType(t)
	case *gen.SimpleDataTypeContext:
		return v.VisitSimpleDataType(t)
	case *gen.CollectionDataTypeContext:
		return v.VisitCollectionDataType(t)
	case *gen.SpatialDataTypeContext:
		return v.VisitSpatialDataType(t)
	case *gen.LongVarcharDataTypeContext:
		return v.VisitLongVarcharDataType(t)
	case *gen.LongVarbinaryDataTypeContext:
		return v.VisitLongVarbinaryDataType(t)
	}

	v.panicWithExpr(ctx.GetStart(), "invalid data type: "+ctx.GetText())
	return invalidDataType
}

// VisitStringDataType visits a parse tree produced by MySqlParser#stringDataType.
func (v *Visitor) VisitStringDataType(ctx *gen.StringDataTypeContext) DataType {
	v.trace(`VisitStringDataType`)
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `CHAR`:
		return Char
	case `CHARACTER`:
		return Character
	case `VARCHAR`:
		return VarChar
	case `TINYTEXT`:
		return TinyText
	case `TEXT`:
		return Text
	case `MEDIUMTEXT`:
		return MediumText
	case `LONGTEXT`:
		return LongText
	case `NCHAR`:
		return NChar
	case `NVARCHAR`:
		return NVarChar
	case `LONG`:
		return LongVarChar
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitNationalStringDataType visits a parse tree produced by MySqlParser#nationalVaryingStringDataType.
func (v *Visitor) VisitNationalStringDataType(ctx *gen.NationalStringDataTypeContext) DataType {
	v.trace(`VisitNationalStringDataType`)
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `VARCHAR`:
		return NVarChar
	case `CHARACTER`:
		return NChar
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitNationalVaryingStringDataType visits a parse tree produced by MySqlParser#nationalVaryingStringDataType.
func (v *Visitor) VisitNationalVaryingStringDataType(_ *gen.NationalVaryingStringDataTypeContext) DataType {
	v.trace("VisitNationalVaryingStringDataType")
	return NVarChar
}

// VisitDimensionDataType visits a parse tree produced by MySqlParser#dimensionDataType.
func (v *Visitor) VisitDimensionDataType(ctx *gen.DimensionDataTypeContext) DataType {
	v.trace("VisitDimensionDataType")
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `BIT`:
		return Bit
	case `TIME`:
		return Time
	case `TIMESTAMP`:
		return Timestamp
	case `DATETIME`:
		return DateTime
	case `BINARY`:
		return Binary
	case `VARBINARY`:
		return VarBinary
	case `BLOB`:
		return Blob
	case `YEAR`:
		return Year
	case `DECIMAL`:
		return Decimal
	case `DEC`:
		return Dec
	case `FIXED`:
		return Fixed
	case `NUMERIC`:
		return Numeric
	case `FLOAT`:
		return Float
	case `FLOAT4`:
		return Float4
	case `FLOAT8`:
		return Float8
	case `DOUBLE`:
		return Double
	case `REAL`:
		return Real
	case `TINYINT`:
		return TinyInt
	case `SMALLINT`:
		return SmallInt
	case `MEDIUMINT`:
		return MediumInt
	case `INT`:
		return Int
	case `INTEGER`:
		return Integer
	case `BIGINT`:
		return BigInt
	case `MIDDLEINT`:
		return MiddleInt
	case `INT1`:
		return Int1
	case `INT2`:
		return Int2
	case `INT3`:
		return Int3
	case `INT4`:
		return Int4
	case `INT8`:
		return Int8
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitSimpleDataType visits a parse tree produced by MySqlParser#simpleDataType.
func (v *Visitor) VisitSimpleDataType(ctx *gen.SimpleDataTypeContext) DataType {
	v.trace("VisitSimpleDataType")
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `DATE`:
		return Date
	case `TINYBLOB`:
		return TinyBlob
	case `MEDIUMBLOB`:
		return MediumBlob
	case `LONGBLOB`:
		return LongBlob
	case `BOOL`:
		return Bool
	case `BOOLEAN`:
		return Boolean
	case `SERIAL`:
		return Serial
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitCollectionDataType visits a parse tree produced by MySqlParser#collectionDataType.
// todo(anqiansong) enum/set value
func (v *Visitor) VisitCollectionDataType(ctx *gen.CollectionDataTypeContext) DataType {
	v.trace("VisitCollectionDataType")
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `ENUM`:
		return Enum
	case `SET`:
		return Set
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitSpatialDataType visits a parse tree produced by MySqlParser#spatialDataType.
func (v *Visitor) VisitSpatialDataType(ctx *gen.SpatialDataTypeContext) DataType {
	v.trace("VisitSpatialDataType")
	text := strings.ToUpper(ctx.GetTypeName().GetText())
	switch text {
	case `GEOMETRYCOLLECTION`:
		return GeometryCollection
	case `GEOMCOLLECTION`:
		return GeomCollection
	case `LINESTRING`:
		return LineString
	case `MULTILINESTRING`:
		return MultiLineString
	case `MULTIPOINT`:
		return MultiPoint
	case `MULTIPOLYGON`:
		return MultiPolygon
	case `POINT`:
		return Point
	case `POLYGON`:
		return Polygon
	case `JSON`:
		return Json
	case `GEOMETRY`:
		return Geometry
	}

	v.panicWithExpr(ctx.GetTypeName(), "invalid data type: "+text)
	return invalidDataType
}

// VisitLongVarcharDataType visits a parse tree produced by MySqlParser#longVarcharDataType.
func (v *Visitor) VisitLongVarcharDataType(_ *gen.LongVarcharDataTypeContext) DataType {
	v.trace("VisitLongVarcharDataType")
	return LongVarChar
}

// VisitLongVarbinaryDataType visits a parse tree produced by MySqlParser#longVarbinaryDataType.
func (v *Visitor) VisitLongVarbinaryDataType(_ *gen.LongVarbinaryDataTypeContext) DataType {
	v.trace("VisitLongVarbinaryDataType")
	return LongVarBinary
}
