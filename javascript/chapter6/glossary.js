const glossary = {
	"variable": "A container to store or hold a value",
	"golang": "A weird but cool language to work with",
	"TDD": "Test Driven Development, write code to test your code",
	"frontend": "What the users see and interract with",
	"svelte": "The best javascript framework currently",
};

Object.entries(glossary).forEach(([k, v]) => console.log(`${k}:\n\t${v}\n`));
