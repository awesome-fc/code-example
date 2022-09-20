<?php

use Aliyun\OTS\OTSClient as OTSClient;
use RingCentral\Psr7\Response;

$otsClient;

function initialize($context) {
    global $otsClient;
    $creds = $context['credentials'];
    $otsClient = new OTSClient([
        'EndPoint' => getenv("ENDPOINT"),
        'AccessKeyID' => $creds["accessKeyId"],
        'AccessKeySecret' => $creds["accessKeySecret"],
        "StsToken" => $creds["securityToken"],
        'InstanceName' => getenv("INSTANCE_NAME")
    ]);
}

function handler($request, $context) {
    global $otsClient;
    $queries = $request->getQueryParams();
    $res = $otsClient->getRow([
        "table_name" => getenv("TABLE_NAME"),
        "primary_key" => [
            ["region", $queries['region']],
            ["id", intval($queries['id'])]
        ],
        "max_versions" => 1
    ]);
    return new Response(
        200,
        array("Content-Type" => "application/json"),
        json_encode($res['attribute_columns'])
    );
}