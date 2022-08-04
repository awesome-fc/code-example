package example;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MnsMessage {
    @JsonProperty(value = "TopicName")
    private String topicName;

    @JsonProperty(value = "TopicOwner")
    private String topicOwner;

    @JsonProperty(value = "Message")
    private String message;

    @JsonProperty(value = "MessageId")
    private String messageId;

    @JsonProperty(value = "MessageMD5")
    private String messageMD5;

    @JsonProperty(value = "Subscriber")
    private String subscriber;

    @JsonProperty(value = "SubscriptionName")
    private String subscriptionName;

    @JsonProperty(value = "PublishTime")
    private long publishTime;

    public String getTopicName() {
        return topicName;
    }

    public void setTopicName(String topicName) {
        this.topicName = topicName;
    }

    public String getTopicOwner() {
        return topicOwner;
    }

    public void setTopicOwner(String topicOwner) {
        this.topicOwner = topicOwner;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String getMessageId() {
        return messageId;
    }

    public void setMessageId(String messageId) {
        this.messageId = messageId;
    }

    public String getMessageMD5() {
        return messageMD5;
    }

    public void setMessageMD5(String messageMD5) {
        this.messageMD5 = messageMD5;
    }

    public String getSubscriber() {
        return subscriber;
    }

    public void setSubscriber(String subscriber) {
        this.subscriber = subscriber;
    }

    public String getSubscriptionName() {
        return subscriptionName;
    }

    public void setSubscriptionName(String subscriptionName) {
        this.subscriptionName = subscriptionName;
    }

    public long getPublishTime() {
        return publishTime;
    }

    public void setPublishTime(long publishTime) {
        this.publishTime = publishTime;
    }

    @java.lang.Override
    public java.lang.String toString() {
        return "MnsMessage{" +
                "topicName='" + topicName + '\'' +
                ", topicOwner='" + topicOwner + '\'' +
                ", message='" + message + '\'' +
                ", messageId='" + messageId + '\'' +
                ", messageMD5='" + messageMD5 + '\'' +
                ", subscriber='" + subscriber + '\'' +
                ", subscriptionName='" + subscriptionName + '\'' +
                ", publishTime=" + publishTime +
                '}';
    }
}