<div class="container mx-auto flex-col justify-center">
	{#each Array(32) as _, i}
		<div class="min-h-screen flex items-center justify-center" class:bg-accent={i % 2 === 1}>
			<img
				src="https://ik.imagekit.io/dannylongeuay/ndsq/nd_fave_{i + 1}.jpg?tr=h-800"
				alt="nd_fave_{i + 1}"
			/>
		</div>
	{/each}
</div>
