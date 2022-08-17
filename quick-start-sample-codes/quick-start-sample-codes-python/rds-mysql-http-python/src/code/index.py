# -*- coding: utf-8 -*-
import logging
from re import S
import json
import pymysql
import os

logger = logging.getLogger()

connection = None

# initialize hook, connect to the database
def initialize(context):
    global connection
    try:
        connection = pymysql.connect(
            # Replace it with your host name.
            host=os.environ['MYSQL_ENDPOINT'],
            # Replace it with your port number.
            port=int(os.environ['MYSQL_PORT']),
            # Replace it with your username.
            user=os.environ['MYSQL_USER'],
            # Replace the  password with the one corresponding to your username.
            passwd=os.environ['MYSQL_PASSWORD'],
            # Replace it with the name of your database.
            db=os.environ['MYSQL_DBNAME'],
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
    logger.info("pre_stop hook finish.")

def handler(environ, start_response):
    if connection is None:
        raise Exception("Mysql connection not initialized.")

    # Check if the server is alive.
    # If the connection is closed, reconnect.
    connection.ping(reconnect=True)

    #get request_body
    try:
        request_body_size = int(environ.get('CONTENT_LENGTH', 0))
    except (ValueError):
        request_body_size = 0
    request_body = environ['wsgi.input'].read(request_body_size)
    request_body_json=json.loads(request_body)
    try:
        # Database write operation
        with connection.cursor() as cursor:
            # the primary key id is self-incrementing
            sql = "INSERT INTO `USERS` (`NAME`, `AGE`) VALUES (%s, %s)"
            cursor.execute(sql, (request_body_json['name'], request_body_json['age']))
            # connection is not autocommit by default. So you must commit to save your changes.
            connection.commit()
            logger.info(
                "Successfully insert a piece of data into the database")
        # Database read operation
        with connection.cursor() as cursor:
            cursor.execute("SELECT * FROM USERS ORDER BY ID DESC")
            result = cursor.fetchone()
            logger.info(result)
            start_response('200 OK', [('Content-type', 'text/plain')])
            return "user: {}".format(result)
    except Exception as e:
        logger.error(
            "ERROR: Unexpected error: Could not connect to MySql instance.")
        raise Exception(str(e))
