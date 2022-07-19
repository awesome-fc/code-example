package com.example.demo.controller;
import static com.mongodb.client.model.Filters.eq;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import com.alicloud.openservices.tablestore.SyncClient;
import com.alicloud.openservices.tablestore.model.GetRowResponse;
import com.alicloud.openservices.tablestore.model.GetRowRequest;
import com.alicloud.openservices.tablestore.model.PrimaryKey;
import com.alicloud.openservices.tablestore.model.PrimaryKeyBuilder;
import com.alicloud.openservices.tablestore.model.PrimaryKeyValue;
import com.alicloud.openservices.tablestore.model.Row;
import com.alicloud.openservices.tablestore.model.SingleRowQueryCriteria;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;

import org.bson.Document;

@RestController
public class DemoController {
    @Value("${mongo.url}")
    String mongoUrl;

    @Value("${mongo.database}")
    String mongoDataBase;

    @Value("${tablestore.endpoint}")
    String tableStoreEndpoint;

    @Value("${tablestore.instanceName}")
    String tableStoreInstanceName;

    @Value("${tablestore.accessKey}")
    String tableStoreAccessKey;

    @Value("${tablestore.accessKeySecret}")
    String tableStoreAccessKeySecret;

    // MongoDB 的 client
    MongoClient mongoClient;
    // Tablestore 的 client
    SyncClient  client;

    
    @PostConstruct
    void init(){
        mongoClient = MongoClients.create(mongoUrl);
        client = new SyncClient(tableStoreEndpoint, tableStoreAccessKey, tableStoreAccessKeySecret, tableStoreInstanceName);
    }

    @GetMapping("/mongo")
    public String mongo(){
        MongoDatabase database = mongoClient.getDatabase(mongoDataBase);

        MongoCollection<Document> collection = database.getCollection("users");
        
        Document doc = collection.find(eq("name", "张三")).first();

        String res = doc.toString();
        return res;
    }

    @GetMapping("/tablestore")
    public String tablestore(){
        // 本示例所用表格存储的主键包含两个主键列：region 和 id
        PrimaryKey primaryKey = PrimaryKeyBuilder.createPrimaryKeyBuilder()
        .addPrimaryKeyColumn("region", PrimaryKeyValue.fromString("abc"))
        .addPrimaryKeyColumn("id", PrimaryKeyValue.fromLong(1))
        .build();

        // 读取一行数据，设置数据表名称。
        SingleRowQueryCriteria criteria = new SingleRowQueryCriteria("fc_test", primaryKey);
        // 设置读取最新版本
        criteria.setMaxVersions(1);
        GetRowResponse getRowResponse = client.getRow(new GetRowRequest(criteria));
        Row row = getRowResponse.getRow();
        String res = row.toString();
        return res;
    }

    @PreDestroy
    void PreDestroy(){
        mongoClient.close();
        client.shutdown();
    }

}