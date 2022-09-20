<?php

use AliyunMNS\Client;
use AliyunMNS\Requests\PublishMessageRequest;
use AliyunMNS\Exception\MnsException;

$client = null;
$topic = null;

function initialize($context) {
  global $client,$topic;
  $creds = $context["credentials"];
  $accessKeyId = $creds["accessKeyId"];
  $accessKeySecret = $creds["accessKeySecret"];
  $endPoint=getenv("MNS_ENDPOINT");
  $topicName=getenv("MNS_TOPIC_NAME");
  // 1.首先初始化一个client。
  $client = new Client($endPoint,$accessKeyId,$accessKeySecret);
  // 2.获取 Topic 的实例。
  $topic = $client->getTopicRef($topicName);
}

function handler($event, $context) {
  global $client,$topic;
  $messageBody = "I am a test message";
  // 3.生成PublishMessageRequest。如果是推送到邮箱，还需要设置MessageAttributes，可以参照Tests/TopicTest.php里面的testPublishMailMessage。
  $request = new PublishMessageRequest($messageBody);
  $logger = $GLOBALS['fcLogger'];
  try
  {
      $res = $topic->publishMessage($request);
      // 4.消息发送成功。
      $logger->info( "MessageSent!");
      $logger->info( "Publish Message Succeed. MessageBody:$messageBody");
      return "Publish Message Succeed. MessageBody:$messageBody";
  }
  catch (MnsException $e)
  {
          // 5.可能因为网络错误，或MessageBody过大等原因造成发送消息失败，这里CatchException并做对应的处理。
          $logger->info("SendMessage Failed: " . $e);
  }   
}