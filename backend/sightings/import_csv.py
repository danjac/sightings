import csv
import decimal
import arrow

from .models import db, Report


def parse_row_data(row):
    if len(row) != 9:
        return None

        (_,
         location,
         shape,
         duration,
         description,
         latitude,
         longitude,
         reported_at,
         occurred_at,
         ) = row

        return Report(
            location=location,
            shape=shape,
            duration=duration,
            description=description,
            latitude=parse_coord(latitude),
            longitude=parse_coord(longitude),
            occurred_at=parse_date(occurred_at),
            reported_at=parse_date(reported_at),
        )


def parse_coord(value):
    try:
        return decimal.Decimal(value.replace("'", "").replace('"', ''))
    except decimal.InvalidOperation:
        return None


def parse_date(value):
    try:
        return arrow.get(value, 'YYYY-MM-DD').date()
    except (arrow.parser.ParserError, ValueError):
        return None


def import_csv_data(filename):
    reader = csv.reader(open(filename))
    for row in reader:
        report = parse_row_data(row)
        if report:
            db.session.add(report)
    db.session.commit()
