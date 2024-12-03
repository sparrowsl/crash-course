import matter from "gray-matter";
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
		//@ts-ignore
		const { content } = matter(file);

		return {
			slug: file_path.split("./").pop()?.replace(".md", ""),
			content,
		};
	});

	const found = topics.find((topic) => topic.slug === params.slug);
	const parsed = marked.parse(String(found?.content), { gfm: true });

	return { content: parsed };

	// const post = (await import(`./${params.topic}.md`));
	// console.log((await modules[`./${params.topic}.md`]()).default);
}
