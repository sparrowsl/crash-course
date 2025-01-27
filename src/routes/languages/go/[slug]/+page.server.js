import { marked } from "marked";

const modules = import.meta.glob("./*.md", {
	eager: true,
	query: "?raw",
	import: "default",
});

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	const topics = Object.keys(modules);

	for (const path of topics) {
		const slug = path.split("./").pop()?.replace(".md", "");
		const file = modules[path];
		const parsed = await marked.parse(/** @type {String} */ (file), { async: true, gfm: true });
		console.log(parsed);

		if (slug === params.slug) {
			return { content: parsed };
		}
	}
}
