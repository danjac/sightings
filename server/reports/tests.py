from django.test import TestCase

from .models import Report


class ReportTest(TestCase):

    def test_to_str(self):

        report = Report(location="Des Moines, Iowa")
        self.assertEquals(str(report), "Des Moines, Iowa")
