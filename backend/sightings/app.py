import os

from flask import Flask
from flask_alembic import Alembic

from .models import db

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = os.environ["DATABASE_URL"]

db.init_app(app)

alembic = Alembic(app)


@app.route("/")
def ping():
    return "OK"
