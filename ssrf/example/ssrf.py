# Usage: python3 ssrf.py
from flask import Flask, abort, request
import json
import re
import subprocess
import requests

app = Flask(__name__)

@app.route("/")
def hello():
    return "SSRF Example!"

@app.route("/ssrf", methods=['GET'])
def ssrf1():
    data = request.values
    res = requests.get(data.get('url'))
    return res.content


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
