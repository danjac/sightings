import decimal
import json

from sightings.models import Report


def test_ping(client):
    response = client.get("/")
    assert response.status_code == 200


def test_report_list(db, client):

    for i in range(5):
        report = Report(
            location="Roswell, NM",
            description="test %d" % i,
            latitude=decimal.Decimal('56.18989'),
            longitude=decimal.Decimal('33.18989'),
        )

        db.session.add(report)
    db.session.commit()

    response = client.get("/reports/")
    assert response.status_code == 200
    payload = json.loads(response.data)

    assert payload['total'] == 5
    assert payload['pages'] == 1
    assert payload['per_page'] == 20

    assert not payload['has_prev']
    assert not payload['has_next']

    assert len(payload['items']) == 5


def test_report_detail(db, client):
    report = Report(
            location="Roswell, NM",
            description="test",
            latitude=decimal.Decimal('56.18989'),
            longitude=decimal.Decimal('33.18989'),
        )

    db.session.add(report)
    db.session.commit()

    response = client.get("/reports/%d/" % report.id)
    assert response.status_code == 200
    payload = json.loads(response.data)
    assert payload['location'] == "Roswell, NM"


