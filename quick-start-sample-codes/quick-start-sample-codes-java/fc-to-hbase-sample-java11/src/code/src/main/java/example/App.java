// 本代码样例主要实现以下功能:
// 初始化 Hbase client
// 从环境变量中获取 Hbase table 名字
// 创建表格并向表中插入一条数据
// 读表中插入行的数据
// 清理该行数据
// 删除表格
//
// This sample code is mainly doing the following things:
// Initialize Hbase client
// Get the Hbase table name from environment variables
// Create this table and insert data into the table
// Read the data of the row in the table
// Clear the data inserted before
// delete the table

package example;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.*;
import org.apache.hadoop.hbase.client.*;


import java.io.*;

import com.aliyun.fc.runtime.*;

public class App implements StreamRequestHandler, FunctionInitializer {
    private static final String ZK_ADDRESS = System.getenv("HBaseZKURL");
    private static final String TABLE_NAME = System.getenv("TableName");
    private static final String CF_DEFAULT = "cf";
    public static final byte[] QUALIFIER = "col1".getBytes();
    private static final byte[] ROWKEY = "rowkey1".getBytes();
    private Connection connection = null;

    @Override
    public void initialize(Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();
        Configuration config = HBaseConfiguration.create();
        config.set(HConstants.ZOOKEEPER_QUORUM, ZK_ADDRESS);

        try {
            connection = ConnectionFactory.createConnection(config);
            logger.info("initialize success");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Override
    public void handleRequest(
            InputStream input, OutputStream output, Context context) throws IOException {
            if (connection == null) {
                System.out.println("initialize failed.");
                output.write(new String("failed").getBytes());
                return;
            }

            try {
                HTableDescriptor tableDescriptor = new HTableDescriptor(TableName.valueOf(TABLE_NAME));
                tableDescriptor.addFamily(new HColumnDescriptor(CF_DEFAULT));
                System.out.println("Creating table. ");
                Admin admin = connection.getAdmin();
                try {
                    admin.getTableDescriptor(tableDescriptor.getTableName());
                } catch (TableNotFoundException e) {
                    admin.createTable(tableDescriptor);
                }

                System.out.println(" Done.");
                Table table = connection.getTable(TableName.valueOf(TABLE_NAME));
                try {
                    admin.enableTable(table.getName());
                } catch (TableNotDisabledException e){
                    System.out.println("table enabled already");
                }

                try {
                    Put put = new Put(ROWKEY);
                    put.addColumn(CF_DEFAULT.getBytes(), QUALIFIER, "this is value".getBytes());
                    table.put(put);
                    Get get = new Get(ROWKEY);
                    Result r = table.get(get);
                    byte[] b = r.getValue(CF_DEFAULT.getBytes(), QUALIFIER);  // returns current version of value
                    System.out.println(new String(b));
                    Delete delete = new Delete(ROWKEY);
                    table.delete(delete);
                } finally {
                    if (table != null) table.close();
                    admin.disableTable(table.getName());
                    admin.deleteTable(table.getName());
                }
                admin.close();
            } catch (Exception e) {
                e.printStackTrace();
                output.write(new String("failed").getBytes());
                return;
            }
            output.write(new String("success").getBytes());
    }
}