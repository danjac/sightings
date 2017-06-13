from rest_framework import serializers

from .models import Report


class ReportSerializer(serializers.ModelSerializer):

    class Meta:
        model = Report

        fields = (
            'id',
            'location',
            'shape',
            'description',
            'latitude',
            'longitude',
            'occurred_at',
            'reported_at',
        )
