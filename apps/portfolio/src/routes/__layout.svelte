<script>
	import '../app.css';
</script>

<div class="navbar mb-2 shadow-lg bg-neutral text-neutral-content">
	<div class="flex-none px-2 mx-2">
		<a class="text-lg font-bold" href="/"> ND </a>
	</div>
	<div class="flex-1 px-2 mx-2">
		<div class="items-stretch">
			<a class="btn btn-ghost btn-sm rounded-btn" href="/"> Home </a>
			<a class="btn btn-ghost btn-sm rounded-btn" href="/portfolio"> Portfolio </a>
			<a class="btn btn-ghost btn-sm rounded-btn" href="/about"> About </a>
			<a class="btn btn-ghost btn-sm rounded-btn" href="/contact"> Contact </a>
		</div>
	</div>
	<div class="flex-none">
		<button class="btn btn-square btn-ghost drawer-button">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				class="inline-block w-6 h-6 stroke-current"
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
</div>
<slot />
