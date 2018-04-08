import csv
import decimal
import arrow

from .models import db, Report


def parse_row_data(row):
    if len(row) != 11:
        return None

    (
        occurred_at,
        city,
        state,
        country,
        shape,
        _,
        duration,
        description,
        reported_at,
        latitude,
        longitude,
    ) = row

    location = ",".join([item for item in [city, state, country] if item])

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
        return decimal.Decimal(value.replace("'", "").replace('"', ""))

    except decimal.InvalidOperation:
        return None


def parse_date(value):
    try:
        return arrow.get(value, "M/D/YYYY").date()

    except (arrow.parser.ParserError, ValueError):
        return None


def import_csv_data(filename):
    reader = csv.reader(open(filename))
    for counter, row in enumerate(reader):
        # skip first
        if counter == 0:
            continue

        report = parse_row_data(row)
        if report:
            db.session.add(report)
    db.session.commit()
