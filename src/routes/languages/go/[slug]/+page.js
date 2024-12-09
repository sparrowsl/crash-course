import { marked } from "marked";

const modules = import.meta.glob("./*.md", {
	eager: true,
	query: "?raw",
	import: "default",
});

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const topics = Object.keys(modules).map((file_path) => {
		const file = modules[file_path];
		const parsed = marked.parse(file, { gfm: true });
		console.log(parsed);

		return {
			slug: file_path.split("./").pop()?.replace(".md", ""),
			content: parsed,
		};
	});

	// console.log(topic);
	const found = topics.find((topic) => topic.slug === params.slug);

	return { content: found?.content };
}
