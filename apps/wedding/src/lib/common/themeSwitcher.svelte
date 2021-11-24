<script>
	let themes = [
		'Aqua',
		'Black',
		'Bumblebee',
		'Corporate',
		'Cupcake',
		'Cyberpunk',
		'Dark',
		'Dracula',
		'Emerald',
		'Fantasy',
		'Forest',
		'Garden',
		'Light',
		'Lofi',
		'Luxury',
		'Pastel',
		'Retro',
		'Synthwave',
		'Valentine',
		'Wireframe'
	];
</script>

<div class="dropdown dropdown-left">
	<div tabindex="0" class="btn">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			class="inline-block w-6 h-6 stroke-current md:mr-2"
			><path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"
			/></svg
		>
		<span class="hidden md:block">Theme</span>
	</div>
	<ul tabindex="0" class="shadow menu dropdown-content bg-neutral-focus w-52 pb-2">
		{#each themes as theme}
			<li>
				<button
					class="btn btn-primary mt-2 mx-4"
					data-set-theme={theme.toLowerCase()}
					data-act-class="ACTIVECLASS">{theme}</button
				>
			</li>
		{/each}
	</ul>
</div>
