package example;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.util.Properties;
import java.util.concurrent.Future;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;

import com.aliyun.fc.runtime.StreamRequestHandler;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;


public class App implements StreamRequestHandler, FunctionInitializer {
    String BOOTSTRAP_SERVERS = null;
    String TOPIC_NAME = null;
    KafkaProducer<String, String> producer;

    @Override
    public void initialize(Context context) {
        // Get the environment variables
        BOOTSTRAP_SERVERS = System.getenv("BOOTSTRAP_SERVERS");
        TOPIC_NAME = System.getenv("TOPIC_NAME");

        Properties props = new Properties();

        // Set the access point
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, BOOTSTRAP_SERVERS);

        // Set Kafka serialize type
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");

        producer = new KafkaProducer<>(props);
    }

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        // Read the inputStream
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String value = result.toString(StandardCharsets.UTF_8.name());

        // Create the message to be sent
        ProducerRecord<String, String> kafkaMessage =  new ProducerRecord<>(TOPIC_NAME, value);
        
        // Produce messages to topic (asynchronously)
        Future<RecordMetadata> metadataFuture = producer.send(kafkaMessage);

        // Flush the internel queue, wait for message deliveries before return
        producer.flush();

        try {
            RecordMetadata recordMetadata = metadataFuture.get();
            context.getLogger().info("Produce ok: " + recordMetadata.toString() + "\n Payload: " + value);
            outputStream.write(("Produce ok: " + recordMetadata.toString() + "\n Payload: " + value).getBytes());
        } catch(Exception e) {
             context.getLogger().error("Send message to kafka fail: " + e.toString());
             outputStream.write(("Produce fail: " + e.toString()).getBytes());
        }  
    }
}
