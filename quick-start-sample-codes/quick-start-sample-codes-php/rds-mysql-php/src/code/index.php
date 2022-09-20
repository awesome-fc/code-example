<?php

$conn = null;

function initialize($context) {
    global $conn;
    $user = getenv("MYSQL_USER");
    $password = getenv("MYSQL_PASSWORD");
    $endpoint = getenv("MYSQL_ENDPOINT");
    $port = intval(getenv("MYSQL_PORT"));
    $dbname = getenv("MYSQL_DBNAME");

    $conn = new mysqli($endpoint, $user, $password, $dbname, $port);

    $logger = $GLOBALS['fcLogger'];
    $logger->info('initializing done');
}

function handler($event, $context) {
    global $conn;
    $logger = $GLOBALS['fcLogger'];
    // insert data
    $sql = "INSERT INTO USERS (NAME, AGE)
    VALUES ('John', 38)";
    if ($conn->query($sql) === TRUE) {
        $logger->info('insert succ');
    } else {
        $logger->info("Error: " . $conn->error);
    }
    // query data 
    $query = $conn->query("select * from `users`", MYSQLI_USE_RESULT);
    $result = $query->fetch_all();
    return $result;
}

function pre_stop($context) {
    global $conn;
    $conn->close();
    $logger = $GLOBALS['fcLogger'];
    $logger->info('pre_stop done');
}