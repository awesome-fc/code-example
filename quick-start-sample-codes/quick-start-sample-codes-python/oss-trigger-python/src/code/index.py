"""
本代码样例主要实现以下功能:
* 从 event 中解析出 OSS 事件触发相关信息
* 根据以上获取的信息，初始化 OSS bucket 客户端
* 从 OSS bucket 下载即将被处理的目标图片
* 改变目标图片尺寸
* 将处理过的图片上传到 OSS bucket 下的 processed 目录


This sample code is mainly doing the following things:
* Get OSS processing related information from event 
* Initiate OSS client with target bucket
* Download the target image to be processed from bucket
* Resize the image
* Upload the processed image copy into the same bucket's processed folder

"""

# -*- coding: utf-8 -*-
import oss2, json
from wand.image import Image

def handler(event, context):
    
    # 可以通过 context.credentials 获取密钥信息
    # Access keys can be fetched through context.credentials
    print("The content in context entity is: \n")
    print(context)
    creds = context.credentials

    # 设置权鉴，供 OSS sdk 使用
    # Setup auth, required by OSS sdk
    auth=oss2.StsAuth(
        creds.access_key_id,
        creds.access_key_secret,
        creds.security_token)

    print("The content in event entity is: \n")
    print(event)
    # Load event content
    oss_raw_data = json.loads(event)
    # Get oss event related parameters passed by oss trigger 
    oss_info_map = oss_raw_data['events'][0]['oss']
    # Get oss bucket name
    bucket_name = oss_info_map['bucket']['name']
    # Set oss service endpoint
    endpoint = 'oss-' +  oss_raw_data['events'][0]['region'] + '-internal.aliyuncs.com'
    # Initiate oss client
    bucket = oss2.Bucket(auth, endpoint, bucket_name)
    object_name = oss_info_map['object']['key']

    # Download original image from oss bucket
    remote_stream = bucket.get_object(object_name)
    if not remote_stream:
        return
    remote_stream = remote_stream.read()
    # Processed images will be saved to processed/
    processed_path = 'processed/' + object_name

    # Resize original image and upload the processed copy into the same bucket's processed folder
    with Image(blob=remote_stream)  as img:
        with img.clone() as i:
            i.resize(128, 128)
            new_blob = i.make_blob()
            bucket.put_object(processed_path, new_blob)