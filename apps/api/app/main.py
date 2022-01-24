import os
import sys
import dramatiq
import logging

from asyncio import sleep
from typing import Optional
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi_auth0 import Auth0
from pydantic import BaseModel
from dramatiq.brokers.redis import RedisBroker
from dramatiq.results.backends import RedisBackend
from dramatiq.results import Results
from dramatiq.results.errors import ResultMissing

from app.connectfour import ConnectFour, DefaultPiece, find_best_move

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_methods=['*'],
    allow_headers=['*'],
)
AUTH0_DOMAIN = os.getenv('AUTH0_DOMAIN', 'notarealauth0domain')
AUTH0_API_AUDIENCE = os.getenv('AUTH0_API_AUDIENCE',
                               'notarealauth0apiaudience')

auth = Auth0(
    domain=AUTH0_DOMAIN,
    api_audience=AUTH0_API_AUDIENCE,
    scopes={
        'read:connectfour': '',
    },
)

REDIS_URL = os.getenv(
    'REDIS_URL',
    'notavalidredisurl',
)

result_backend = RedisBackend(url=REDIS_URL)
broker = RedisBroker(
    url=REDIS_URL,
    middleware=[],
)
broker.add_middleware(Results(backend=result_backend))
dramatiq.set_broker(broker)


class ConnectFourRequest(BaseModel):
    player_piece: Optional[str] = DefaultPiece.PLAYER.value
    computer_piece: Optional[str] = DefaultPiece.COMPUTER.value
    empty_piece: Optional[str] = DefaultPiece.EMPTY.value
    depth: Optional[int] = 4
    board: list[list[str]]

    class Config:
        schema_extra = {
            'example': {
                'player_piece':
                'X',
                'computer_piece':
                'O',
                'empty_piece':
                '.',
                'depth':
                4,
                'board': [
                    ['.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.'],
                    ['.', '.', 'X', '.', '.'],
                ],
            }
        }


class ConnectFourResponse(BaseModel):
    column: int
    value: int


class ConnectFourMessagePayload(BaseModel):
    width: int
    height: int
    board: list[list[str]]
    player_piece: str
    computer_piece: str
    empty_piece: str
    depth: int


@app.get('/ping')
def pong():
    return {'message': 'pong'}


@app.post('/connectfour')
async def connectfour(req: ConnectFourRequest):
    height = len(req.board)
    width = len(req.board[0])
    connectfour_msg_payload = ConnectFourMessagePayload(
        width=width,
        height=height,
        board=req.board,
        player_piece=req.player_piece,
        computer_piece=req.computer_piece,
        empty_piece=req.empty_piece,
        depth=req.depth,
    )
    message = connectfour_best_move.send(connectfour_msg_payload.dict())
    result = None
    retries = 40  # ~20 seconds
    while result is None:
        try:
            result = message.get_result()
        except ResultMissing as err:
            logging.debug('unable to retreive result for: %s', err)
        await sleep(0.5)
        retries -= 1
        if retries <= 0:
            break

    response_data = {
        'column': 0,
        'value': -1,
    }
    if result:
        response_data['column'] = result.get('column')
        response_data['value'] = result.get('value')

    return ConnectFourResponse(**response_data)


@dramatiq.actor(store_results=True)
def connectfour_best_move(data: ConnectFourMessagePayload):
    payload = ConnectFourMessagePayload(**data)
    game = ConnectFour(
        payload.width,
        payload.height,
        board=payload.board,
        player=payload.player_piece,
        computer=payload.computer_piece,
        empty=payload.empty_piece,
    )
    val, col = find_best_move(
        game,
        payload.depth,
        -sys.maxsize - 1,
        sys.maxsize,
        game.computer,
    )
    connectfour_response = ConnectFourResponse(
        column=col,
        value=val,
    )
    return connectfour_response.dict()
