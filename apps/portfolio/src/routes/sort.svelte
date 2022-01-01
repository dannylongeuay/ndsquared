<script>
	import { flip } from 'svelte/animate';

	const primaryBorderColorClass = 'border-yellow-300';
	const secondaryBorderColorClass = 'border-cyan-400';

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

	let highlightClasses = {};

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
					highlightClasses = {};
					highlightClasses[arr[j + 1].val] = `border-2 ${primaryBorderColorClass}`;
					highlightClasses[arr[j].val] = `border-2 ${secondaryBorderColorClass}`;
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
				highlightClasses = {};
				highlightClasses[arr[i].val] = `border-2 ${primaryBorderColorClass}`;
				highlightClasses[arr[minIndex].val] = `border-2 ${secondaryBorderColorClass}`;
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
			highlightClasses = {};
			while (j >= 0 && key.val < arr[j].val) {
				swap(arr, j + 1, j);
				highlightClasses[arr[j].val] = `border-2 ${primaryBorderColorClass}`;
				highlightClasses[arr[j + 1].val] = `border-2 ${secondaryBorderColorClass}`;
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
	let selectedInterval = 1000;

	const selectedMap = {
		bubbleSort: {
			alg: bubbleSort,
			arr: inPlace,
			interval: 1000,
			legendInfo: [
				{
					classes: `border-4 ${primaryBorderColorClass}`,
					text: 'Lower Swap'
				},
				{
					classes: `border-4 ${secondaryBorderColorClass}`,
					text: 'Upper Swap'
				}
			]
		},
		selectionSort: {
			alg: selectionSort,
			arr: inPlace,
			interval: 2000,
			legendInfo: [
				{
					classes: `border-4 ${primaryBorderColorClass}`,
					text: 'Min Element'
				},
				{
					classes: `border-4 ${secondaryBorderColorClass}`,
					text: 'Swap Element'
				}
			]
		},
		insertionSort: {
			alg: insertionSort,
			arr: inPlace,
			interval: 2000,
			legendInfo: [
				{
					classes: `border-4 ${primaryBorderColorClass}`,
					text: 'Key Element'
				},
				{
					classes: `border-4 ${secondaryBorderColorClass}`,
					text: 'Shift Element(s)'
				}
			]
		}
	};

	let selected;

	const sort = async () => {
		const sortIter = selectedAlgorithm(selectedArr);
		let next = sortIter.next();
		while (!next.done) {
			selectedArr = next.value;
			await sleep(selectedInterval);
			next = sortIter.next();
		}
		highlightClasses = {};
		selectedArr = next.value;
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
		selectedInterval = selectedMap[selected].interval;
	};

	const getHighlightClasses = (val) => {
		if (val in highlightClasses) {
			return highlightClasses[val];
		}
		return '';
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
			<div class={getHighlightClasses(item.val)}>
				<div
					class="{item.bgColor} {item.height} {item.textColor} m-1 p-1 md:m-2 md:p-4 text-xl font-bold"
				>
					{item.val}
				</div>
			</div>
		</div>
	{/each}
</div>
{#if selected}
	<div class="flex justify-center">
		<div class="flex-col px-4 m-4 border-2 rounded-lg bg-base-200 border-secondary">
			<h1 class="py-4 text-2xl text-center text-primary">Legend</h1>
			<div class="flex py-4 justify-evenly">
				{#each selectedMap[selected].legendInfo as info}
					<div class="flex flex-col items-center">
						<div class="w-6 h-20 {info.classes}" />
						<h3 class="p-2 text-base-content">{info.text}</h3>
					</div>
				{/each}
			</div>
		</div>
	</div>
{/if}
