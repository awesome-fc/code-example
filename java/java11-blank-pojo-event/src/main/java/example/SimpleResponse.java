package example;

public class SimpleResponse {
    String message;

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public SimpleResponse() {}
    public SimpleResponse(String message) {
        this.message = message;
    }
}