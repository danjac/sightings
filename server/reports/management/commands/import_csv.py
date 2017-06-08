import csv
import arrow

from django.core.management import BaseCommand

from reports.models import Report


class Command(BaseCommand):
    help = "Imports CSV reports data into database"

    def add_arguments(self, parser):
        parser.add_argument('filename')

    def parse_row_data(self, row):
        if len(row) != 9:
            return None

        (_,
                occurred_at,
                reported_at,
                location,
                shape,
                duration,
                description,
                latitude,
                longitude) = row

        return Report(
            occurred_at=self.parse_date(occurred_at),
            reported_at=self.parse_date(reported_at),
            location=location,
            shape=shape,
            duration=duration,
            description=description,
            latitude=self.parse_coord(latitude),
            longitude=self.parse_coord(longitude),
        )

    def parse_coord(self, value):
        return value.replace("'", "").replace('"', '')

    def parse_date(self, value):
        try:
            return arrow.get(value, 'YYYYMMDD').date()
        except arrow.ParserError:
            return None

    def handle(self, *args, **options):

        filename = options['filename']

        try:
            reader = csv.reader(open(filename), delimiter='\t')
        except IOError:
            self.stderr.write(
                self.style.ERROR('File %s could not be opened' % filename)
            )
            return

        reports = []
        for row in reader:
            report = self.parse_row_data(row)
            if report is not None:
                reports.append(report)
        Report.objects.bulk_create(reports)
