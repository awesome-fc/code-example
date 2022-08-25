<?php
use RingCentral\Psr7\Response;

$mongoManager;

function initialize($context) {
    global $mongoManager;
    $mongoManager = new MongoDB\Driver\Manager(getenv("MONGO_URL"));
}

function handler($request, $context){
    global $mongoManager;
    $queries = $request->getQueryParams();
    $dbName = getenv("MONGO_DATABASE");
    $filter = ['name' => $queries['name']];
    $options = [];
    $query = new MongoDB\Driver\Query($filter, $options);
    $cursor = $mongoManager->executeQuery($dbName . ".users", $query);
    $res = [];
    foreach ($cursor as $doc){
        array_push($res, $doc);
    }
    return new Response(
        200,
        array("Content-Type" => "application/json"),
        json_encode($res)
    );
}

