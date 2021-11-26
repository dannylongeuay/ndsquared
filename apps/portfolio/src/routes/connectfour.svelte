<script>
	const BOARD_WIDTH = 9;
	const BOARD_HEIGHT = 7;
	let board = initBoard(BOARD_WIDTH, BOARD_HEIGHT);
	let currentColor = 'red';
	function initBoard(xdim, ydim) {
		let tempBoard = new Array(ydim);
		for (let y = 0; y < ydim; y++) {
			tempBoard[y] = new Array(xdim);
			for (let x = 0; x < xdim; x++) {
				tempBoard[y][x] = { fillColor: 'white' };
			}
		}
		return tempBoard;
	}

	function setCellFillColor(x, y, color) {
		if (!isValidMove(x, y, color)) {
			console.log('invalid move', x, y, board[y][x]);
			return;
		}
		// console.log(x, y, board[y][x]);
		board[y][x].fillColor = color;
		if (currentColor === 'red') {
			currentColor = 'blue';
		} else {
			currentColor = 'red';
		}
	}

	function isValidMove(x, y, color) {
		if (board[y][x].fillColor !== 'white') {
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
				if (board[y][x].fillColor !== 'white') {
					continue;
				}
				foundValidMove[x] = true;
				validMoves[`${x},${y}`] = true;
			}
		}
		return validMoves;
	}
</script>

<div class="container mx-auto">
	<div class="flex justify-center mb-8">
		<h1 class="text-4xl md:text-8xl text-secondary-focus">Connect Four</h1>
	</div>
	<div class="grid object-center grid-cols-9 bg-base-300">
		{#each board as boardRows, y}
			{#each boardRows as cell, x}
				<div class="flex justify-center p-1">
					<svg
						viewBox="0 0 100 100"
						fill={cell.fillColor}
						xmlns="http://www.w3.org/2000/svg"
						on:click={() => {
							setCellFillColor(x, y, currentColor);
						}}
					>
						<circle cx="50" cy="50" r="50" stroke="black" stroke-width="3" />
					</svg>
				</div>
			{/each}
		{/each}
	</div>
</div>
