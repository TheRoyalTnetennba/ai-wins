from flask import Flask
app = Flask(__name__)

dictionary = Sowpods()

@app.route("/")
def hello():
    return "Hello World!"

@app.route('/user/<endpoint>')
def user(endpoint):
    """add user verification then switch statement for various resources"""