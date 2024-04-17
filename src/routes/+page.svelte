<script>
import "iconify-icon";
</script>

<main class="bg-[#fafafb] min-h-[100dvh] font-fira-sans">
	<section class="grid place-content-center min-h-[70dvh] text-center">
		<p class="text-sm border px-3 py-2 bg-white rounded-full w-fit mx-auto">
			Made by developers, for developers
		</p>
		<h1 class="text-6xl text-accent mb-2">Crash Course</h1>
		<p class="font-light max-w-xl text-center text-lg">
			Bringing you up to speed in multiple languages so you can build programs
			and tools that solve problems!
		</p>
		<a
			href="/"
			class="mt-3 bg-accent text-white px-6 py-2 rounded-md max-w-fit mx-auto"
		>
			Choose a Language
		</a>
	</section>

	<!-- Overview -->
	<section class="mx-auto max-w-3xl py-10">
		<h2 class="mb-10 text-2xl max-w-lg mx-auto text-center">
			Multiple languages, with byte size snippets and a lot of exercises!
		</h2>

		<div class="grid gap-5 grid-cols-2">
			{#each { length: 4 } as _}
				<!--Overview Card-->
				<div class="bg-white border rounded-lg p-5 shadow">
					<h3 class="mb-1 font-bold text-accent">Lorem ipsum dolor sit.</h3>
					<p class="text-sm font-roboto font-light">
						Lorem, ipsum dolor sit amet consectetur adipisicing elit. Totam
						sint, dolorem iusto maxime doloribus ipsam.
					</p>
				</div>
			{/each}
		</div>
	</section>

	<!-- List of languages supported -->
	<section class="mx-auto max-w-3xl py-10">
		<h2 class="mb-10 text-2xl max-w-lg mx-auto text-center">
			Languages currently supported!
		</h2>

		<div class="flex justify-center gap-10 flex-wrap max-w-xl mx-auto">
			{#each { length: 8 } as _}
				<a href="/" class="block">
					<span
						class="mx-auto p-2 rounded-full grid place-content-center w-fit"
					>
						<iconify-icon icon="simple-icons:javascript"></iconify-icon>
					</span>
					<h4 class="font-bold text-gray-700 mt-2">JavaScript</h4>
				</a>
			{/each}
		</div>
	</section>
</main>
