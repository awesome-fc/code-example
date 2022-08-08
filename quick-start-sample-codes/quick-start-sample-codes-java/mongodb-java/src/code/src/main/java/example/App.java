package example;
import static com.mongodb.client.model.Filters.eq;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.StreamRequestHandler;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.PreStopHandler;

import org.bson.Document;

public class App implements StreamRequestHandler, FunctionInitializer, PreStopHandler {
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
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        MongoDatabase database = mongoClient.getDatabase(MONGO_DATABASE);
        MongoCollection<Document> collection = database.getCollection("users");
        
        Document doc = collection.find(eq("name", "张三")).first();

        context.getLogger().info("get user: " + doc.toString());
        outputStream.write(doc.toString().getBytes());
    }
}
