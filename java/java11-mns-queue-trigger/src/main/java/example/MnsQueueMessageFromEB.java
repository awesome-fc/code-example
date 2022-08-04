package example;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
// MNS queue message from EventBridge
public class MnsQueueMessageFromEB {
    private String id;
    private String source;
    private String specversion;
    private String type;
    private String datacontenttype;
    private String subject;
    private String time;
    private String aliyunaccountid;
    private String aliyunpublishtime;
    private String aliyunoriginalaccountid;
    private String aliyuneventbusname;
    private String aliyunregionid;
    private String aliyunpublishaddr;

    private MnsQueueData data;

    @JsonIgnoreProperties(ignoreUnknown = true)
    public class MnsQueueData {
        private String requestId;
        private String messageId;
        private String messageBody;

        public String getRequestId() {
            return requestId;
        }

        public void setRequestId(String requestId) {
            this.requestId = requestId;
        }

        public String getMessageId() {
            return messageId;
        }

        public void setMessageId(String messageId) {
            this.messageId = messageId;
        }

        public String getMessageBody() {
            return messageBody;
        }

        public void setMessageBody(String messageBody) {
            this.messageBody = messageBody;
        }

        @java.lang.Override
        public java.lang.String toString() {
            return "{" +
                    "requestId='" + requestId + '\'' +
                    ", messageId='" + messageId + '\'' +
                    ", messageBody='" + messageBody + '\'' +
                    '}';
        }
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

    public void setSpecversion(String specversion) {
        this.specversion = specversion;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getDatacontenttype() {
        return datacontenttype;
    }

    public void setDatacontenttype(String datacontenttype) {
        this.datacontenttype = datacontenttype;
    }

    public String getSubject() {
        return subject;
    }

    public void setSubject(String subject) {
        this.subject = subject;
    }

    public String getTime() {
        return time;
    }

    public void setTime(String time) {
        this.time = time;
    }

    public String getAliyunaccountid() {
        return aliyunaccountid;
    }

    public void setAliyunaccountid(String aliyunaccountid) {
        this.aliyunaccountid = aliyunaccountid;
    }

    public String getAliyunpublishtime() {
        return aliyunpublishtime;
    }

    public void setAliyunpublishtime(String aliyunpublishtime) {
        this.aliyunpublishtime = aliyunpublishtime;
    }

    public String getAliyunoriginalaccountid() {
        return aliyunoriginalaccountid;
    }

    public void setAliyunoriginalaccountid(String aliyunoriginalaccountid) {
        this.aliyunoriginalaccountid = aliyunoriginalaccountid;
    }

    public String getAliyuneventbusname() {
        return aliyuneventbusname;
    }

    public void setAliyuneventbusname(String aliyuneventbusname) {
        this.aliyuneventbusname = aliyuneventbusname;
    }

    public String getAliyunregionid() {
        return aliyunregionid;
    }

    public void setAliyunregionid(String aliyunregionid) {
        this.aliyunregionid = aliyunregionid;
    }

    public String getAliyunpublishaddr() {
        return aliyunpublishaddr;
    }

    public void setAliyunpublishaddr(String aliyunpublishaddr) {
        this.aliyunpublishaddr = aliyunpublishaddr;
    }

    public example.MnsQueueMessageFromEB.MnsQueueData getData() {
        return data;
    }

    public void setData(example.MnsQueueMessageFromEB.MnsQueueData data) {
        this.data = data;
    }

    @java.lang.Override
    public java.lang.String toString() {
        return "MnsQueueMessageFromEB{" +
                "id='" + id + '\'' +
                ", source='" + source + '\'' +
                ", specversion='" + specversion + '\'' +
                ", type='" + type + '\'' +
                ", datacontenttype='" + datacontenttype + '\'' +
                ", subject='" + subject + '\'' +
                ", time='" + time + '\'' +
                ", aliyunaccountid='" + aliyunaccountid + '\'' +
                ", aliyunpublishtime='" + aliyunpublishtime + '\'' +
                ", aliyunoriginalaccountid='" + aliyunoriginalaccountid + '\'' +
                ", aliyuneventbusname='" + aliyuneventbusname + '\'' +
                ", aliyunregionid='" + aliyunregionid + '\'' +
                ", aliyunpublishaddr='" + aliyunpublishaddr + '\'' +
                ", data=" + data +
                '}';
    }
}