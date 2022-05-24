import unittest

from fastapi.testclient import TestClient

from app.main import DelayRequest, delay_actor


def test_ping(client: TestClient):
    resp = client.get("/ping")

    assert resp.status_code == 200
    assert resp.json() == {"message": "pong"}


def test_delay(client: TestClient, stub_broker, stub_worker):
    req = DelayRequest(amount=1)
    resp = client.post(
        "/delay",
        json=req.dict(),
    )
    stub_broker.join(delay_actor.queue_name)
    stub_worker.join()

    assert resp.status_code == 200
    assert resp.json() == {"message": "Delayed for 1 second(s)"}


class TestFoo(unittest.TestCase):
    def test_bar(self):
        self.assertTrue(True)
        self.assertFalse(False)
