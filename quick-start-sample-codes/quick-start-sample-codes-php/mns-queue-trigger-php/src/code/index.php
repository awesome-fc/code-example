<?php

function handler($event, $context) {
  $logger = $GLOBALS['fcLogger'];
  $logger->info("Receive mns queue whole message:$event");
  $eventJson = json_decode($event,true);
  $logger->info(sprintf("MessageBody is: %s ,MessageID is: %s",$eventJson['data']['messageBody'],$eventJson['data']['messageId']));
  return sprintf("MessageBody is: %s ,MessageID is: %s",$eventJson['data']['messageBody'],$eventJson['data']['messageId']);
}
