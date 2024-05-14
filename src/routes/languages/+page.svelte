<script>
import "iconify-icon";
// biome-ignore lint/correctness/noUnusedImports: <explanation>
import { languages } from "./languages.js";
</script>

<article class="mx-auto max-w-3xl py-10">
	<h2 class="mt-10 text-2xl max-w-lg mx-auto text-center">
		Languages currently supported!
	</h2>
	<p class="italic text-center font-light">
		<strong>Disclaimer:</strong> some text/sentences are taken from books, docs,
		etc. Almost all the exercises are taken from different programming books with
		different languages _(also probably different concepts)_.
	</p>

	<section class="grid grid-cols-2 gap-5 mt-10">
		{#each languages as language (language.name)}
			<figure class="bg-white shadow-md p-5 flex gap-6 items-center rounded-md">
				<iconify-icon icon={language.icon?.name} class="text-5xl self-start">
				</iconify-icon>

				<figcaption class="w-full grid gap-1">
					<h2 class="font-semibold mt-2 text-accent text-xl">
						{language.name}
					</h2>
					<p class="text-sm">
						<strong class="text-gray-500">Concepts:</strong>
						{language.concepts} concepts
					</p>
					<p class="text-sm mb-2">
						<strong class="text-gray-500">Exercises:</strong>
						{language.exercises} exercises
					</p>
					<!-- <p class="border text-sm w-fit">Status - WIP (coming soon)</p> -->
					{#if language.supported}
						<a
							href={language.url}
							class="block w-fit text-xs ml-auto text-accent font-bold px-3 py-1 rounded-sm mt-2"
						>
							Get Started ->
						</a>
					{:else}
						<span class="block w-fit text-xs ml-auto bg-accent/20 px-3 py-1">
							WIP - coming soon
						</span>
					{/if}
				</figcaption>
			</figure>
		{/each}
	</section>
</article>
