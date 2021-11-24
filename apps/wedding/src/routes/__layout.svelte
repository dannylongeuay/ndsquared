<script>
	import '../app.css';
	import '../app.css';
	import NavLinks from '$lib/common/navLinks.svelte';
	import ThemeSwitcher from '$lib/common/themeSwitcher.svelte';
	import { onMount } from 'svelte';
	import { themeChange } from 'theme-change';

	onMount(async () => {
		themeChange(false);
	});
	let mobileMenuOpen = false;
	let hideScrollTop = true;
	let showOnPx = 150;

	function goTop() {
		document.body.scrollIntoView();
	}

	function scrollContainer() {
		return document.documentElement || document.body;
	}

	function handleOnScroll() {
		if (!scrollContainer()) {
			return;
		}
		if (scrollContainer().scrollTop > showOnPx) {
			hideScrollTop = false;
		} else {
			hideScrollTop = true;
		}
	}
</script>

<svelte:head>
	<title>NDSquared | Wedding</title>
</svelte:head>
<svelte:window on:scroll={handleOnScroll} />

<aside
	class:translate-x-full={mobileMenuOpen}
	class="absolute z-10 w-full h-full transition duration-500 ease-in-out transform bg-neutral-focus -left-full"
	on:click={() => (mobileMenuOpen = !mobileMenuOpen)}
>
	<div class="flex flex-col text-xl text-neutral-content">
		<NavLinks />
	</div>
</aside>

<nav class="navbar bg-neutral text-neutral-content">
	<div class="flex-none">
		<button
			class="z-20 btn btn-square btn-ghost lg:hidden"
			on:click={() => (mobileMenuOpen = !mobileMenuOpen)}
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
	<div class="items-center flex-1">
		<div class="px-2 pb-2 mx-2 flex-0">
			<a class="text-4xl font-bold text-primary" href="/"> ND </a>
		</div>
		<div class="flex-1 hidden px-2 mx-2 lg:block">
			<div class="items-stretch">
				<NavLinks />
			</div>
		</div>
	</div>
	<div class="flex-none">
		<ThemeSwitcher />
	</div>
</nav>
<main class="min-h-screen">
	<slot />
</main>
<div class="fixed bottom-0 right-0 mb-2 mr-2" class:hidden={hideScrollTop}>
	<button class="btn btn-primary btn-circle" on:click={goTop}>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			class="w-5 h-5"
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
		</svg>
	</button>
</div>
<footer class="flex justify-center w-full p-12 bg-neutral">
	<span class="text-xl text-neutral-content">
		Copyright &copy; 2021 NDSquared. All Rights Reserved
	</span>
</footer>
