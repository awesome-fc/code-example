'use strict';

const { MongoClient } = require("mongodb");

var client;

exports.initialize = (context, callback) => {
  client = new MongoClient(process.env.MONGO_URL);
  callback(null, "succ");
};

exports.handler = (request, response, context) => {
  var queries = request.queries
  client.connect(err => {
    if (err) {
      throw err;
    }
    const database = client.db(process.env.MONGO_DATABASE);
    const collections = database.collection("users");
    collections.find({"name":queries['name']}).toArray((err, result) => {
      if (err) {
        throw err;
      }
      response.setHeader('content-type', 'application/json');
      response.setStatusCode(200);
      response.send(Buffer.from(JSON.stringify(result)));
      client.close()
    })
  })
}

exports.pre_stop = (context, callback) => {
  client.close();
  callback(null, "")
}