<script>
	import { appSettingsStore } from '$lib/stores/appsettings';

	let appSettings;

	appSettingsStore.subscribe((value) => {
		appSettings = value;
	});

	const BOARD_WIDTH = 9;
	const BOARD_HEIGHT = 7;
	const PLAYER_COLOR = 'red';
	const OPPONENT_COLOR = 'yellow';
	const NEUTRUAL_COLOR = 'gray';
	const MODE_ONE_PLAYER = '1 Player';
	const MODE_TWO_PLAYER = '2 Players';
	const DIFFICULTY_EASY = 3;
	const DIFFICULTY_NORMAL = 4;
	const DIFFICULTY_HARD = 5;
	const GAME_PAUSED_MSGS = [
		'I am plotting your defeat...',
		'Please wait while the hamster wheel turns...',
		'I can see the future, can you?',
		'Minimax with alpha-beta pruning calculation in progress...',
		'Give me a minute I am plotting mass destruction...',
		'Hold my pencil while I think...',
		'NO NOT THERE!',
		'It was not your turn, but fine...',
		'Are you sure about that move?',
		'But that is where I wanted to go!',
		'Hold that thought...',
		'Well this is going to be easy...',
		'That move must have been an accident...',
		'Gonna make a sandwich, be right back...',
		'Give me a minute to think of some more clever lines...',
		'I have thought of the perfect move, hold my beer...',
		'Is that all you got?',
		'I thought this was supposed to be a challenge...',
		'Do not say I never gave you a chance',
		'You really think that is your best option?',
		'Anywhere but there!',
		'I would not do that if I were you...',
		'Did the hamster tell you to go there?',
		'I have run out of clever remarks'
	];

	let gamePaused = false;
	let gameOver = false;
	let currentColor = PLAYER_COLOR;
	let winningMsg = '';
	let board = initBoard(BOARD_WIDTH, BOARD_HEIGHT);
	let mode = MODE_ONE_PLAYER;
	let difficulty = DIFFICULTY_NORMAL;
	let gamePausedMsg = 'Computer is thinking...';

	function initBoard(xdim, ydim) {
		let tempBoard = new Array(ydim);
		for (let y = 0; y < ydim; y++) {
			tempBoard[y] = new Array(xdim);
			for (let x = 0; x < xdim; x++) {
				tempBoard[y][x] = { fillColor: NEUTRUAL_COLOR };
			}
		}
		return tempBoard;
	}

	function resetBoard() {
		board = initBoard(BOARD_WIDTH, BOARD_HEIGHT);
		gamePaused = false;
		gameOver = false;
		currentColor = PLAYER_COLOR;
		winningMsg = '';
	}

	function getGameBoard(xdim, ydim) {
		let tempBoard = new Array(ydim);
		for (let y = 0; y < ydim; y++) {
			tempBoard[y] = new Array(xdim);
			for (let x = 0; x < xdim; x++) {
				switch (board[y][x].fillColor) {
					case PLAYER_COLOR:
						tempBoard[y][x] = 'X';
						break;
					case OPPONENT_COLOR:
						tempBoard[y][x] = 'O';
						break;
					default:
						tempBoard[y][x] = '.';
				}
			}
		}
		return tempBoard;
	}

	function getOpenRow(x) {
		for (let y = BOARD_HEIGHT - 1; y >= 0; y--) {
			if (board[y][x].fillColor === NEUTRUAL_COLOR) {
				return y;
			}
		}
		return null;
	}

	function dropPiece(x, color) {
		let y = getOpenRow(x);
		if (y === null) {
			return;
		}
		board[y][x].fillColor = color;
	}

	async function handleTurn(x, y, color) {
		if (gamePaused) return;
		if (!isValidMove(x, y)) {
			console.log('invalid move', x, y, board[y][x]);
			return;
		}
		board[y][x].fillColor = color;

		if (isWinningMove(color)) {
			gamePaused = true;
			gameOver = true;
			gamePausedMsg = `${color.toUpperCase()} won!`;
			return;
		}
		if (currentColor === PLAYER_COLOR) {
			currentColor = OPPONENT_COLOR;
		} else {
			currentColor = PLAYER_COLOR;
		}

		if (mode === MODE_TWO_PLAYER) {
			return;
		}

		const gameBoard = getGameBoard(BOARD_WIDTH, BOARD_HEIGHT);
		const payload = {
			board: gameBoard,
			depth: difficulty
		};
		gamePaused = true;
		gamePausedMsg = GAME_PAUSED_MSGS[Math.floor(Math.random() * GAME_PAUSED_MSGS.length)];
		const connect_four_url = appSettings.API_BASE_URL + '/connectfour';
		const response = await fetch(connect_four_url, {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			body: JSON.stringify(payload)
		});
		const data = await response.json();

		color = OPPONENT_COLOR;

		dropPiece(data.column, color);

		if (isWinningMove(color)) {
			gamePaused = true;
			gameOver = true;
			gamePausedMsg = `${color.toUpperCase()} won!`;
			return;
		}
		if (currentColor === PLAYER_COLOR) {
			currentColor = OPPONENT_COLOR;
		} else {
			currentColor = PLAYER_COLOR;
		}
		gamePaused = false;
	}

	function isValidMove(x, y) {
		if (board[y][x].fillColor !== NEUTRUAL_COLOR) {
			return false;
		}
		if (!(`${x},${y}` in getValidMoves())) {
			return false;
		}
		return true;
	}

	function getValidMoves() {
		let validMoves = {};
		let foundValidMove = {};
		for (let y = BOARD_HEIGHT - 1; y >= 0; y--) {
			for (let x = 0; x < BOARD_WIDTH; x++) {
				if (x in foundValidMove) {
					continue;
				}
				if (board[y][x].fillColor !== NEUTRUAL_COLOR) {
					continue;
				}
				foundValidMove[x] = true;
				validMoves[`${x},${y}`] = true;
			}
		}
		return validMoves;
	}

	function windowCount(window, color) {
		let count = 0;
		for (let cell of window) {
			if (cell.fillColor === color) {
				count++;
			}
		}
		return count;
	}

	function getVerticalWindow(x, yStart, yEnd) {
		let window = [];
		for (let y = yStart; y < yEnd; y++) {
			window.push(board[y][x]);
		}
		return window;
	}

	function getDiagonalWindow(xStart, yStart, direction) {
		let window = [];
		for (let i = 0; i < 4; i++) {
			window.push(board[yStart + i * direction][xStart + i]);
		}
		return window;
	}

	function checkWinningCount(window, color) {
		let count = windowCount(window, color);
		if (count === 4) {
			for (let cell of window) {
				cell.fillColor = 'black';
			}
			return true;
		}
		return false;
	}

	function isWinningMove(color) {
		// Horizontal
		for (let y = 0; y < BOARD_HEIGHT; y++) {
			for (let x = 0; x < BOARD_WIDTH - 3; x++) {
				let window = board[y].slice(x, x + 4);
				if (checkWinningCount(window, color)) {
					return true;
				}
			}
		}

		// Vertical
		for (let x = 0; x < BOARD_WIDTH; x++) {
			for (let y = 0; y < BOARD_HEIGHT - 3; y++) {
				let window = getVerticalWindow(x, y, y + 4);
				if (checkWinningCount(window, color)) {
					return true;
				}
			}
		}

		// Positive Diagonal
		for (let y = BOARD_HEIGHT - 1; y > 2; y--) {
			for (let x = 0; x < BOARD_WIDTH - 3; x++) {
				let window = getDiagonalWindow(x, y, -1);
				if (checkWinningCount(window, color)) {
					return true;
				}
			}
		}

		// Negative Diagonal
		for (let y = 0; y < BOARD_HEIGHT - 3; y++) {
			for (let x = 0; x < BOARD_WIDTH - 3; x++) {
				let window = getDiagonalWindow(x, y, 1);
				if (checkWinningCount(window, color)) {
					return true;
				}
			}
		}
		return false;
	}
