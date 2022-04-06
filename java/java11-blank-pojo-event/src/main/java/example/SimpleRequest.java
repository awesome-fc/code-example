package example;

public class SimpleRequest {
    String firstName;
    String lastName;

    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public SimpleRequest() {}
    public SimpleRequest(String firstName, String lastName) {
        this.firstName = firstName;
        this.lastName = lastName;
    }
}