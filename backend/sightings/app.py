import os
import click

from flask import Flask
from flask_alembic import Alembic
from flask_cors import CORS

from .models import db, Report
from .schemas import ma, ReportSchema, ReportPaginationSchema
from .import_csv import import_csv_data

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = os.environ["DATABASE_URL"]

CORS(app)

# DB initialization

db.init_app(app)
alembic = Alembic(app)

# Schema initialization

ma.init_app(app)

report_schema = ReportSchema()
report_pagination_schema = ReportPaginationSchema()


@app.route("/")
def ping():
    return "OK"


@app.route("/reports/")
def report_list():
    reports = Report.query.order_by(Report.occurred_at.desc()).paginate()
    return report_pagination_schema.jsonify(reports)


@app.route("/reports/<int:id>/")
def report_detail(id):
    report = Report.query.get_or_404(id)
    return report_schema.jsonify(report)


@app.cli.command()
@click.argument("filename")
def import_csv(filename):
    click.echo("Loading data from %s" % filename)
    import_csv_data(filename)


@app.shell_context_processor
def make_shell_context():
    return {
        'db': db,
        'Report': Report,
    }
