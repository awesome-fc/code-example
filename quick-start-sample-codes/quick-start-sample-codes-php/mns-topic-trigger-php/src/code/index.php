<?php

function handler($event, $context) {
    $logger = $GLOBALS['fcLogger'];
    $eventJson = json_decode($event,true);
    if (isset($eventJson['Message'])) {
        $logger->info(sprintf("MessageBody is: %s ,MessageID is: %s",$eventJson['Message'],$eventJson['MessageId']));
        return sprintf("MessageBody is: %s ,MessageID is: %s",$eventJson['Message'],$eventJson['MessageId']); 
    }
    $logger->info("the event format is STREAM and mns topic message content is:$event");
        return "the event format is STREAM and mns topic message content is:$event";
    
}