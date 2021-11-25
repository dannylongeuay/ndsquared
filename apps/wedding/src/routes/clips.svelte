<script>
	let clips = [
		'nd_down_the_aisle',
		'nd_ceremony_armpull',
		'nd_up_the_aisle',
		'nd_cheer',
		'nd_pond_pan',
		'nd_jump',
		'nd_dance_twirl',
		'nd_father_daughter_dance',
		'nd_silly_butt',
		'nd_trio_trouble',
		'nd_brother_groove',
		'nd_getting_down_n_dirty',
		'nd_family_pan',
		'nd_final_send_off',
		'nd_final_kiss'
	];
</script>

<div class="container flex-col justify-center mx-auto">
	{#each clips as clip, i}
		<div
			class="flex items-center justify-center min-h-screen border-b-8 border-base-300"
			class:bg-accent={i % 2 === 0}
		>
			<!-- svelte-ignore a11y-media-has-caption -->
			<video autoplay loop muted playsinline class="border-8 border-neutral-focus">
				<!-- <source
					src="https://ik.imagekit.io/dannylongeuay/ndsq/{clip}.mp4?tr=f-webm"
					type="video/webm"
				/> -->
				<source src="https://ik.imagekit.io/dannylongeuay/ndsq/{clip}.mp4" type="video/mp4" />
			</video>
		</div>
	{/each}
</div>
