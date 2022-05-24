import pytest
from dramatiq import Worker
from fastapi.testclient import TestClient

from app.main import app, broker


@pytest.fixture()
def stub_broker():
    broker.flush_all()
    return broker


@pytest.fixture()
def stub_worker():
    worker = Worker(broker, worker_timeout=100)
    worker.start()
    yield worker
    worker.stop()


@pytest.fixture()
def client():
    client = TestClient(app)
    return client
