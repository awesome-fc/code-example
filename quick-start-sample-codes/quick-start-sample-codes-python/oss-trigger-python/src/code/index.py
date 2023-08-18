"""
本代码样例主要实现以下功能:
* 从 event 中解析出 OSS 事件触发相关信息
* 根据以上获取的信息，初始化 OSS bucket 客户端
* 将源图片 resize 后持久化到OSS bucket 下指定的目标图片路径，从而实现图片备份


This sample code is mainly doing the following things:
* Get OSS processing related information from event 
* Initiate OSS client with target bucket
* Resize the source image and then store the processed image into the same bucket's copy folder to backup the image

"""

# -*- coding: utf-8 -*-
import oss2, json
import base64

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
        print(f'{object_name} does not exist in bucket {bucket_name}')
        return
    # Processed images will be saved to processed/
    processed_path = object_name.replace('source/', 'processed/')

    # 将图片缩放为固定宽高128 px。
    style = 'image/resize,m_fixed,w_128,h_128'
    # 指定处理后图片名称。如果图片不在Bucket根目录，需携带文件完整访问路径，例如exampledir/example.jpg。
    process = "{0}|sys/saveas,o_{1},b_{2}".format(style, 
        oss2.compat.to_string(base64.urlsafe_b64encode(oss2.compat.to_bytes(processed_path))),
        oss2.compat.to_string(base64.urlsafe_b64encode(oss2.compat.to_bytes(bucket_name))))
    result = bucket.process_object(object_name, process)
    print(result)