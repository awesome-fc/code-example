'use strict';

// more information about nodejs mysql: https://github.com/mysqljs/mysql
var mysql = require('mysql');
var connection;

exports.initialize = (context, callback) => {
  console.log('initializing');
  connection = mysql.createConnection({
    host: process.env.MYSQL_ENDPOINT,
    user: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    port: process.env.MYSQL_PORT,
    database: process.env.MYSQL_DBNAME
  });
  connection.connect((err) => {
    if (err) {
      console.log('[MYSQL CONNECTION ERROR] - ', err.message);
      callback(err)
      return;
    }
    callback(null, 'succ');

  });

};

exports.handler = (event, context, callback) => {
  //insert data
  var addSql = 'INSERT INTO USERS(NAME,AGE) VALUES(?,?)';
  var addSqlParams = ['wanger', '38'];
  connection.query(addSql, addSqlParams, function (err, result) {
    if (err) {
      console.log('[INSERT ERROR] - ', err.message);
      callback(err)
    }
    console.log('INSERT ID:', result);
  });
  //query data
  var sql = 'SELECT * FROM USERS ORDER BY ID DESC LIMIT 1';
  connection.query(sql, function (err, result) {
    if (err) {
      console.log('[SELECT ERROR] - ', err.message);
      callback(err)
    }
    console.log(result);
    callback(null, result);
  });
}

exports.pre_stop = (context, callback) => {
  console.log('pre_stop start');
  connection.end();
  console.log('pre_stop end');
  callback(null, '');
}