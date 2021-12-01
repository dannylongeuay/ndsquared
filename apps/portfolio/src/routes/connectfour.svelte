<script>
	const BOARD_WIDTH = 9;
	const BOARD_HEIGHT = 7;
	let currentColor = 'red';
	let neutralColor = 'gray';
	let board = initBoard(BOARD_WIDTH, BOARD_HEIGHT);
	function initBoard(xdim, ydim) {
		let tempBoard = new Array(ydim);
		for (let y = 0; y < ydim; y++) {
			tempBoard[y] = new Array(xdim);
			for (let x = 0; x < xdim; x++) {
				tempBoard[y][x] = { fillColor: neutralColor };
			}
		}
		return tempBoard;
	}

	function setCellFillColor(x, y, color) {
		if (!isValidMove(x, y, color)) {
			console.log('invalid move', x, y, board[y][x]);
			return;
		}
		board[y][x].fillColor = color;
		if (currentColor === 'red') {
			currentColor = 'yellow';
		} else {
			currentColor = 'red';
		}
	}

	function isValidMove(x, y, color) {
		if (board[y][x].fillColor !== neutralColor) {
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
				if (board[y][x].fillColor !== neutralColor) {
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
	<div class="flex items-center justify-center my-8">
		<h1 class="text-4xl md:text-8xl text-primary">Connect Four</h1>
		<svg class="h-12 mx-2 md:h-32" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
			<circle cx="50" cy="50" r="50" stroke="black" stroke-width="3" fill={currentColor} />
		</svg>
	</div>
	<div
		class="grid max-w-screen-sm grid-cols-9 mx-auto bg-blue-700 md:max-w-screen-md lg:max-w-screen-lg"
	>
		{#each board as boardRows, y}
			{#each boardRows as cell, x}
				<div class="px-1 py-2">
					<svg
						viewBox="0 0 100 100"
						xmlns="http://www.w3.org/2000/svg"
						on:click={() => {
							setCellFillColor(x, y, currentColor);
						}}
					>
						<circle cx="50" cy="50" r="50" fill={cell.fillColor} />
					</svg>
				</div>
			{/each}
		{/each}
	</div>
</div>
