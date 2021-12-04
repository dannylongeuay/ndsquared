import sys
from enum import Enum
from pprint import pprint
from typing import Tuple, Union, Optional
from copy import deepcopy


class DefaultPiece(Enum):
    EMPTY = '.'
    PLAYER = 'X'
    COMPUTER = 'O'


class ConnectFour:
    def __init__(
        self,
        width: int,
        height: int,
        board: Optional[list[list[str]]] = None,
        player: Optional[str] = DefaultPiece.PLAYER.value,
        computer: Optional[str] = DefaultPiece.COMPUTER.value,
        empty: Optional[str] = DefaultPiece.EMPTY.value,
    ):
        self.width = width
        self.height = height
        self.player = player
        self.computer = computer
        self.empty = empty
        if board is None:
            self.board = self.create_board(width, height)
        else:
            self.board = board

    def create_board(self, width: int, height: int) -> list[list[str]]:
        board = []
        for _ in range(height):
            row = []
            for _ in range(width):
                row.append(self.empty)
            board.append(row)
        return board

    def print_board(self):
        pprint([[str(x + 1) for x in range(self.width)]])
        pprint(self.board)

    def get_next_open_row(self, col: int) -> Union[int, None]:
        for i in range(self.height - 1, -1, -1):
            if self.board[i][col] == self.empty:
                return i
        return None

    def get_valid_moves(self) -> list[int]:
        valid_moves = []
        for col in range(self.width):
            if self.get_next_open_row(col) is None:
                continue
            valid_moves.append(col)
        return valid_moves

    def drop_piece(self, col: int, piece: str):
        row = self.get_next_open_row(col)
        if row is None:
            return
        self.board[row][col] = piece

    def get_vertical_window(
        self,
        col: int,
        row_start: int,
        row_stop: int,
    ) -> list[str]:
        window = []
        for i in range(row_start, row_stop):
            window.append(self.board[i][col])
        return window

    def get_diagonal_window(
        self,
        col_start: int,
        row_start: int,
        direction: int,
    ) -> list[str]:
        window = []
        for i in range(0, 4):
            window.append(self.board[row_start + i * direction][col_start + i])
        return window

    def window_count(self, piece: str, window: list[str]) -> int:
        count = 0
        for window_piece in window:
            if window_piece == piece:
                count += 1
        return count

    def window_score(self, piece: str, window: list[str]) -> int:
        score = 0
        opponent_piece = self.computer
        if piece == self.computer:
            opponent_piece = self.player

        if self.window_count(piece, window) == 4:
            score += 100
        elif self.window_count(piece, window) == 3 and self.window_count(
                self.empty, window) == 1:
            score += 5
        elif self.window_count(piece, window) == 2 and self.window_count(
                self.empty, window) == 2:
            score += 2

        if self.window_count(opponent_piece,
                             window) == 3 and self.window_count(
                                 self.empty, window) == 1:
            score -= 4

        return score

    def evaluate(self, piece: str) -> int:
        score = 0

        # Center
        center_col_window = self.get_vertical_window(
            self.width // 2,
            0,
            self.height,
        )
        center_count = self.window_count(piece, center_col_window)
        score += center_count * 3

        # Horizontal
        for row in range(self.height):
            for col in range(self.width - 3):
                window = self.board[row][col:col + 4]
                score += self.window_score(piece, window)

        # Vertical
        for col in range(self.width):
            for row in range(self.height - 3):
                window = self.get_vertical_window(col, row, row + 4)
                score += self.window_score(piece, window)

        # Positive Diagonal
        for row in range(self.height - 1, 2, -1):
            for col in range(self.width - 3):
                window = self.get_diagonal_window(col, row, -1)
                score += self.window_score(piece, window)

        # Negative Diagonal
        for row in range(self.height - 3):
            for col in range(self.width - 3):
                window = self.get_diagonal_window(col, row, 1)
                score += self.window_score(piece, window)

        return score

    def is_winning_move(self, piece: str) -> bool:
        # Horizontal
        for row in range(self.height):
            for col in range(self.width - 3):
                window = self.board[row][col:col + 4]
                if self.window_count(piece, window) == 4:
                    return True

        # Vertical
        for col in range(self.width):
            for row in range(self.height - 3):
                window = self.get_vertical_window(col, row, row + 4)
                if self.window_count(piece, window) == 4:
                    return True

        # Positive Diagonal
        for row in range(self.height - 1, 2, -1):
            for col in range(self.width - 3):
                window = self.get_diagonal_window(col, row, -1)
                if self.window_count(piece, window) == 4:
                    return True

        # Negative Diagonal
        for row in range(self.height - 3):
            for col in range(self.width - 3):
                window = self.get_diagonal_window(col, row, 1)
                if self.window_count(piece, window) == 4:
                    return True

        return False


