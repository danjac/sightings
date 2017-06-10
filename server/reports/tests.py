import factory

from django.test import TestCase
from rest_framework.test import APIRequestFactory

from .models import Report
from .views import ReportViewSet


class ReportFactory(factory.django.DjangoModelFactory):

    class Meta:
        model = Report

    location="Des Moines, Iowa"
    shape="Rectangle"
    duration="10 mins"
    description="..."


class ReportTest(TestCase):

    def test_to_str(self):

        report = Report(location="Des Moines, Iowa")
        self.assertEquals(str(report), "Des Moines, Iowa")

    def test_search(self):

        report = ReportFactory()

        qs = Report.objects.search("Iowa")
        self.assertEqual(qs.count(), 1)
        self.assertEqual(qs.first(), report)


class APITests(TestCase):

    def setUp(self):
        self.factory = APIRequestFactory()

    def test_fetch_all(self):

        ReportFactory()

        request = self.factory.get("/api/reports/")
        view = ReportViewSet.as_view({'get': 'list'})
        response = view(request)
        self.assertContains(response, 'Iowa')

    def test_search(self):

        ReportFactory()
        ReportFactory(location='Area 51, NM')

        request = self.factory.get("/api/reports/", {"s": "Area 51"})
        view = ReportViewSet.as_view({'get': 'list'})
        response = view(request)
        self.assertContains(response, 'Area 51')
        self.assertNotContains(response, 'Iowa')

    def test_fetch_one(self):

        report = ReportFactory()

        request = self.factory.get("/api/reports/%s/" % report.id)
        view = ReportViewSet.as_view({'get': 'retrieve'})
        response = view(request, pk=report.id)
        self.assertContains(response, 'Iowa')

