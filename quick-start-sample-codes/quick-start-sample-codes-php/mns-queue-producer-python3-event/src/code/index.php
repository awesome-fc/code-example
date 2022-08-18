<?php

use AliyunMNS\Client;
use AliyunMNS\Requests\SendMessageRequest;
use AliyunMNS\Exception\MnsException;

$client = null;
$queue = null;

function initialize($context) {
    global $client,$queue;
    $creds = $context["credentials"];
    $accessKeyId = $creds["accessKeyId"];
    $accessKeySecret = $creds["accessKeySecret"];
    $endPoint=getenv("MNS_ENDPOINT");
    $queueName=getenv("MNS_QUEUE_NAME");
    // 1.首先初始化一个client。
    $client = new Client($endPoint,$accessKeyId,$accessKeySecret);
    // 2.获取Queue的实例。
    // PHP SDK默认会对发送的消息做Base64 Encode，对接收到的消息做Base64 Decode。
    // 如果不希望SDK做这样的Base64操作，可以在getQueueRef的时候，传入参数$base64=FALSE。即$queue = $client->getQueueRef($queueName, FALSE);
    $queue = $client->getQueueRef($queueName,FALSE);
}

function handler($event, $context) {
    global $client,$queue;
    $messageBody = "I am a test message";
    // 3.生成一个SendMessageRequest对象。
    // SendMessageRequest对象本身也包含了DelaySeconds和Priority属性可以设置。
    // 对于Message的属性，请参见QueueMessage。
    $bodyMD5 = md5(base64_encode($messageBody));
    $request = new SendMessageRequest($messageBody);
    $logger = $GLOBALS['fcLogger'];
    try
    {
        $res = $queue->sendMessage($request);
        // 4.消息发送成功。
        $logger->info( "MessageSent!");
        return "Send Message Succeed. MessageBody:$messageBody";
    }
    catch (MnsException $e)
    {
        // 5.可能因为网络错误，或MessageBody过大等原因造成发送消息失败，这里CatchException并做对应的处理。
        $logger->info("SendMessage Failed: " . $e);
    }   
}