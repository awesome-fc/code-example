from flask import Flask
from flask import request

app = Flask(__name__)


@app.route('/')
def index():
    return '''<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>Serverless Devs - Powered By Serverless Devs</title>
    <link href="https://example-static.oss-cn-beijing.aliyuncs.com/web-framework/style.css" rel="stylesheet" type="text/css"/>
</head>
<body>
<div class="website">
    <div class="ri-t">
        <h1>Devsapp</h1>
        <h2>这是一个 Flask 项目</h2>
        <span>自豪的通过Serverless Devs进行部署</span>
        <br/>
        <p>您也可以快速体验： <br/>
            • 下载Serverless Devs工具：npm install @serverless-devs/s<br/>
            • 初始化项目：s init start-flask<br/>

            • 项目部署：s deploy<br/>
            <br/>
            Serverless Devs 钉钉交流群：33947367
        </p>
    </div>
</div>
</body>
</html>
'''

@app.route('/initialize', methods=['POST'])
def init_invoke():
    rid = request.headers.get("x-fc-request-id")
    print("FC Initialize Start RequestId: " + rid)
    # do your things
    print("FC Initialize End RequestId: " + rid)
    return "OK"

@app.route('/pre-freeze', methods=['GET'])
def pre_freeze_invoke():
    rid = request.headers.get("x-fc-request-id")
    print("FC PreFreeze Start RequestId: " + rid)
    # do your things
    print("FC PreFreeze End RequestId: " + rid)
    return "OK"

@app.route('/pre-stop', methods=['GET'])
def pre_stop_invoke():
    rid = request.headers.get("x-fc-request-id")
    print("FC PreStop Start RequestId: " + rid)
    # do your things
    print("FC PreStop End RequestId: " + rid)
    return "OK"
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=9000)
