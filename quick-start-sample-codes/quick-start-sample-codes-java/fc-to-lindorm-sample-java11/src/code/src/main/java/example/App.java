/**
 * 本代码样例主要实现以下功能:
 * 1.初始化 lindorm 连接
 * 2.从环境变量中获取 lindorm 宽表引擎 table 名字
 * 3.创建此表并向表中插入一条数据
 * 4.读表中所有数据，查看是否符合预期
 *
 * This sample code is mainly doing the following things:
 * 1.Initialize lindorm connection
 * 2.Get the lindorm table name from environment variables
 * 3.Create this table and insert data into the table
 * 4.Read all the data in the table to see if it is as expected
 */

package example;

import java.io.*;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.Properties;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import com.aliyun.fc.runtime.*;

public class App implements StreamRequestHandler, FunctionInitializer {
    public static final String LINDORM_URL = System.getenv("DatabaseURL");
    private static final String USER_NAME = System.getenv("LindormUserName");
    private static final String PASSWORD = System.getenv("LindormPassword");
    private static final String SQL_TABLE_NAME = System.getenv("SQLTableName");

    private Connection pconn = null;
    private Statement statement = null;

    @Override
    public void initialize(Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();
        Properties properties = new Properties();
        properties.put("user", USER_NAME);
        properties.put("password", PASSWORD);
        properties.put("database", "default");

        try {
            pconn = DriverManager.getConnection(LINDORM_URL, properties);
            logger.info("initialize success");
            statement = pconn.createStatement();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Override
    public void handleRequest(
            InputStream input, OutputStream output, Context context) throws IOException {
        if (pconn == null) {
            System.out.println("initialize failed.");
            output.write(new String("failed").getBytes());
            return;
        }

        try {
            // 如果输入的 table 存在，则先删除，保证表数据清洁
            // If the input table exists, delete it first to ensure that the table data is clean
            String sql = "drop table if exists " + SQL_TABLE_NAME;
            statement.executeUpdate(sql);

            // 创建新表，新表共有两列：c1、c2
            // Create a new table, the new table has two columns: c1, c2
            sql = "create table if not exists " + SQL_TABLE_NAME
                    + "(c1 int, c2 int, primary key(c1))";
            statement.executeUpdate(sql);

            // 向表中插入数据，两列数据分别为 20 和 30
            // Insert data into the table, the two columns of data are 20 and 30
            StringBuilder sqlBuilder = new StringBuilder();
            sqlBuilder.append("upsert into " + SQL_TABLE_NAME + "(c1,c2) values(20,30)");
            PreparedStatement ps = pconn.prepareStatement(sqlBuilder.toString());
            ps.executeUpdate();

            // 查询表中的数据并输出
            // Query the data in the table and output
            sql = "select * from " + SQL_TABLE_NAME ;
            ps = pconn.prepareStatement(sql);
            ResultSet rs = ps.executeQuery();
            while (rs.next()) {
                Integer c1 = rs.getInt(1);
                Integer c2 = rs.getInt(2);
                System.out.println("row data: c1=" + c1 + ", " + "c2=" + c2);
            }
            deleteTable();
        } catch (Exception e) {
            e.printStackTrace();
            try {
                deleteTable();
            } catch (SQLException ex) {
                ex.printStackTrace();
            }
            output.write(new String("failed").getBytes());
        }
        output.write(new String("success").getBytes());
    }

    // 删除表
    // delete table
    public void deleteTable() throws SQLException {
        try {
            String sql = "OFFLINE TABLE " + SQL_TABLE_NAME;
            statement.executeUpdate(sql);

            sql = "drop table if exists " + SQL_TABLE_NAME;
            statement.executeUpdate(sql);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}