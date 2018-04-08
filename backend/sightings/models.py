from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()


class Report(db.Model):

    __tablename__ = "reports"

    id = db.Column(db.BigInteger, primary_key=True)

    location = db.Column(db.Text, nullable=True)
    shape = db.Column(db.Text, nullable=True)
    duration = db.Column(db.Text, nullable=True)
    description = db.Column(db.Text, nullable=True)

    latitude = db.Column(db.Numeric, nullable=True)
    longitude = db.Column(db.Numeric, nullable=True)

    reported_at = db.Column(db.DateTime, nullable=True)
    occurred_at = db.Column(db.DateTime, nullable=True)
