<script>
	import '../app.css';
	import NavLinks from '$lib/common/navLinks.svelte';
	import ThemeSwitch from '$lib/common/themeSwitch.svelte';
	import { onMount } from 'svelte';
	import { themeChange } from 'theme-change';

	onMount(async () => {
		themeChange(false);
	});
	let sidebarOpen = false;
</script>

<aside
	class:translate-x-full={sidebarOpen}
	class="absolute z-10 w-full h-full transition duration-500 ease-in-out transform -left-full bg-neutral-focus"
	on:click={() => (sidebarOpen = !sidebarOpen)}
>
	<div class="flex flex-col texl-xl text-neutral-content">
		<NavLinks />
	</div>
</aside>
<nav class="flex mb-2 shadow-lg navbar bg-neutral text-neutral-content">
	<div class="flex-none">
		<button
			class="z-20 btn btn-square btn-ghost lg:hidden"
			on:click={() => (sidebarOpen = !sidebarOpen)}
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				class="inline-block w-10 h-10 stroke-current"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 6h16M4 12h16M4 18h16"
				/>
			</svg>
		</button>
	</div>
	<div class="flex-1">
		<div class="px-2 mx-2 flex-0">
			<a class="text-lg font-bold" href="/"> ND </a>
		</div>
		<div class="flex-1 hidden px-2 mx-2 lg:block">
			<div class="items-stretch">
				<NavLinks />
			</div>
		</div>
	</div>
	<div class="flex-none">
		<ThemeSwitch />
	</div>
</nav>
<main class="h-screen">
	<slot />
</main>
<footer class="flex justify-center w-full p-12 bg-neutral">
	<span class="text-xl text-neutral-content">
		Copyright &copy; 2021 NDSquared. All Rights Reserved
	</span>
</footer>
