package example;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.util.Properties;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;

import com.aliyun.fc.runtime.StreamRequestHandler;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;


public class App implements StreamRequestHandler, FunctionInitializer {
    String BOOTSTRAP_SERVERS = null;
    String TOPIC_NAME = null;
    KafkaProducer<String, String> producer;

    @Override
     public void initialize(Context context) {
        // 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
        BOOTSTRAP_SERVERS = System.getenv("bootstrap_servers");
        TOPIC_NAME = System.getenv("topic_name");

        Properties props = new Properties();
        //设置接入点，请通过控制台获取对应Topic的接入点
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, BOOTSTRAP_SERVERS);

        //Kafka消息的序列化方式
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");

        producer = new KafkaProducer<>(props);
     }

    @Override
     public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {

        // 读出要发送的消息
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String value = result.toString(StandardCharsets.UTF_8.name());

        ProducerRecord<String, String> kafkaMessage =  new ProducerRecord<>(TOPIC_NAME, value);
        producer.send(kafkaMessage);

        producer.flush();

        outputStream.write(("Produce ok: " + value).getBytes());
     }
        
}
