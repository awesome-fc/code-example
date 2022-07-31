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
        //断线自动重连
        producer.connect();
    })

    producer.on('event.log', function(event) {
        console.log("event.log", event);
    });

    producer.on("error", function(error) {
        console.log("error:" + error);

    });

    producer.on('delivery-report', function(err, report) {
        console.log("delivery-report: producer ok");
    });
    // Any errors we encounter, including connection errors
    producer.on('event.error', function(err) {
        console.error('event.error:' + err);
    })
    // Poll for events every 10 ms
    producer.setPollInterval(10);

    callback(null, "");
};

function wait(ms) {
    return new Promise(resolve =>setTimeout(() =>resolve(), ms));
};

exports.handler = async(event, context, callback) => {
    // Connect to the broker manually
    producer.connect();
    // 等待connect成功
    await wait(5000);
    producer.produce(
        topic_name,   
        null,      
        Buffer.from(event),      
        null,   
        Date.now()
    );
    producer.flush();
    // 等待发送成功
    await wait(5000);
    
    callback(null, "Finish sending the message:" + event);
}