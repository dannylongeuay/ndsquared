<script>
	export let show = false;
	export let src;
	export let alt;
</script>

{#if show}
	<div
		class="fixed top-0 left-0 z-20 flex justify-center w-screen h-screen bg-opacity-90 bg-base-300"
	>
		<button on:click class="fixed top-0 right-0 mt-2 mr-2 btn btn-primary btn-circle">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="w-5 h-5"
				viewBox="0 0 20 20"
				fill="currentColor"
			>
				<path
					fill-rule="evenodd"
					d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
					clip-rule="evenodd"
				/>
			</svg>
		</button>
		<div class="flex items-center justify-center max-w-screen-sm max-h-screen md:max-w-screen-md">
			<img {src} class="md:border-8 md:border-base-content" alt="{alt}_modal" />
		</div>
	</div>
{/if}
