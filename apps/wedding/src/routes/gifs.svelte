<script>
	let gifs = ['nd_up_the_aisle', 'nd_ceremony_armpull', 'nd_down_the_aisle', 'nd_dance_twirl'];
</script>

<div class="container mx-auto flex justify-center">
	<div class="flex-col">
		{#each gifs as gif}
			<img class="rounded-lg p-2" src="https://static.ndsquared.net/{gif}.gif" alt={gif} />
		{/each}
	</div>
</div>
