package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.StreamRequestHandler;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.PreStopHandler;

import java.util.HashMap;
import java.util.Map;
import java.sql.*;

// Notice, do not import com.mysql.jdbc.*
// or you will have problems!

public class App implements StreamRequestHandler, FunctionInitializer, PreStopHandler {
    // jdbc:mysql://<hostname>:<port>/<db>
    String JDBC_URL = null;
    String JDBC_USER = null;
    String JDBC_PASSWORD = null;

    private Connection connect = null;
    private Statement stmt = null;
    private ResultSet rs = null;

    @Override
    public void initialize(Context context) throws IOException {
        // jdbc:mysql://<hostname>:<port>/<db>
        JDBC_URL = "jdbc:mysql://"+System.getenv("MYSQL_ENDPOINT")+":"+System.getenv("MYSQL_PORT")+"/"+System.getenv("MYSQL_DBNAME");
        JDBC_USER = System.getenv("MYSQL_USER");
        JDBC_PASSWORD = System.getenv("MYSQL_PASSWORD");
        long start = System.currentTimeMillis();
        try {
            // The newInstance() call is a work around for some
            // broken Java implementations
            Class.forName("com.mysql.cj.jdbc.Driver").getDeclaredConstructor().newInstance();
            // Setup the connection with the DB
            connect = DriverManager.getConnection(JDBC_URL, JDBC_USER, JDBC_PASSWORD);
        } catch (Exception ex) {
            // handle any errors
            context.getLogger().error("SQLException: " + ex.getMessage());
            throw new RuntimeException(ex.getMessage());
        } finally {
            context.getLogger().info("database connection time cost: " + (System.currentTimeMillis() - start) + "ms");
        }
    }

    @Override
    public void preStop(Context context) throws IOException {
        context.getLogger().info("preStop start");
        if (connect != null) {
            try {
                connect.close();
            } catch (SQLException sqlEx) {
                context.getLogger().error("SQLException: " + sqlEx.getMessage());
                context.getLogger().error("SQLState: " + sqlEx.getSQLState());
                context.getLogger().error("VendorError: " + sqlEx.getErrorCode());
            }
            connect = null;
        }
        context.getLogger().info("preStop end");
    }

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        Map<String, Object> result = new HashMap<String, Object>();
        try {
            // Create a Statement instance that we can use for
            // 'normal' result sets assuming you have a
            // Connection 'connect' to a MySQL database already available
            stmt = connect.createStatement();

            // Insert a user item into the database
            int rowsEffected = stmt.executeUpdate("INSERT INTO users(name, age) VALUES('wanger', 5)");
            context.getLogger().info("Success - " + rowsEffected + " rows affected.");

            // Get a user item from database
            // Result set get the result of the SQL query
            rs = stmt.executeQuery("SELECT * FROM `users` ORDER BY `id` DESC LIMIT 1");

            // ResultSet is initially before the first data set
            while (rs.next()) {
                result.put("id", rs.getLong("id"));
                result.put("name", rs.getString("name"));
                result.put("age", rs.getShort("age"));
                break;
            }
        } catch (SQLException ex) {
            // handle any errors
            context.getLogger().error("SQLException: " + ex.getMessage());
            context.getLogger().error("SQLState: " + ex.getSQLState());
            context.getLogger().error("VendorError: " + ex.getErrorCode());
        } finally {
            // it is a good idea to release
            // resources in a finally{} block
            // in reverse-order of their creation
            // if they are no-longer needed
            if (rs != null) {
                try {
                    rs.close();
                } catch (SQLException sqlEx) { } // ignore
                rs = null;
            }

            if (stmt != null) {
                try {
                    stmt.close();
                } catch (SQLException sqlEx) { } // ignore

                stmt = null;
            }
        }
        context.getLogger().info("get user: " + result.toString());
        outputStream.write(result.toString().getBytes());
    }
}
