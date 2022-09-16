package example;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.StreamRequestHandler;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.cbor.CBORFactory;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

// 本代码样例主要实现以下功能:
// * 打印 event 信息


// This sample code is mainly doing the following things:
// * print event


// event 示例见文档：https://help.aliyun.com/document_detail/169672.html，event结构如下所示：
// {
//     "Version": "Sync-v1",
//     "Records": [
//         {
//             "Type": "PutRow",
//             "Info": {
//                 "Timestamp": 1506416585740836
//             },
//             "PrimaryKey": [
//                 {
//                     "ColumnName": "pk_0",
//                     "Value": 1506416585881590900
//                 },
//                 {
//                     "ColumnName": "pk_1",
//                     "Value": "2017-09-26 17:03:05.8815909 +0800 CST"
//                 },
//                 {
//                     "ColumnName": "pk_2",
//                     "Value": 1506416585741000
//                 }
//             ],
//             "Columns": [
//                 {
//                     "Type": "Put",
//                     "ColumnName": "attr_0",
//                     "Value": "hello_table_store",
//                     "Timestamp": 1506416585741
//                 },
//                 {
//                     "Type": "Put",
//                     "ColumnName": "attr_1",
//                     "Value": 1506416585881590900,
//                     "Timestamp": 1506416585741
//                 }
//             ]
//         }
//     ]
// }

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(InputStream input, OutputStream output, Context context) throws IOException {
        CBORFactory f = new CBORFactory();
        ObjectMapper mapper = new ObjectMapper(f);
        mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);

        JsonNode node = mapper.readTree(input);
        System.out.println("Version: " + node.get("Version"));

        JsonNode records = node.get("Records");

        if (records.isArray()) {
            for (JsonNode record : records) {
                System.out.println("Type: " + record.get("Type"));
                JsonNode Info = record.get("Info");
                System.out.println("Timestamp: " + Info.get("Timestamp"));

                JsonNode PrimaryKey = record.get("PrimaryKey");
                if (PrimaryKey.isArray()){
                    for (JsonNode objNode : PrimaryKey){
                        System.out.println(objNode.get("ColumnName") + ": " + objNode.get("Value"));
                    }
                }

                JsonNode Columns = record.get("Columns");
                if (Columns.isArray()){
                    for (JsonNode objNode : Columns){
                        System.out.println(objNode.get("Type") + "  " + objNode.get("ColumnName") + ": " + objNode.get("Value") + "  >>" + objNode.get("Timestamp"));
                    }
                }
            }
        }
        output.write(new String("ok").getBytes());
    }
}