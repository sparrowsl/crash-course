import { marked } from "marked";

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const modules = import.meta.glob("./*.md", {
		eager: true,
		query: "?raw",
		import: "default",
	});

	const topics = Object.keys(modules).map((file_path) => {
		const file = modules[file_path];
		const parsed = marked.parse(file);

		return {
			slug: file_path.split("./").pop()?.replace(".md", ""),
			content: parsed,
		};
	});
	// console.log(topic);
	const found = topics.find((topic) => topic.slug === params.slug);

	return { content: found?.content };
	return { content: "" };

	// const post = (await import(`./${params.topic}.md`));
	// console.log((await modules[`./${params.topic}.md`]()).default);
}
