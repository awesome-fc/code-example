from flask import Flask
from flask import request
import logging
import os

REQUEST_ID_HEADER = 'x-fc-request-id'

app = Flask(__name__)

format_str = '[%(asctime)s] %(levelname)s in %(module)s: %(message)s'
logging.basicConfig(filename='/tmp/log/fc-flask.log', filemode='w', 
    format=format_str, encoding='utf-8', level=logging.DEBUG)

@app.route("/invoke", methods = ["POST"])
def hello_world():
    rid = request.headers.get(REQUEST_ID_HEADER)
    logger = logging.getLogger()

    print("FC Invoke Start RequestId: " + rid)
    logger.info("FC Invoke Start RequestId: " + rid)

    data = request.stream.read()
    print(str(data))
    logger.info("receive event: {}".format(str(data)))
    
    print("FC Invoke End RequestId: " + rid)
    logger.info("FC Invoke Start RequestId: " + rid)
    return "Hello, World!"

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=9000)