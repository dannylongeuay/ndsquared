import sys

from typing import Optional
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel

from app.connectfour import ConnectFour, DefaultPiece, find_best_move

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_methods=['*'],
)


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


@app.get('/ping')
def pong():
    return {'message': 'pong'}


@app.post('/connectfour')
def connectfour(req: ConnectFourRequest):
    height = len(req.board)
    width = len(req.board[0])
    game = ConnectFour(
        width,
        height,
        board=req.board,
        player=req.player_piece,
        computer=req.computer_piece,
        empty=req.empty_piece,
    )
    val, col = find_best_move(
        game,
        req.depth,
        -sys.maxsize - 1,
        sys.maxsize,
        game.computer,
    )
    return {
        'column': col,
        'value': val,
    }
