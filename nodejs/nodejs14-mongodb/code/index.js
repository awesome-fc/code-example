'use strict';

const { MongoClient } = require("mongodb");

var client;

exports.initialize = (context, callback) => {
  client = new MongoClient(process.env.MONGO_URL);
  callback(null, "succ");
};

exports.handler = (event, context, callback) => {
  client.connect(err => {
    if (err) {
      callback(err); 
      throw err;
    }
    const database = client.db(process.env.MONGO_DATABASE);
    const collections = database.collection("users");
    collections.find({"name":"张三"}).toArray((err, result) => {
      if (err) {
        callback(err); 
        throw err;
      }
      callback(null, result);
      client.close()
    })
  })
}

exports.pre_stop = (context, callback) => {
  client.close();
  callback(null, "")
}