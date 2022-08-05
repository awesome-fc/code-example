'use strict';

const Kafka = require('node-rdkafka');

var bootstrap_servers
var topic_name
var producer

exports.initialize = (context, callback) => {
    bootstrap_servers = process.env.bootstrap_servers
    topic_name = process.env.topic_name
    console.log("Servers: ", bootstrap_servers);
    console.log("TopicName: ", topic_name);
    producer = new Kafka.Producer({
        'api.version.request': 'true',
        'bootstrap.servers': bootstrap_servers,
        'dr_cb': true,
        'dr_msg_cb': true,
        'message.send.max.retries': 10
    });

    var connected = false

    // Wait for the ready event before proceeding
    producer.on('ready', function() {
        connected = true
        console.log("connect ok")
    });

    producer.on("disconnected", function() {
        connected = false;
        // Auto reconnect
        producer.connect();
    })

    producer.on('event.log', function(event) {
        console.log("event.log", event);
        throw new Error(event);
    });

    producer.on("error", function(error) {
        console.log("error:" + error);
        throw new Error(error);
    });

    // Any errors we encounter, including connection errors
    producer.on('event.error', function(err) {
        console.error('event.error:' + err);
        throw new Error(err);
    })
    
    // Poll for events every 10 ms
    producer.setPollInterval(10);

    callback(null, "");
};

exports.handler = async(event, context, callback) => {
    // Connect to the broker manually
    producer.connect();

    // Wait for connection
    await producer.on('ready', function() {
        producer.produce(
            topic_name,   
            null,      
            Buffer.from(event),      
            null,   
            Date.now()
        );
        producer.flush();
    });

    // waiting for sending
    await producer.on('delivery-report', function(err, report) {
        console.log("delivery-report err: ", err);
        console.log("delivery-report content: ", report);

        if (err == null) {
            callback(null, "Finish sending the message:" + event);
        } else {
            callback(err, "Send message fail!");
        }  
    });
}