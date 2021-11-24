<div class="container mx-auto flex-col justify-center">
	{#each Array(9) as _, i}
		<div class="min-h-screen flex items-center justify-center" class:bg-accent={i % 2 === 1}>
			<img
				src="https://ik.imagekit.io/dannylongeuay/ndsq/nd_fave_{i}.jpg?tr=h-600"
				alt="nd_fave_{i}"
			/>
		</div>
	{/each}
</div>
