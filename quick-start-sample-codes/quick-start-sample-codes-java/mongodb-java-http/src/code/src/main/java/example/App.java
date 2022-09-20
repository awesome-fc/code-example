package example;
import static com.mongodb.client.model.Filters.eq;

import java.io.IOException;
import java.io.OutputStream;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.aliyun.fc.runtime.Context;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.HttpRequestHandler;
import com.aliyun.fc.runtime.PreStopHandler;

import org.bson.Document;

public class App implements HttpRequestHandler, FunctionInitializer, PreStopHandler {
    // mongodb://<hostname>:<port>
    String MONGO_URL = null;
    String MONGO_DATABASE = null;

    private MongoClient mongoClient = null;

    @Override
    public void initialize(Context context) {
        // 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
        MONGO_URL = System.getenv("MONGO_URL");
        MONGO_DATABASE = System.getenv("MONGO_DATABASE");
        mongoClient = MongoClients.create(MONGO_URL);
    }

    @Override
    public void preStop(Context context) {
        mongoClient.close();
    }

    @Override
    public void handleRequest(HttpServletRequest request, HttpServletResponse response, Context context) throws IOException {
        MongoDatabase database = mongoClient.getDatabase(MONGO_DATABASE);
        MongoCollection<Document> collection = database.getCollection("users");
        
        Document doc = collection.find(eq("name", request.getParameter("name"))).first();

        context.getLogger().info("get user: " + doc.toString());
        response.setStatus(200);
        response.setHeader("Content-Type", "application/json");
        OutputStream out = response.getOutputStream();
        out.write(doc.toString().getBytes());
        out.flush();
        out.close();
    }
}
