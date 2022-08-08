'use strict';

const Kafka = require('node-rdkafka');

var BOOTSTRAP_SERVERS
var TOPIC_NAME
var producer

exports.initialize = (context, callback) => {
    BOOTSTRAP_SERVERS = process.env.BOOTSTRAP_SERVERS
    TOPIC_NAME = process.env.TOPIC_NAME
    console.log("Servers: ", BOOTSTRAP_SERVERS);
    console.log("TopicName: ", TOPIC_NAME);
    producer = new Kafka.Producer({
        'api.version.request': 'true',
        'bootstrap.servers': BOOTSTRAP_SERVERS,
        'dr_cb': true,
        'dr_msg_cb': true,
        'message.send.max.retries': 10
    });

    producer.on('event.log', function(event) {
        console.log("event.log", event);
        callback(new Error(event.message), "");
    });

    producer.on("error", function(err) {
        console.log("error:" + err);
        callback(err, "");
    });

    // Any errors we encounter, including connection errors
    producer.on('event.error', function(err) {
        console.error('event.error:' + err);
        callback(err, "");
    })
    
    // Poll for events every 10 ms
    producer.setPollInterval(10);

    producer.connect();

    // Wait for the ready event before proceeding
    producer.on('ready', function() {
        console.log("connect ok")
        callback(null, "");
    });
};

exports.handler = async(event, context, callback) => {
    // Wait for connection
    producer.produce(
        TOPIC_NAME,   
        null,      
        Buffer.from(event),
        null,
        Date.now()
    );
    producer.flush();

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

module.exports.preStop = function(context, callback) {
    console.log('preStop hook start');
    if (producer != null) {
        producer.disconnect();
    }
    console.log('preStop hook finish');
    callback(null, "");
}