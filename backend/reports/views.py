from rest_framework import viewsets

from .models import Report
from .serializers import ReportSerializer


class ReportViewSet(viewsets.ReadOnlyModelViewSet):

    serializer_class = ReportSerializer
    paginate_by = 20

    def get_queryset(self):

        qs = Report.objects.get_queryset().order_by('-occurred_at')
        search_term = self.request.GET.get('s', '').strip()
        if search_term:
            qs = qs.search(search_term)

        return qs
