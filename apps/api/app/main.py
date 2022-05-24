import logging
import time
from asyncio import sleep
from enum import Enum

import dramatiq
from dramatiq import Worker
from dramatiq.brokers.redis import RedisBroker
from dramatiq.brokers.stub import StubBroker
from dramatiq.results import Results
from dramatiq.results.backends import RedisBackend, StubBackend
from dramatiq.results.errors import ResultMissing
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, BaseSettings


class Environment(Enum):
    PRODUCTION = "production"
    DEVELOPMENT = "development"
    TESTING = "testing"


class Settings(BaseSettings):
    REDIS_URL: str = "redis://user:pass@localhost:6379/1"
    API_ENVIRONMENT: Environment = Environment.PRODUCTION


app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

result_backend = RedisBackend(url=Settings().REDIS_URL)
broker = RedisBroker(
    url=Settings().REDIS_URL,
    middleware=[],
)

STUB_ENVIRONMENTS = (Environment.DEVELOPMENT, Environment.TESTING)

if Settings().API_ENVIRONMENT in STUB_ENVIRONMENTS:
    broker = StubBroker(middleware=[])
    result_backend = StubBackend()

if Settings().API_ENVIRONMENT == Environment.DEVELOPMENT:
    worker = Worker(broker, worker_timeout=100)
    worker.start()

broker.add_middleware(Results(backend=result_backend))
dramatiq.set_broker(broker)


class PingResponse(BaseModel):
    message: str

    class Config:
        schema_extra = {
            "example": {
                "message": "pong",
            }
        }


class DelayRequest(BaseModel):
    amount: int

    class Config:
        schema_extra = {
            "example": {
                "amount": 3,
            }
        }


class DelayResponse(BaseModel):
    message: str

    class Config:
        schema_extra = {
            "example": {
                "message": "Delayed for 3 second(s)",
            }
        }


@app.get("/ping", response_model=PingResponse)
def ping():
    return PingResponse(message="pong")


@app.post("/delay", response_model=DelayResponse)
async def delay(req: DelayRequest):
    logging.info(req)
    actor = delay_actor.send(req.dict())
    result = None
    while result is None:
        try:
            result = actor.get_result()
        except ResultMissing as err:
            logging.debug("unable to retrieve result for: %s", err)
        await sleep(0.5)
    return DelayResponse(message=result.get("message"))


@dramatiq.actor(store_results=True)
def delay_actor(data: dict) -> dict:
    amount = data.get("amount")
    time.sleep(amount)
    return {"message": f"Delayed for {amount} second(s)"}
