"""
本代码样例主要实现以下功能:
* 初始化 lindorm 连接
* 从环境变量中获取 lindorm 宽表引擎 table 名字
* 创建此表并向表中插入一条数据
* 读表中所有数据，查看是否符合预期

This sample code is mainly doing the following things:
* Initialize lindorm connection
* Get the lindorm table name from environment variables
* Create this table and insert data into the table
* Read all the data in the table to see if it is as expected

"""

# -*- coding: utf-8 -*-

import os
import json
import logging
import phoenixdb

logger = logging.getLogger()
connection = None

# initialize 做初始化连接工作，这样在一个实例中的多次请求可以复用同一个连接
# Initialize lindorm connection and multiple requests in an instance can reuse the same connection
def initialize(context):
    try:
        # 从环境变量获取初始化连接配置
        # Get initial connection configuration from environment variables
        env = os.environ
        lindorm_user_name = env.get("LindormUserName")
        lindorm_password = env.get("LindormPassword")
        database_url = env.get("DatabaseURL")

        # 初始化 Lindorm 连接
        # Initialize lindorm connection.
        connect_kw_args = {'user': lindorm_user_name, 'password': lindorm_password, 'readonly': 'True'}
        global connection
        connection = phoenixdb.connect(database_url, autocommit=True, **connect_kw_args)
    except Exception as e:
        logger.info(e)
        raise json.dumps({"success": False, "error_message": str(e)})
    return

# 删除表
# delete table
def delete_table(sql_table_name):
    with connection.cursor() as statement:
        offline_table = f"OFFLINE TABLE {sql_table_name}"
        statement.execute(offline_table)

        sql_drop_table = f"drop table if exists {sql_table_name}"
        statement.execute(sql_drop_table)
        statement.close()

def handler(event, context):
    if connection is None:
        raise Exception("lindorm connection not initialized.")

    env = os.environ
    sql_table_name = env.get("SQLTableName")
    try:
        with connection.cursor() as statement:
            # 如果输入的 table 存在，则先删除，保证表数据清洁
            # If the input table exists, delete it first to ensure that the table data is clean
            sql_drop_table = f"drop table if exists {sql_table_name}"
            statement.execute(sql_drop_table)

            # 创建新表，新表共有两列：c1、c2
            # Create a new table, the new table has two columns: c1, c2
            sql_create_table = f"create table if not exists {sql_table_name}(c1 int, c2 int, primary key(c1))"
            statement.execute(sql_create_table)

            # 向表中插入数据，两列数据分别为 20 和 30
            # Insert data into the table, the two columns of data are 20 and 30
            sql_upsert = f"upsert into {sql_table_name}(c1, c2) values(20,30)"
            statement.execute(sql_upsert)

            # 查询表中的数据并输出
            # Query the data in the table and output
            sql_select = f"SELECT * FROM {sql_table_name}"
            statement.execute(sql_select)
            rows = statement.fetchall()
            logger.info("rows data: " + str(rows))

            delete_table(sql_table_name)
        connection.close
    except Exception as e:
        logger.info(e)
        delete_table(sql_table_name)
        raise json.dumps({"success": False, "error_message": str(e)})

    return json.dumps({"success": True, "error_message": ""})