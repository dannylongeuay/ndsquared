<script>
	import ImageModal from '$lib/common/imageModal.svelte';

	let visibleImageModals = {};
	function handleClick(i) {
		visibleImageModals = {};
		visibleImageModals[i] = true;
	}

	function reset() {
		visibleImageModals = {};
	}
</script>

<div class="container flex-col justify-center mx-auto">
	{#each Array(32) as _, i}
		<div
			class="flex items-center justify-center min-h-screen border-b-8 border-base-300"
			class:bg-accent={i % 2 === 1}
		>
			<img
				class="border-8 border-neutral-focus"
				src="https://ik.imagekit.io/dannylongeuay/ndsq/nd_fave_{i + 1}.jpg?tr=h-800"
				alt="nd_fave_{i + 1}"
				on:click={() => handleClick(i)}
			/>
		</div>
		<ImageModal
			show={visibleImageModals[i]}
			src="https://ik.imagekit.io/dannylongeuay/ndsq/nd_fave_{i + 1}.jpg"
			alt="nd_fave_{i + 1}"
			on:click={reset}
		/>
	{/each}
</div>
