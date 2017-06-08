from django.test import TestCase

from .models import Report


class ReportTest(TestCase):

    def test_to_str(self):

        report = Report(location="Des Moines, Iowa")
        self.assertEquals(str(report), "Des Moines, Iowa")

    def test_search(self):

        report = Report.objects.create(
                location="Des Moines, Iowa",
                shape="Rectangle",
                duration="10 mins",
                description="...",
        )

        qs = Report.objects.search("Iowa")
        self.assertEqual(qs.count(), 1)
        self.assertEqual(qs.first(), report)
