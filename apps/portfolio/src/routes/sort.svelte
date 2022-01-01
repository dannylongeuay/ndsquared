<script>
	import { flip } from 'svelte/animate';
	import { quintOut } from 'svelte/easing';

	let inPlace = [
		{
			val: 3,
			height: 'h-32',
			bgColor: 'bg-green-300',
			textColor: 'text-green-700'
		},
		{
			val: 8,
			height: 'h-80',
			bgColor: 'bg-violet-300',
			textColor: 'text-violet-700'
		},
		{
			val: 9,
			height: 'h-96',
			bgColor: 'bg-lime-300',
			textColor: 'text-lime-700'
		},
		{
			val: 5,
			height: 'h-52',
			bgColor: 'bg-red-300',
			textColor: 'text-red-700'
		},
		{
			val: 2,
			height: 'h-24',
			bgColor: 'bg-pink-300',
			textColor: 'text-pink-700'
		},
		{
			val: 4,
			height: 'h-40',
			bgColor: 'bg-fuchsia-300',
			textColor: 'text-fuchsia-700'
		},
		{
			val: 6,
			height: 'h-60',
			bgColor: 'bg-blue-300',
			textColor: 'text-blue-700'
		},
		{
			val: 1,
			height: 'h-10',
			bgColor: 'bg-sky-300',
			textColor: 'text-sky-700'
		},
		{
			val: 7,
			height: 'h-72',
			bgColor: 'bg-orange-300',
			textColor: 'text-orange-700'
		}
	];

	function swap(arr, i, j) {
		// console.log(`swapped ${arr[i].val} with ${arr[j].val}`);
		[arr[i], arr[j]] = [arr[j], arr[i]];
	}

	function sleep(ms) {
		return new Promise((resolve) => setTimeout(resolve, ms));
	}

	function* bubbleSort(arr) {
		for (let i = 0; i < arr.length - 1; i++) {
			for (let j = 0; j < arr.length - 1; j++) {
				if (arr[j].val > arr[j + 1].val) {
					swap(arr, j, j + 1);
					yield arr;
				}
			}
		}
		return arr;
	}

	function* selectionSort(arr) {
		for (let i = 0; i < arr.length; i++) {
			let minIndex = i;
			for (let j = i + 1; j < arr.length; j++) {
				if (arr[minIndex].val > arr[j].val) {
					minIndex = j;
				}
			}
			if (arr[i].val !== arr[minIndex].val) {
				swap(arr, i, minIndex);
				yield arr;
			}
		}
		return arr;
	}

	function* insertionSort(arr) {
		for (let i = 1; i < arr.length; i++) {
			const key = arr[i];
			let j = i - 1;
			let swapped = false;
			while (j >= 0 && key.val < arr[j].val) {
				swap(arr, j + 1, j);
				j -= 1;
				swapped = true;
			}
			if (swapped) {
				yield arr;
			}
		}
		return arr;
	}

	let selectedAlgorithm = bubbleSort;
	let selectedArr = inPlace;
	const selectedMap = {
		bubbleSort: {
			alg: bubbleSort,
			arr: inPlace
		},
		selectionSort: {
			alg: selectionSort,
			arr: inPlace
		},
		insertionSort: {
			alg: insertionSort,
			arr: inPlace
		}
	};

	let selected;

	const sort = async () => {
		const sortIter = selectedAlgorithm(selectedArr);
		let next = sortIter.next();
		while (!next.done) {
			selectedArr = next.value;
			await sleep(1000);
			next = sortIter.next();
		}
	};

	const shuffle = () => {
		for (let i = selectedArr.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			swap(selectedArr, i, j);
		}
		selectedArr = selectedArr;
	};

	const select = () => {
		selectedAlgorithm = selectedMap[selected].alg;
		selectedArr = selectedMap[selected].arr;
	};
</script>

<div class="items-center justify-center m-4 md:flex">
	<div class="flex justify-center">
		<button on:click={sort} class="m-2 btn btn-primary">Sort</button>
		<button on:click={shuffle} class="m-2 btn btn-primary">Shuffle</button>
	</div>
	<select
		bind:value={selected}
		on:change={select}
		class="w-full max-w-xs m-4 select select-bordered select-primary"
	>
		<option value="bubbleSort">Bubble Sort</option>
		<option value="selectionSort">Selection Sort</option>
		<option value="insertionSort">Insertion Sort</option>
	</select>
</div>
<div class="flex items-end justify-center overflow-x-auto">
	{#each selectedArr as item (item)}
		<div animate:flip>
			<div
				class="{item.bgColor} {item.height} {item.textColor} m-1 p-1 md:m-2 md:p-4 text-xl font-bold"
			>
				{item.val}
			</div>
		</div>
	{/each}
</div>
