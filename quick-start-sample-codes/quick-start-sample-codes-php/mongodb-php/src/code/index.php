<?php

$mongoManager;

function initialize($context) {
    global $mongoManager;
    $mongoManager = new MongoDB\Driver\Manager(getenv("MONGO_URL"));
}

function handler($event, $context){
    global $mongoManager;
    $dbName = getenv("MONGO_DATABASE");
    $filter = ['name' => "张三"];
    $options = [];
    $query = new MongoDB\Driver\Query($filter, $options);
    $cursor = $mongoManager->executeQuery($dbName . ".users", $query);
    $res = [];
    foreach ($cursor as $doc){
        array_push($res, $doc);
    }
    return $res;
}

