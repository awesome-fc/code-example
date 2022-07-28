package com.example.demo.controller;
import static com.mongodb.client.model.Filters.eq;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

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

    // MongoDB 的 client
    MongoClient mongoClient;

    
    @PostConstruct
    void init(){
        mongoClient = MongoClients.create(mongoUrl);
    }

    @GetMapping("/mongo")
    public String mongo(){
        MongoDatabase database = mongoClient.getDatabase(mongoDataBase);

        MongoCollection<Document> collection = database.getCollection("users");
        
        Document doc = collection.find(eq("name", "张三")).first();

        return doc.toString();
    }

    @PreDestroy
    void PreDestroy(){
        mongoClient.close();
    }

}