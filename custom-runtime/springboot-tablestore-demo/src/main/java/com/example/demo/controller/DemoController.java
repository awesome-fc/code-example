package com.example.demo.controller;

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

@RestController
public class DemoController {
    @Value("${tablestore.endpoint}")
    String tableStoreEndpoint;

    @Value("${tablestore.instanceName}")
    String tableStoreInstanceName;

    @Value("${tablestore.accessKey}")
    String tableStoreAccessKey;

    @Value("${tablestore.accessKeySecret}")
    String tableStoreAccessKeySecret;

    // Tablestore 的 client
    SyncClient  client;

    @PostConstruct
    void init(){
        client = new SyncClient(tableStoreEndpoint, tableStoreAccessKey, tableStoreAccessKeySecret, tableStoreInstanceName);
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
        return row.toString();
    }

    @PreDestroy
    void PreDestroy(){
        client.shutdown();
    }

}