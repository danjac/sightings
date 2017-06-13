from django.contrib import admin

from .models import Report


class ReportAdmin(admin.ModelAdmin):

    list_display = (
        'location',
        'shape',
        'reported_at',
    )

    exclude = (
        'search_vector',
    )

    date_hierarchy = 'reported_at'


admin.site.register(Report, ReportAdmin)
