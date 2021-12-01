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
		<svg height="100" width="100">
			<circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" />
		</svg>
		<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
			<circle cx="50" cy="50" r="50" fill="blue" />
		</svg>
	</div>
	<div class="grid grid-cols-9 bg-base-300">
		<div class="flex justify-center p-1">
			<svg height="100" width="100">
				<circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" />
			</svg>
		</div>
		<div class="flex items-center justify-center p-1">
			<svg height="50" width="50">
				<circle cx="25" cy="25" r="25" stroke="black" stroke-width="3" fill="green" />
			</svg>
		</div>
		<div class="flex items-center justify-center p-1">
			<svg height="20" width="20">
				<circle cx="10" cy="10" r="10" stroke="black" stroke-width="3" fill="purple" />
			</svg>
		</div>
		<div class="flex justify-center p-1">
			<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
				<circle cx="50" cy="50" r="50" fill="blue" />
			</svg>
		</div>
		<div class="flex justify-center p-1">
			<svg viewBox="0 0 50 50" xmlns="http://www.w3.org/2000/svg">
				<circle cx="25" cy="25" r="25" fill="magenta" />
			</svg>
		</div>
		<div class="flex justify-center p-1">
			<svg viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
				<circle cx="10" cy="10" r="10" fill="pink" />
			</svg>
		</div>
	</div>

	<div class="grid grid-cols-9 bg-base-300">
		{#each board as boardRows, y}
			{#each boardRows as cell, x}
				<div class="flex justify-center p-1">
					<svg height="100" width="100">
						<circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" />
					</svg>
					<!-- <svg
						viewBox="0 0 100 100"
						xmlns="http://www.w3.org/2000/svg"
						on:click={() => {
							setCellFillColor(x, y, currentColor);
						}}
					>
						<circle cx="50" cy="50" r="50" fill={cell.fillColor} />
						Unable to render SVG! {cell.fillColor}
					</svg> -->
				</div>
			{/each}
		{/each}
	</div>
</div>
