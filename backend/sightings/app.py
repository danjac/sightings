import os
import click

from flask import Flask
from flask_alembic import Alembic

from .models import db
from .import_csv import import_csv_data

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = os.environ["DATABASE_URL"]

db.init_app(app)

alembic = Alembic(app)


@app.route("/")
def ping():
    return "OK"


@app.cli.command()
@click.argument('filename')
def import_csv(filename):
    import_csv_data(filename)
