import os
import pytest

from sightings.app import app
from sightings.app import db as _db


class TestConfig:
    TESTING = True
    SQLALCHEMY_DATABASE_URI = os.environ["TEST_DATABASE_URL"]
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    DEBUG = False


def truncate_all():
    """
    Empties all the tables, so we have clean data for each test.
    """
    for table in _db.metadata.tables.keys():
        _db.engine.execute("truncate %s" % table)


@pytest.fixture(scope="session")
def flask_app():
    app.config.from_object(TestConfig)
    with app.app_context():
        _db.create_all()
        yield app

        # call this to prevent hanging
        _db.session.close()
        _db.drop_all()


@pytest.fixture
def db():
    truncate_all()
    return _db


@pytest.fixture
def client(flask_app):
    with flask_app.app_context():
        yield flask_app.test_client()
