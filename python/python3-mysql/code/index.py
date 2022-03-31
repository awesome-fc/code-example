# -*- coding: utf-8 -*-
import logging
from re import S
import pymysql
import os

logger = logging.getLogger()

connection = None

# initialize hook, connect to the database
def initialize(context):
    global connection
    try:
        connection = pymysql.connect(
            host=os.environ['MYSQL_ENDPOING'],   # 替换为您的HOST名称。
            port=int(os.environ['MYSQL_PORT']),  # 替换为您的端口号。
            user=os.environ['MYSQL_USER'],       # 替换为您的用户名。
            passwd=os.environ['MYSQL_PASSWORD'], # 替换为您的用户名对应的密码。
            db=os.environ['MYSQL_DBNAME'],       # 替换为您的数据库名称。
            connect_timeout=5)
    except Exception as e:
        logger.error(e)
        logger.error(
            "ERROR: Unexpected error: Could not connect to MySql instance.")
        raise Exception(str(e))

def pre_stop(context):
    logger.info("pre_stop hook start.")
    if connection != None:
        connection.close()


def handler(event, context):
    if connection is None:
        raise Exception("Mysql connection not initialized.")
    
    # Check if the server is alive.
    # If the connection is closed, reconnect.
    connection.ping()

    try:
        with connection.cursor() as cursor:
            cursor.execute("SELECT * FROM users")
            result = cursor.fetchone()
            logger.info(result)
            return "user: {}".format(result)
    except Exception as e:
        logger.error(
            "ERROR: Unexpected error: Could not connect to MySql instance.")
        raise Exception(str(e))


