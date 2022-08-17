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

exports.handler = function (request, response, context) {
  //get request body
  var body = request.body
  var bodyJson=JSON.parse(body)
  //insert data
  var addSql = 'INSERT INTO USERS(NAME,AGE) VALUES(?,?)';
  var addSqlParams = [bodyJson.name, bodyJson.age];
  connection.query(addSql, addSqlParams, function (err, result) {
    if (err) {
      console.log('[INSERT ERROR] - ', err.message);
    }
    console.log('INSERT ID:', result);
  });
  //query data
  var sql = 'SELECT * FROM USERS ORDER BY ID DESC LIMIT 1';
  connection.query(sql, function (err, result) {
    if (err) {
      console.log('[SELECT ERROR] - ', err.message);
    }
    console.log(result);
  });
  response.setStatusCode(200)
  response.setHeader('content-type', 'text/plain')
  response.send('succ')
}

exports.pre_stop = (context, callback) => {
  console.log('pre_stop start');
  connection.end();
  console.log('pre_stop end');
  callback(null, '');
}