def minimax(
    game: ConnectFour,
    depth: int,
    alpha: int,
    beta: int,
    maximizer: str,
) -> Tuple[int, int]:
    return_value = sys.maxsize
    if maximizer == game.computer:
        return_value = -sys.maxsize - 1
    player_won = game.is_winning_move(game.player)
    computer_won = game.is_winning_move(game.computer)
    valid_moves = game.get_valid_moves()

    if computer_won:
        return sys.maxsize, -1
    if player_won:
        return -sys.maxsize - 1, -1
    if depth == 0:
        return game.evaluate(game.computer), -1
    if len(valid_moves) == 0:
        return 0, -1

    return_col = valid_moves[0]

    for col in valid_moves:
        new_board = deepcopy(game.board)
        new_game = ConnectFour(
            game.width,
            game.height,
            new_board,
            player=game.player,
            computer=game.computer,
            empty=game.empty,
        )
        new_game.drop_piece(col, maximizer)
        if new_game.is_winning_move(maximizer):
            return_col = col
            return_value = sys.maxsize
            if maximizer == game.player:
                return_value = -sys.maxsize - 1
            break
        next_maximizer = game.computer
        if maximizer == game.computer:
            next_maximizer = game.player
        value, _ = minimax(new_game, depth - 1, alpha, beta, next_maximizer)
        if maximizer == game.computer:
            if value > return_value:
                return_value = value
                return_col = col
            if return_value > alpha:
                alpha = return_value
            if alpha >= beta:
                break
        else:
            if value < return_value:
                return_value = value
                return_col = col
            if return_value < beta:
                beta = value
            if beta <= alpha:
                break

    return return_value, return_col


def find_best_move(
    game: ConnectFour,
    depth: int,
    alpha: int,
    beta: int,
    maximizer: str,
) -> Tuple[int, int]:
    valid_moves = game.get_valid_moves()

    if len(valid_moves) == 0:
        return 0, -1

    for col in valid_moves:
        new_board = deepcopy(game.board)
        new_game = ConnectFour(
            game.width,
            game.height,
            new_board,
            player=game.player,
            computer=game.computer,
            empty=game.empty,
        )
        new_game.drop_piece(col, maximizer)
        if new_game.is_winning_move(maximizer):
            return sys.maxsize, col

    return minimax(game, depth, alpha, beta, maximizer)


def play(game: ConnectFour):
    try:
        while True:
            game.print_board()
            valid_moves = game.get_valid_moves()
            try:
                raw_col = input("Enter column number: ")
                col = int(raw_col) - 1
            except ValueError:
                print(f'{raw_col} is not a valid column number')
                continue
            if col not in valid_moves:
                print(f'{col} is not a valid move')
                continue
            game.drop_piece(col, DefaultPiece.PLAYER.value)
            player_won = game.is_winning_move(DefaultPiece.PLAYER.value)
            if player_won:
                print('Player has won!')
                game.print_board()
                break
            game.print_board()
            print('Computer Thinking...')
            val, col = find_best_move(
                game,
                5,
                -sys.maxsize - 1,
                sys.maxsize,
                DefaultPiece.COMPUTER.value,
            )
            print(val, col)
            game.drop_piece(col, DefaultPiece.COMPUTER.value)
            computer_won = game.is_winning_move(DefaultPiece.COMPUTER.value)
            if computer_won:
                print('Computer has won!')
                game.print_board()
                break
    except KeyboardInterrupt:
        print('\nKeyboard interrupt detected, exiting game...')


if __name__ == '__main__':
    init_game = ConnectFour(9, 7)
    play(init_game)
