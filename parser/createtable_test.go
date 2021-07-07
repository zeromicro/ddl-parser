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

func TestVisitor_VisitCreateTable(t *testing.T) {
	p := NewParser(WithDebugMode(true))
	accept := func(p *gen.MySqlParser, visitor *Visitor) interface{} {
		ctx := p.CreateTable()
		return visitor.VisitCreateTable(ctx)
	}

	t.Run("copyCreateTableContext", func(t *testing.T) {
		_, err := p.testMysqlSyntax("test.sql", accept,
			`create table new_t  (like t1);`)
		assert.Error(t, err)
	})

	t.Run("queryCreateTable", func(t *testing.T) {
		_, err := p.testMysqlSyntax("test.sql", accept,
			`CREATE TABLE test (a INT NOT NULL AUTO_INCREMENT,PRIMARY KEY (a), 
				KEY(b))ENGINE=InnoDB SELECT b,c FROM test2;`)
		assert.Error(t, err)
	})

	t.Run("columnCreateTable", func(t *testing.T) {
		v, err := p.testMysqlSyntax("test.sql", accept,
			"CREATE TABLE `user` (\n  "+
				"`id` bigint NOT NULL AUTO_INCREMENT,\n  "+
				"`number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',\n  "+
				"`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名称',\n "+
				" `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',\n "+
				" `gender` char(5) NOT NULL COMMENT '男｜女｜未公开',\n  "+
				"`create_time` timestamp NULL DEFAULT NULL,\n  "+
				"`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  "+
				"PRIMARY KEY (`id`),\n  "+
				"UNIQUE KEY `number_unique` (`number`) USING BTREE,\n  "+
				"UNIQUE KEY `number_unique2` (`number`"+
				") USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
		assert.Nil(t, err)
		_, ok := v.(*CreateTable)
		assert.True(t, ok)
		assert.Equal(t, true, func() bool {
			return true
		})
	})
}
