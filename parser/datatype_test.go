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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/ddl-parser/gen"
)

func TestDataType(t *testing.T) {
	p := NewParser(WithDebugMode(true))
	accept := func(p *gen.MySqlParser, visitor *Visitor) interface{} {
		return visitor.VisitDataType(p.DataType())
	}

	t.Run("stringDataType", func(t *testing.T) {
		testData := map[string]DataType{
			`CHAR(10)`:      Char,
			`CHARACTER(10)`: Character,
			`VARCHAR(10)`:   VarChar,
			`TINYTEXT`:      TinyText,
			`TEXT`:          Text,
			`MEDIUMTEXT`:    MediumText,
			`LONGTEXT`:      LongText,
			`NCHAR(20)`:     NChar,
			`NVARCHAR(20)`:  NVarChar,
			`LONG`:          Long,
		}

		for sql, dataType := range testData {
			actual, err := p.testMysqlSyntax("test.sql", accept, sql)
			assert.Nil(t, err)
			assertTypeEqual(t, dataType, actual)
		}
	})
}

func assertTypeEqual(t *testing.T, expected DataType, actual interface{}) {
	assert.Equal(t, expected, actual.(DataType))
}
