package com.example.demo.controller;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;

import org.codehaus.jackson.map.ObjectMapper;
import org.codehaus.jackson.type.TypeReference;

import java.io.IOException;
import java.util.Map;
import java.util.List;

@SpringBootApplication
@RestController
public class DemoController {

    @PostMapping("/invoke")
    public ResponseEntity<String> invoke(@RequestHeader Map<String, String> headers, @RequestBody String events) throws IOException {
        ObjectMapper objectMapper = new ObjectMapper();
        System.out.println(events);
        System.out.println();
        List<Event> eventsObj = objectMapper.readValue(events, new TypeReference<List<Event>>(){});
        for (int i = 0; i < eventsObj.size(); i++) {
          Event event = eventsObj.get(i);
          RabbitMQData rabbitmqData = event.getData();
          System.out.println(rabbitmqData.getBody());
        }
        return new ResponseEntity<>("rabbitmq demo", HttpStatus.OK);
    }

}

class Event {
    RabbitMQData data;
    String id;
    String source;
    String specversion;
    String type;
    String datacontenttype;
    String time;
    String subject;
    String aliyunaccountid;
    
    public Event() {}

    public Event(RabbitMQData data,
                 String id,
                 String source,
                 String specversion,
                 String type,
                 String datacontenttype,
                 String time,
                 String subject,
                 String aliyunaccountid) {
        this.data = data;
        this.id = id;
        this.source = source;
        this.specversion = specversion;
        this.type = type;
        this. datacontenttype = datacontenttype;
        this.time = time;
        this.subject = subject;
        this.aliyunaccountid = aliyunaccountid;
    }

    public RabbitMQData getData() {
        return data;
    }

    public void setData(RabbitMQData data) {
        this.data = data;
    }

    public String getAliyunaccountid() {
        return aliyunaccountid;
    }

    public void setAliyunaccountid(String aliyunaccountid) {
        this.aliyunaccountid = aliyunaccountid;
    }

    public String getDatacontenttype() {
        return datacontenttype;
    }

    public void setDatacontenttype(String datacontenttype) {
        this.datacontenttype = datacontenttype;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

    public String getSpecversion() {
        return specversion;
    }
    public String getType() {
        return type;
    }

    public String getTime() {
        return time;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getSubject() {
        return subject;
    }

    public void setSpecversion(String specversion) {
        this.specversion = specversion;
    }

    public void setTime(String time) {
        this.time = time;
    }

    public void setSubject(String subject) {
        this.subject = subject;
    }
}

class RabbitMQData {
    String body;
    RabbitEnvelope envelope;
    RabbitProps props;

    public RabbitMQData() {}

    public RabbitMQData(String body,
                        RabbitEnvelope envelope,
                        RabbitProps props) {
        this.body = body;
        this.envelope = envelope;
        this.props = props;
    }

    public String getBody() {
        return body;
    }

    public void setBody(String body) {
        this.body = body;
    }

    public RabbitEnvelope getEnvelope() {
        return envelope;
    }

    public void setEnvelope(RabbitEnvelope envelope) {
        this.envelope = envelope;
    }

    public RabbitProps getProps() {
        return props;
    }

    public void setProps(RabbitProps props) {
        this.props = props;
    }
}

class RabbitEnvelope {
    int deliveryTag;
    boolean redeliver;
    String exchange;
    String routingKey;

    public RabbitEnvelope() {}
    public RabbitEnvelope(int deliveryTag,
                          boolean redeliver,
                          String exchange,
                          String routingKey) {
        this.deliveryTag = deliveryTag;
        this.redeliver = redeliver;
        this.exchange = exchange;
        this.routingKey = routingKey;
    }

    public int getDeliveryTag() {
        return deliveryTag;
    }

    public void setDeliveryTag(int deliveryTag) {
        this.deliveryTag = deliveryTag;
    }

    public boolean isRedeliver() {
        return redeliver;
    }

    public void setRedeliver(boolean redeliver) {
        this.redeliver = redeliver;
    }

    public String getExchange() {
        return exchange;
    }

    public void setExchange(String exchange) {
        this.exchange = exchange;
    }

    public String getRoutingKey() {
        return routingKey;
    }

    public void setRoutingKey(String routingKey) {
        this.routingKey = routingKey;
    }
}

class RabbitProps {
    String messageId;

    public RabbitProps() {}

    public RabbitProps(String messageId) {
        this.messageId = messageId;
    }

    public String getMessageId() {
        return messageId;
    }

    public void setMessageId(String messageId) {
        this.messageId = messageId;
    }
}