</script>

{#if gamePaused}
	<div class="fixed inset-x-0">
		<div class="flex justify-center items-center bg-neutral bg-opacity-95 rounded-2xl mx-4 py-4">
			<h1 class="text-md md:text-4xl text-neutral-content text-center">{gamePausedMsg}</h1>
			{#if gameOver}
				<button class="btn btn-primary mx-4" on:click={resetBoard}>Reset</button>
			{:else}
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="currentColor"
					class="mx-4 w-10 h-10"
					viewBox="0 0 16 16"
				>
					<path
						d="M2 1.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-1v1a4.5 4.5 0 0 1-2.557 4.06c-.29.139-.443.377-.443.59v.7c0 .213.154.451.443.59A4.5 4.5 0 0 1 12.5 13v1h1a.5.5 0 0 1 0 1h-11a.5.5 0 1 1 0-1h1v-1a4.5 4.5 0 0 1 2.557-4.06c.29-.139.443-.377.443-.59v-.7c0-.213-.154-.451-.443-.59A4.5 4.5 0 0 1 3.5 3V2h-1a.5.5 0 0 1-.5-.5zm2.5.5v1a3.5 3.5 0 0 0 1.989 3.158c.533.256 1.011.791 1.011 1.491v.702c0 .7-.478 1.235-1.011 1.491A3.5 3.5 0 0 0 4.5 13v1h7v-1a3.5 3.5 0 0 0-1.989-3.158C8.978 9.586 8.5 9.052 8.5 8.351v-.702c0-.7.478-1.235 1.011-1.491A3.5 3.5 0 0 0 11.5 3V2h-7z"
					/>
				</svg>
			{/if}
		</div>
	</div>
{/if}
<div class="container mx-auto">
	<div class="flex items-center justify-center my-8">
		{#if winningMsg}
			<h1 class="text-4xl md:text-8xl text-primary">{winningMsg.toUpperCase()}</h1>
		{:else}
			<h1 class="text-4xl md:text-8xl text-primary">Connect Four</h1>
		{/if}
		<svg class="h-12 mx-2 md:h-32" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
			<circle cx="50" cy="50" r="50" stroke="black" stroke-width="3" fill={currentColor} />
		</svg>
	</div>
	<div class="md:flex justify-center">
		<div class="flex items-center justify-center mb-4">
			<div><h2 class="mx-2 text-lg md:text-2xl text-secondary">Mode:</h2></div>
			<div class="mx-2">
				<button
					class:btn-primary={mode === MODE_ONE_PLAYER}
					class="btn"
					on:click={() => {
						mode = MODE_ONE_PLAYER;
						resetBoard();
					}}>{MODE_ONE_PLAYER}</button
				>
				<button
					class:btn-primary={mode === MODE_TWO_PLAYER}
					class="btn"
					on:click={() => {
						mode = MODE_TWO_PLAYER;
						resetBoard();
					}}>{MODE_TWO_PLAYER}</button
				>
			</div>
		</div>
		<div class="flex items-center justify-center mb-4">
			<div><h2 class="mx-2 text-lg md:text-2xl text-secondary">Difficulty:</h2></div>
			<div class="mx-2">
				<button
					class:btn-primary={difficulty === DIFFICULTY_EASY}
					class="btn"
					on:click={() => {
						difficulty = DIFFICULTY_EASY;
						resetBoard();
					}}>Easy</button
				>
				<button
					class:btn-primary={difficulty === DIFFICULTY_NORMAL}
					class="btn"
					on:click={() => {
						difficulty = DIFFICULTY_NORMAL;
						resetBoard();
					}}>Normal</button
				>
				<button
					class:btn-primary={difficulty === DIFFICULTY_HARD}
					class="btn"
					on:click={() => {
						difficulty = DIFFICULTY_HARD;
						resetBoard();
					}}>Hard</button
				>
			</div>
		</div>
	</div>
	<div
		class="grid max-w-screen-sm grid-cols-9 mx-auto border-4 border-primary bg-blue-700 rounded-2xl md:max-w-screen-md lg:max-w-screen-lg"
	>
		{#each board as boardRows, y}
			{#each boardRows as cell, x}
				<div class="px-1 py-2">
					<svg
						viewBox="0 0 100 100"
						xmlns="http://www.w3.org/2000/svg"
						on:click={() => {
							handleTurn(x, y, currentColor);
						}}
					>
						<circle cx="50" cy="50" r="50" fill={cell.fillColor} />
					</svg>
				</div>
			{/each}
		{/each}
	</div>
	<div class="flex justify-center my-4">
		<button class="btn btn-primary" on:click={resetBoard}>Reset</button>
	</div>
</div>
