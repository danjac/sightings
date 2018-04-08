from marshmallow import fields
from flask_marshmallow import Marshmallow
from flask_marshmallow.fields import URLFor

from .models import Report

ma = Marshmallow()


class ReportSchema(ma.ModelSchema):

    latitude = fields.Decimal(as_string=True)
    longitude = fields.Decimal(as_string=True)

    url = URLFor("report_detail", id="<id>")

    class Meta:
        model = Report
        strict = True


class ReportPaginationSchema(ma.Schema):

    has_next = fields.Boolean()
    has_prev = fields.Boolean()

    page = fields.Integer()

    pages = fields.Integer()
    per_page = fields.Integer()
    total = fields.Integer()

    next_num = fields.Integer()
    prev_num = fields.Integer()

    items = fields.Nested(ReportSchema, many=True)
