<?php

use Aliyun\OTS\OTSClient as OTSClient;

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

function handler($event, $context) {
    global $otsClient;
    $res = $otsClient->getRow([
        "table_name" => getenv("TABLE_NAME"),
        "primary_key" => [
            ["region", "abc"],
            ["id", 1]
        ],
        "max_versions" => 1
    ]);
    return $res['attribute_columns'];
}

