from django.db import models


class ReportQuerySet(models.QuerySet):

    def search(self, search_term):
        return self.filter(location__icontains=search_term)


class Report(models.Model):

    location = models.CharField(max_length=250, blank=True)
    shape = models.CharField(max_length=50, blank=True)
    duration = models.CharField(max_length=50, blank=True)
    description = models.TextField(blank=True)

    latitude = models.DecimalField(
        null=True,
        blank=True,
        decimal_places=2,
        max_digits=10,
    )

    longitude = models.DecimalField(
        null=True,
        blank=True,
        decimal_places=2,
        max_digits=10,
    )

    reported_at = models.DateField(null=True, blank=True)
    occurred_at = models.DateField(null=True, blank=True)

    objects = ReportQuerySet.as_manager()
